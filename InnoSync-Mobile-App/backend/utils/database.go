package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/team-18-project/InnoSync/backend/config"
)

// Database holds the database connection
type Database struct {
	DB *sqlx.DB
}

// NewDatabase creates a new database connection
func NewDatabase(cfg *config.Config) (*Database, error) {
	db, err := sqlx.Connect("postgres", cfg.GetDatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connected successfully")
	return &Database{DB: db}, nil
}

// Close closes the database connection
func (d *Database) Close() error {
	return d.DB.Close()
}

// Migrate runs database migrations
func (d *Database) Migrate() error {
	// Check if tables exist and create if not
	migrations := []string{
		// `DROP TYPE IF EXISTS education_enum CASCADE;`,
		`DO $$ BEGIN CREATE TYPE education_enum AS ENUM ('NO_DEGREE', 'BACHELOR', 'MASTER', 'PHD'); EXCEPTION WHEN duplicate_object THEN null; END $$;`,
		// `DROP TYPE IF EXISTS expertise_level_enum CASCADE;`,
		`DO $$ BEGIN CREATE TYPE expertise_level_enum AS ENUM ('ENTRY', 'JUNIOR', 'MID', 'SENIOR', 'RESEARCHER'); EXCEPTION WHEN duplicate_object THEN null; END $$;`,
		`DO $$ BEGIN CREATE TYPE invitation_status_enum AS ENUM ('INVITED', 'ACCEPTED', 'DECLINED', 'REVOKED'); EXCEPTION WHEN duplicate_object THEN null; END $$;`,
		`DO $$ BEGIN CREATE TYPE application_status_enum AS ENUM ('PENDING', 'UNDER_REVIEW', 'ACCEPTED', 'REJECTED', 'WITHDRAWN'); EXCEPTION WHEN duplicate_object THEN null; END $$;`,
	}

	for _, migration := range migrations {
		if _, err := d.DB.Exec(migration); err != nil {
			log.Printf("Migration warning: %v", err)
		}
	}

	log.Println("Database migrations completed")
	return nil
}

// Health checks database health
func (d *Database) Health() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return d.DB.PingContext(ctx)
}

// Transaction executes a function within a database transaction
func (d *Database) Transaction(fn func(*sqlx.Tx) error) error {
	tx, err := d.DB.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
