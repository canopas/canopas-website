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

	password := os.Getenv("DB_PASSWORD")

	host := os.Getenv("DB_HOST")

	port := os.Getenv("DB_PORT")

	name := os.Getenv("DB_NAME")

	db = sqlx.MustConnect("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+name)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	db.SetConnMaxLifetime(time.Minute * 1)

	return db
}
