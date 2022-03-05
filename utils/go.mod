module utils

go 1.17

replace db => ../db

require (
	db v0.0.0
	github.com/aws/aws-sdk-go v1.43.5
	github.com/jmoiron/sqlx v1.3.4
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)
