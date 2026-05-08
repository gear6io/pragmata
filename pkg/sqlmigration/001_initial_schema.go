package sqlmigration

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

type initialSchema struct{}

func NewInitialSchema() SQLMigration {
	return &initialSchema{}
}

// Register must be called from this file so bun names the migration after it.
func (m *initialSchema) Register(migrations *migrate.Migrations) error {
	return migrations.Register(m.Up, m.Down)
}

func (m *initialSchema) Up(ctx context.Context, db *bun.DB) error {
	// table:pipes
	if _, err := db.NewCreateTable().
		Model(&struct {
			bun.BaseModel    `bun:"table:pipes"`
			Name             string    `bun:"name,pk,type:text"`
			Type             string    `bun:"type,notnull,type:text"`
			Description      string    `bun:"description,type:text,default:''"`
			Tags             string    `bun:"tags,notnull,type:text,default:'[]'"`
			Nodes            string    `bun:"nodes,notnull,type:text,default:'[]'"`
			Datasource       string    `bun:"datasource,type:text,default:''"`
			TargetDatasource string    `bun:"target_datasource,type:text,default:''"`
			CopySchedule     string    `bun:"copy_schedule,type:text,default:''"`
			CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp"`
			UpdatedAt        time.Time `bun:"updated_at,notnull,default:current_timestamp"`
		}{}).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}

	// table:backfill_jobs
	if _, err := db.NewCreateTable().
		Model(&struct {
			bun.BaseModel      `bun:"table:backfill_jobs"`
			ID                 string    `bun:"id,pk,type:text"`
			PipeID             string    `bun:"pipe_id,notnull,type:text"`
			TotalIntervals     int       `bun:"total_intervals,notnull,default:0"`
			CompletedIntervals int       `bun:"completed_intervals,notnull,default:0"`
			Status             string    `bun:"status,notnull,type:text,default:'running'"`
			LastError          string    `bun:"last_error,notnull,type:text,default:''"`
			UpdatedAt          time.Time `bun:"updated_at,notnull,default:current_timestamp"`
		}{}).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}

	// table:tokens
	if _, err := db.NewCreateTable().
		Model(&struct {
			bun.BaseModel `bun:"table:tokens"`
			ID            string    `bun:"id,pk,type:text"`
			Name          string    `bun:"name,notnull,type:text"`
			Value         string    `bun:"value,unique,notnull,type:text"`
			Scope         string    `bun:"scope,notnull,type:text,default:'PIPES:READ'"`
			CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
		}{}).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (m *initialSchema) Down(_ context.Context, _ *bun.DB) error {
	return nil
}
