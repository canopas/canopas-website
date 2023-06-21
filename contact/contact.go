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
	Name             string `json:"name"`
	Email            string `json:"email"`
	ProjectInfo      string `json:"project_info"`
	PhoneNumber      string `json:"phone_number"`
	Reference        string `json:"reference"`
	ContactType      string `json:"contact_type"`
	Token            string `json:"token"`
	Invest           string `json:"invest"`
	NDA              bool   `json:"nda"`
	SendMailToClient bool   `json:"send_mail_to_client" form:"send_mail_to_client"`
}

type Template struct {
	templates *template.Template
	UtilsRepo utils.UtilsRepository
}

func New(templateFs embed.FS, utilsRepo utils.UtilsRepository) *Template {
	templates := template.Must(template.ParseFS(templateFs, "templates/*.html"))

	return &Template{
		templates: templates, UtilsRepo: utilsRepo,
	}
}
func (repository *Template) SendContactMail(c *gin.Context) {
	var input ContactDetails

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if input.Token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	success, err := repository.UtilsRepo.VerifyRecaptcha(input.Token)

	if err != nil || !success {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	emailTemplate := repository.getEmailTemplate(input)

	statusCode := repository.UtilsRepo.SendEmail(emailTemplate, nil)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}
	if input.SendMailToClient {
		clientEmailTemplate := repository.getClientEmailTemplate(input)

		statusCode = repository.UtilsRepo.SendEmail(clientEmailTemplate, nil)

		if statusCode != 0 {
			c.AbortWithStatus(statusCode)
			return
		}
	}
	c.JSON(http.StatusOK, "Contact mail sent successfully")
}

func (repository *Template) getEmailTemplate(input ContactDetails) (template *ses.SendEmailInput) {

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input, "contact-email-template.html")

	extendedTitle := ""
	if input.ContactType != "" {
		extendedTitle = "(By " + input.ContactType + ")"
	}

	title := "Canopas Website - Contact Information " + extendedTitle

	template = GetEmailTemplate(htmlBody, input, title, os.Getenv("CONTACT_RECEIVER"))

	return
}

func (repository *Template) getClientEmailTemplate(input ContactDetails) (template *ses.SendEmailInput) {

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input, "client-thankyou-email-template.html")

	title := "Thank you " + input.Name + " for choosing Canopas!"

	template = GetEmailTemplate(htmlBody, input, title, input.Email)

	return
}

func (repository *Template) getHTMLBodyOfEmailTemplate(input ContactDetails, templateName string) string {

	var templateBuffer bytes.Buffer

	err := repository.templates.ExecuteTemplate(&templateBuffer, templateName, input)

	if err != nil {
		log.Error(err)
		return ""
	}

	return templateBuffer.String()
}

func GetEmailTemplate(htmlBody string, data ContactDetails, title string, receiver string) (template *ses.SendEmailInput) {

	SENDER := os.Getenv("CONTACT_SENDER")
	RECEIVER := receiver
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
