module canopas-website

go 1.17

replace contact => ./contact

replace db => ./db

replace jobs => ./jobs

replace sitemap => ./sitemap

replace utils => ./utils

require (
	contact v0.0.0-00010101000000-000000000000
	db v0.0.0
	github.com/apex/gateway v1.1.2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/jmoiron/sqlx v1.3.4
	jobs v0.0.0
	sitemap v0.0.0
)

require (
	cloud.google.com/go/compute v1.6.1 // indirect
	github.com/aws/aws-lambda-go v1.17.0 // indirect
	github.com/aws/aws-sdk-go v1.43.5 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20220513224357-95641704303c // indirect
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5 // indirect
	golang.org/x/sys v0.0.0-20220513210249-45d2b4557a2a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/api v0.80.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/guregu/null.v3 v3.5.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	utils v0.0.0 // indirect
)
