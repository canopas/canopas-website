package main

import (
	"blogs"
	"contact"
	"context"
	"contribution"
	"db"
	"embed"
	"jobs"
	"leave"
	"net/http"
	"notification"
	"os"
	"sitemap"
	"utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//go:embed templates/*.html
var templateFS embed.FS

var router *gin.Engine
var ginLambda *ginadapter.GinLambda

func inLambda() bool {
	log.Info("LAMBDA_TASK_ROOT: ", os.Getenv("LAMBDA_TASK_ROOT"))
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Info("In Lambda Handler...")
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func setupRouter() *gin.Engine {

	sqlDb := db.NewSql()

	router := gin.Default()

	router.Use(cors.New(corsConfig()))
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	utilsRepo := utils.NewEmail()
	contactRepo := contact.New(templateFS, utilsRepo)
	jobsRepo := jobs.New(sqlDb, templateFS, utilsRepo)
	sitemapRepo := sitemap.New(jobsRepo)
	leaveRepo := leave.New(templateFS, utilsRepo)
	notificationRepo := notification.New(templateFS, utilsRepo)

	router.POST("/api/send-contact-mail", contactRepo.SendContactMail)

	router.GET("/api/careers", jobsRepo.Careers)

	router.GET("/api/careers/:unique_id", jobsRepo.CareerById)

	router.POST("/api/send-jobs-applications", jobsRepo.SaveApplicationsData)

	router.GET("/api/sitemap", sitemapRepo.GenerateSitemap)

	router.GET("/api/blogs", blogs.Get)

	router.GET("/api/github/stars", contribution.GetStargazers)

	router.POST("/api/leave/new", leaveRepo.SendLeaveRequest)

	router.POST("/api/leave/update", leaveRepo.SendUpdateLeaveMail)

	router.POST("/api/invitation", notificationRepo.SendInvitationMail)

	router.POST("/api/acceptence", notificationRepo.SendAcceptenceMail)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return router
}

func corsConfig() cors.Config {
	defaultCors := cors.DefaultConfig()
	defaultCors.AllowAllOrigins = true
	return defaultCors
}

func main() {
	log.Info("router: ", router)
	router = setupRouter()
	log.Info("router: ", router)
	ginLambda = ginadapter.New(router)
	inLambda := inLambda()
	log.Info("inLambda: ", inLambda)
	if inLambda {
		log.Info("running aws lambda in aws")
		lambda.Start(LambdaHandler)
	} else {
		log.Info("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":8082", router))
	}
}
