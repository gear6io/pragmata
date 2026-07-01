//go:build integration

package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/testcontainers/testcontainers-go"
	tcnetwork "github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	baseURL string
	chConn  driver.Conn

	// TODO(auth): server has no auth middleware yet; authToken is unused.
	// Once auth is enforced, populate authToken before m.Run():
	// options: pre-insert a token row into pragmata.db, add --dev-token flag to serve,
	// or read PRAGMATA_TEST_TOKEN env var.
	authToken = ""
)

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

func run(m *testing.M) int {
	ctx := context.Background()

	// Isolated Docker network so containers can reach each other by hostname.
	net, err := tcnetwork.New(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create network:", err)
		return 1
	}
	defer net.Remove(ctx) //nolint:errcheck
	networkName := net.Name

	// ClickHouse single-node — no Zookeeper needed (MergeTree, no replication).
	// Alias "clickhouse" lets the Pragmata container reach it by hostname.
	chCtr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "clickhouse/clickhouse-server:25.5.6",
			ExposedPorts: []string{"9000/tcp"},
			WaitingFor:   wait.ForListeningPort("9000/tcp").WithStartupTimeout(60 * time.Second),
			Networks:     []string{networkName},
			NetworkAliases: map[string][]string{
				networkName: {"clickhouse"},
			},
 			Env: map[string]string{
				"CLICKHOUSE_PASSWORD": "test",
			},
		},
		Started: true,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "start ClickHouse:", err)
		return 1
	}
	defer chCtr.Terminate(ctx) //nolint:errcheck

	// chConn uses the host-mapped port for direct data seeding from test code.
	chHost, err := chCtr.Host(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ClickHouse host:", err)
		return 1
	}
	chPort, err := chCtr.MappedPort(ctx, "9000")
	if err != nil {
		fmt.Fprintln(os.Stderr, "ClickHouse port:", err)
		return 1
	}
	chConn, err = clickhouse.Open(&clickhouse.Options{
		Addr:     []string{fmt.Sprintf("%s:%s", chHost, chPort.Port())},
		Protocol: clickhouse.Native,
		Auth:     clickhouse.Auth{Database: "default", Username: "default", Password: "test"},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "open ClickHouse conn:", err)
		return 1
	}
	defer chConn.Close()

	// TODO(migration): replace with the actual migration utility once it exists.
	if err := setupClickHouse(ctx, chConn); err != nil {
		fmt.Fprintln(os.Stderr, "setup ClickHouse:", err)
		return 1
	}

	repoRoot, err := repoRootDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "repo root:", err)
		return 1
	}

	// Pragmata config passed directly into the container via Reader — no host temp file needed.
	// ClickHouse URL uses the internal network alias, not the host-mapped port.
	cfgContent := `
server:
  host: 0.0.0.0
  port: 7181
clickhouse:
  url: clickhouse://default:test@clickhouse:9000
sqlmesh:
  project_dir: /tmp/pragmata-sqlmesh
database:
  path: /tmp/pragmata.db
`

	// Build and run Pragmata from the repo Dockerfile.
	// Docker layer cache makes subsequent runs fast after the first full compile.
	// frontend/dist must be pre-built (`make build-ui`) before running these tests.
	pgCtr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:        repoRoot,
				Dockerfile:     "tests/integration/Dockerfile",
				BuildLogWriter: os.Stderr,
				KeepImage:      true, // reuse built image across runs; Docker cache handles staleness
			},
			ExposedPorts: []string{"7181/tcp"},
			WaitingFor: wait.ForHTTP("/api/v0/pipes").
				WithPort("7181/tcp").
				WithStartupTimeout(120 * time.Second),
			Networks: []string{networkName},
			Files: []testcontainers.ContainerFile{
				{
					Reader:            strings.NewReader(cfgContent),
					ContainerFilePath: "/config.yaml",
					FileMode:          0o644,
				},
			},
			Cmd: []string{"serve", "-c", "/config.yaml"},
		},
		Started: true,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "start Pragmata:", err)
		return 1
	}
	defer pgCtr.Terminate(ctx) //nolint:errcheck

	pgHost, err := pgCtr.Host(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Pragmata host:", err)
		return 1
	}
	pgPort, err := pgCtr.MappedPort(ctx, "7181")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Pragmata port:", err)
		return 1
	}
	baseURL = fmt.Sprintf("http://%s:%s", pgHost, pgPort.Port())

	result := m.Run()

	// Dump all Pragmata container logs after tests finish — Logs() is a snapshot,
	// not a stream, so reading after m.Run() captures everything generated during tests.
	fmt.Fprintln(os.Stderr, "\n--- pragmata container logs ---")
	if logReader, err := pgCtr.Logs(ctx); err == nil {
		io.Copy(os.Stderr, logReader) //nolint:errcheck
		logReader.Close()             //nolint:errcheck
	}

	return result
}

