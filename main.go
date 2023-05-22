package main

import (
	"blogs"
	"contact"
	"contribution"
	"db"
	"embed"
	"jobs"
	"sitemap"
	"utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//go:embed templates/contact-email-template.html
//go:embed templates/career-email-template.html
//go:embed templates/client-thankyou-email-template.html
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

	utilsRepo := utils.NewEmail()

	contactRepo := contact.New(templateFS, utilsRepo)
	jobsRepo := jobs.New(sqlDb, templateFS, utilsRepo)
	sitemapRepo := sitemap.New(jobsRepo)

	router.POST("/api/send-contact-mail", contactRepo.SendContactMail)

	router.GET("/api/careers", jobsRepo.Careers)

	router.GET("/api/careers/:unique_id", jobsRepo.CareerById)

	router.POST("/api/send-career-mail", jobsRepo.SendCareerMail)

	router.GET("/api/sitemap", sitemapRepo.GenerateSitemap)

	router.GET("/api/blogs", blogs.Get)

	router.GET("/api/github/stars", contribution.GetStargazers)

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	router.Run(":8080")

	return router
}

func corsConfig() cors.Config {
	defaultCors := cors.DefaultConfig()
	defaultCors.AllowAllOrigins = true
	return defaultCors
}
