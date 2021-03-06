package contact

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
	"os"
	"utils"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	CHARSET = "utf-8"
)

type ContactDetails struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	ProjectInfo string `json:"project_info"`
	Reference   string `json:"reference"`
	ContactType string `json:"contact_type"`
}

type Template struct {
	templates *template.Template
	EmailRepo utils.EmailRepository
}

func New(templateFs embed.FS, emailRepo utils.EmailRepository) *Template {
	templates, _ := template.ParseFS(templateFs, "templates/contact-email-template.html")
	return &Template{
		templates: templates, EmailRepo: emailRepo,
	}
}

func (repository *Template) SendContactMail(c *gin.Context) {
	var input ContactDetails

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailTemplate := repository.getEmailTemplate(input)

	statusCode := repository.EmailRepo.SendEmail(emailTemplate, nil)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, input)
}

func (repository *Template) getEmailTemplate(input ContactDetails) (template *ses.SendEmailInput) {

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input)

	title := "Canopas Website - Contact Information (By " + input.ContactType + ")"

	template = GetEmailTemplate(htmlBody, input, title)

	return
}

func (repository *Template) getHTMLBodyOfEmailTemplate(input ContactDetails) string {

	var templateBuffer bytes.Buffer

	err := repository.templates.ExecuteTemplate(&templateBuffer, "contact-email-template.html", input)

	if err != nil {
		log.Error(err)
		return ""
	}

	return templateBuffer.String()
}

func GetEmailTemplate(htmlBody string, data ContactDetails, title string) (template *ses.SendEmailInput) {

	SENDER := os.Getenv("SENDER")
	RECEIVER := os.Getenv("CONTACT_RECEIVER")

	template = &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(RECEIVER),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CHARSET),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CHARSET),
					Data:    aws.String("Contact Info"),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(title),
			},
		},
		Source: aws.String(SENDER),
	}
	return
}
