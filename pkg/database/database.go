package database

import (
	"database/sql"
	"erp-opity/pkg/config"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// DB representa a conexão com o banco de dados
type DB struct {
	*sql.DB
}

// Connect estabelece conexão com o banco de dados
func Connect(cfg config.DatabaseConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com banco: %w", err)
	}

	// Configurar pool de conexões
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime))

	// Testar conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao testar conexão com banco: %w", err)
	}

	return &DB{db}, nil
}

// Close fecha a conexão com o banco de dados
func (db *DB) Close() error {
	return db.DB.Close()
}

// Ping testa a conexão com o banco de dados
func (db *DB) Ping() error {
	return db.DB.Ping()
}

// Stats retorna estatísticas do pool de conexões
func (db *DB) Stats() sql.DBStats {
	return db.DB.Stats()
}
