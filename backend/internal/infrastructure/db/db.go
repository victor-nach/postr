package db

import (
    "database/sql"
    "fmt"
    "path/filepath"
    "time"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/victor-nach/postr-backend/pkg/migrator"
    _ "modernc.org/sqlite"
)

// New initialzes the sqlite db and applies the latest migrations
func New() (*gorm.DB, *sql.DB, error) {
    dbFile := filepath.Join(".", "data", "app.db")

    dsn := fmt.Sprintf("file:%s?cache=shared&_busy_timeout=5000&_journal_mode=WAL", dbFile)

    sqlDB, err := sql.Open("sqlite", dsn)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to open sql db: %w", err)
    }

    if err := sqlDB.Ping(); err != nil {
        return nil, nil, fmt.Errorf("failed to ping db: %w", err)
    }

    // Configure connection pool settings
    sqlDB.SetMaxOpenConns(2)
    sqlDB.SetMaxIdleConns(2)
    sqlDB.SetConnMaxLifetime(time.Hour)
    sqlDB.SetConnMaxIdleTime(10 * time.Minute)

    gormDB, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
    if err != nil {
        return nil, nil, fmt.Errorf("failed to open gorm db: %w", err)
    }

    // Apply latest migrations
    if err := migrator.Migrate(sqlDB, "file://migrations"); err != nil {
        return nil, nil, fmt.Errorf("failed to apply latest migrations: %w", err)
    }

    return gormDB, sqlDB, nil
}
