package db

import (
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	log "github.com/sirupsen/logrus"
)

func NewSql() *sqlx.DB {

	log.Info("Initializing New Mysql Instance")

	var db *sqlx.DB
	username := os.Getenv("DB_USERNAME")

	if username == "" {
		username = "root"
	}

	password := os.Getenv("DB_PASSWORD")

	if password == "" {
		password = "root"
	}

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")

	if port == "" {
		port = "9906"
	}

	name := os.Getenv("DB_NAME")

	if name == "" {
		name = "website_admin"
	}

	db = sqlx.MustConnect("mysql", username+":"+password+"@("+host+":"+port+")/"+name)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	db.SetConnMaxLifetime(time.Minute * 1)

	return db
}
