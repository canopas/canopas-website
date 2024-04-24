package db

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func BlogDBInstance() *sqlx.DB {

	log.Info("Initializing New Blog Postgres Instance")

	username := os.Getenv("DB_USERNAME")
	if username == "" {
		username = "docker_user"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "docker_user"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5433"
	}

	name := os.Getenv("BLOG_DB_NAME")
	if name == "" {
		name = "canopas_website"
	}

	sslmode := "require"

	if os.Getenv("DB_ENV") == "test" {
		sslmode = "disable"
	}

	var db *sqlx.DB
	
	db, err := sqlx.Open("postgres", "postgres://"+username+":"+password+"@"+host+":"+port+"/"+name+"?sslmode="+sslmode)
	if err != nil {
		fmt.Println("Error while connecting to the blog database: ", err)
	}

	log.Info("Blog database instance initialized successfully")
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetConnMaxIdleTime(15 * time.Second)
	db.SetMaxOpenConns(2)

	return db
}
