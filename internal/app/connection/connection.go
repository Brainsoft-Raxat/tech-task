package connection

import (
	"fmt"
	"log"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection struct {
	Postgres *sqlx.DB
}

func New(cfg *config.Configs) (*Connection, error) {
	postgres, err := postgresConnection(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("postgres сonnection: %v", err)
	}
	
	return &Connection{
		Postgres: postgres,
	}, nil
}

func (c *Connection) Close() {
	// Close connections
	if c.Postgres != nil {
		_ = c.Postgres.Close()
	}
}

func postgresConnection(cfg config.Postgres) (*sqlx.DB, error) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	log.Print(datasource)

	return sqlx.Open("postgres", datasource)
}