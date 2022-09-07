module utils

go 1.18

replace db => ../db

require (
	cloud.google.com/go/recaptchaenterprise/v2 v2.0.1
	db v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v1.44.83
	github.com/jmoiron/sqlx v1.3.5
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/api v0.94.0
	google.golang.org/genproto v0.0.0-20220902135211-223410557253
)

require (
	cloud.google.com/go/compute v1.7.0 // indirect
	cloud.google.com/go/recaptchaenterprise v1.3.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/oauth2 v0.0.0-20220822191816-0ebed06d0094 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
