package sourcetypes

import (
	"testing"

	"github.com/gear6io/pragmata/pkg/types/querybuildertypes"
)

func TestCreateTableStmt(t *testing.T) {
	tests := []struct {
		name    string
		src     Source
		expectedErr bool
		expectedSQL string
	}{
		{
			name:    "no fields returns error",
			src:     Source{Name: "empty"},
			expectedErr: true,
		},
		{
			name: "defaults engine to MergeTree and includes all fields",
			src: Source{
				Name:   "events",
				Engine: EngineUndefined,
				Fields: []querybuildertypes.Field{
					{Name: "id", Type: querybuildertypes.FieldDataTypeInt64},
					{Name: "ts", Type: querybuildertypes.FieldDataTypeDateTime64},
				},
			},
			expectedSQL: "CREATE TABLE IF NOT EXISTS events (id Int64, ts DateTime64(3), __attrs__ JSON) ENGINE = MergeTree() ORDER BY (id,ts)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctb, err := tc.src.CreateTableStmt()
			if tc.expectedErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			sql, _ := ctb.Build()
			if sql != tc.expectedSQL {
				t.Errorf("SQL mismatch\ngot:  %s\nwant: %s", sql, tc.expectedSQL)
			}
		})
	}
}