// apiDo sends a JSON request; caller must not use resp after returning — readResponse
// and assertNoContent handle body drain/close.
func apiDo(t *testing.T, method, path string, body any) *http.Response {
	t.Helper()
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("marshal request: %v", err)
		}
		buf = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(context.Background(), method, baseURL+path, buf)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if authToken != "" {
		req.Header.Set("Authorization", "Bearer "+authToken)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("do %s %s: %v", method, path, err)
	}
	return resp
}

// readResponse reads + closes the body, asserts status == wantStatus, returns the data field.
func readResponse(t *testing.T, resp *http.Response, wantStatus int) json.RawMessage {
	t.Helper()
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if resp.StatusCode != wantStatus {
		t.Fatalf("want HTTP %d, got %d — %s", wantStatus, resp.StatusCode, raw)
	}
	var ar struct {
		Status string          `json:"status"`
		Data   json.RawMessage `json:"data"`
		Error  string          `json:"error"`
	}
	if err := json.Unmarshal(raw, &ar); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	return ar.Data
}

// assertNoContent closes the body and asserts HTTP 204.
func assertNoContent(t *testing.T, resp *http.Response) {
	t.Helper()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("want HTTP 204, got %d — %s", resp.StatusCode, body)
	}
}

// seedTestData inserts raw test rows into ClickHouse after source registration has created the tables.
// page_views: last row has empty user_id (intentionally dirty, filtered by clean_page_views).
// user_profiles: u4 has empty plan (filtered by clean_profiles) and country JP (no data after filter).
// Intermediate tables (clean_page_views, clean_profiles, enriched_sessions) are populated by sqlmesh
// when the TABLE pipes execute — do not seed them here.
func seedTestData(t *testing.T) {
	t.Helper()
	ctx := context.Background()

	if err := chConn.Exec(ctx, `
		INSERT INTO pragmata_source.page_views (timestamp, user_id, page) VALUES
		('2024-01-15 10:00:00', 'u1', '/home'),
		('2024-01-15 11:00:00', 'u2', '/pricing'),
		('2024-01-16 09:00:00', 'u1', '/home'),
		('2024-01-16 10:00:00', 'u3', '/docs'),
		('2024-01-16 11:00:00', 'u2', '/home'),
		('2024-01-16 12:00:00', '',   '/broken')
	`); err != nil {
		t.Fatalf("seed page_views: %v", err)
	}

	if err := chConn.Exec(ctx, `
		INSERT INTO pragmata_source.user_profiles (user_id, plan, country, signup_date) VALUES
		('u1', 'pro',  'US', '2023-01-01'),
		('u2', 'pro',  'US', '2023-02-01'),
		('u3', 'free', 'DE', '2023-03-01'),
		('u4', '',     'JP', '2023-04-01')
	`); err != nil {
		t.Fatalf("seed user_profiles: %v", err)
	}
}

// setupClickHouse creates the pragmata_source database that the server expects to exist.
// TODO(migration): remove once a proper migration utility creates it on server startup.
func setupClickHouse(ctx context.Context, conn driver.Conn) error {
	return conn.Exec(ctx, "CREATE DATABASE IF NOT EXISTS pragmata_source")
}

func repoRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found from %s", dir)
		}
		dir = parent
	}
}
