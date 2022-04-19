package main

import (
	"contact"
	"db"
	"embed"
	"jobs"
	"log"
	"os"
	"sitemap"

	"github.com/apex/gateway"
	"github.com/gin-contrib/cors"
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

	contactRepo := contact.New(templateFS)
	jobsRepo := jobs.New(sqlDb, templateFS)
	sitemapRepo := sitemap.New(jobsRepo)

	router.POST("/api/send-contact-mail", contactRepo.SendContactMail)

	router.GET("/api/careers", jobsRepo.Careers)

	router.GET("/api/careers/:id", jobsRepo.CareerById)

	router.GET("/api/update-uniqueid", jobsRepo.UpdateUniqueId)

	router.POST("/api/send-career-mail", jobsRepo.SendCareerMail)

	router.GET("/sitemap", sitemapRepo.GenerateSitemap)

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
