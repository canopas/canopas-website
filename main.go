package main

import (
	"canopas-website/contact"
	"embed"
	"log"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

//go:embed templates/email-template.html
var tmplFS embed.FS

func main() {
	contact.TemplateFs = tmplFS
	router := gin.Default()

	contactRepo := contact.New()

	router.POST("/api/send-contact-mail", contactRepo.ContactDetail)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	if inLambda() {
		log.Fatal(gateway.ListenAndServe(":8080", router))
	} else {
		router.Run()
	}
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}
