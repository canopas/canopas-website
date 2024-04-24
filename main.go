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
	"post"
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
//go:embed templates/*.txt
var templateFS embed.FS

var router *gin.Engine
var ginLambda *ginadapter.GinLambda

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	sitemapRepo := sitemap.New(jobsRepo, templateFS)
	leaveRepo := leave.New(templateFS, utilsRepo)
	notificationRepo := notification.New(templateFS, utilsRepo)
	lifePerksImagesRepo := jobs.NewLifePerksImages(sqlDb)

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

	router.GET("/api/lifeimages", lifePerksImagesRepo.LifeImages)

	router.GET("/api/perksimages", lifePerksImagesRepo.PerksImages)

	blogSqlDb := db.BlogDBInstance()

	postRepo := post.New(blogSqlDb)

	router.GET("/api/posts", postRepo.Get)
	router.GET("/api/posts/:slug", postRepo.Show)

	router.GET("/api/posts/tags/:slug", postRepo.GetPostsByTag)

	router.GET("/api/posts/author/:username", postRepo.GetPostsByAuthor)

	return router
}

func corsConfig() cors.Config {
	defaultCors := cors.DefaultConfig()
	defaultCors.AllowAllOrigins = true
	return defaultCors
}

func main() {
	router = setupRouter()
	ginLambda = ginadapter.New(router)
	if inLambda() {
		log.Info("running aws lambda in aws")
		lambda.Start(LambdaHandler)
	} else {
		log.Info("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":8082", router))
	}
}
