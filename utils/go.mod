module utils

go 1.18

replace db => ../db

require (
	github.com/aws/aws-sdk-go v1.44.83
	github.com/jmoiron/sqlx v1.3.5
	github.com/sirupsen/logrus v1.9.0
)

require (
	db v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
