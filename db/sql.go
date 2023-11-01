package db

import (
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func NewSql() *sqlx.DB {

	log.Info("Initializing New Postgresql Instance")

	var db *sqlx.DB
	username := os.Getenv("DB_USERNAME")

	if username == "" {
		username = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")

	if password == "" {
		password = "postgres"
	}

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")

	if port == "" {
		port = "5432"
	}

	name := os.Getenv("DB_NAME")

	if name == "" {
		name = "postgres"
	}

	sslmode := "require"
	if os.Getenv("DB_ENV") == "test" {
		sslmode = "disable"
	}

	db = sqlx.MustConnect("postgres", "postgres://"+username+":"+password+"@"+host+":"+port+"/"+name+"?sslmode="+sslmode)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	db.SetConnMaxLifetime(time.Minute * 1)

	return db
}
