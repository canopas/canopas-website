package main

import (
	"contact"
	"db"
	"embed"
	"jobs"
	"log"
	"os"
	"sitemap"
	"utils"

	"github.com/apex/gateway"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//go:embed templates/contact-email-template.html
//go:embed templates/career-email-template.html
var templateFS embed.FS

func main() {
	sqlDb := db.NewSql()
	defer sqlDb.Close()

	r := setupRouter(sqlDb)

	_ = r.Run(":8080")
}

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(corsConfig()))
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	emailRepo := utils.NewEmail()

	contactRepo := contact.New(templateFS, emailRepo)
	jobsRepo := jobs.New(sqlDb, templateFS, emailRepo)
	sitemapRepo := sitemap.New(jobsRepo)

	router.POST("/api/send-contact-mail", contactRepo.SendContactMail)

	router.GET("/api/careers", jobsRepo.Careers)

	router.GET("/api/careers/:unique_id", jobsRepo.CareerById)

	router.POST("/api/send-career-mail", jobsRepo.SendCareerMail)

	router.GET("/api/sitemap", sitemapRepo.GenerateSitemap)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	if inLambda() {
		log.Fatal(gateway.ListenAndServe(":8080", router))
	} else {
		router.Run()
	}

	return router
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func corsConfig() cors.Config {
	defaultCors := cors.DefaultConfig()
	defaultCors.AllowAllOrigins = true
	return defaultCors
}
