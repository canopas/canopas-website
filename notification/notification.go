package notification

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
	"utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

const (
	CHARSET = "utf-8"
)

type NotificationData struct {
	Receiver    string `json:"receiver" form:"receiver"`
	CompanyName string `json:"companyname" form:"companyname"`
	SpaceLink   string `json:"spacelink" form:"spacelink"`
	Sender      string `json:"sender" form:"sender"`
}

type NotificationRepository struct {
	templates *template.Template
	UtilsRepo utils.UtilsRepository
}

func New(templateFs embed.FS, utilsRepo utils.UtilsRepository) *NotificationRepository {
	templates, _ := template.ParseFS(templateFs, "templates/*.html")

	return &NotificationRepository{
		templates: templates, UtilsRepo: utilsRepo,
	}
}

func (repository *NotificationRepository) SendInvitationMail(c *gin.Context) {
	var input NotificationData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailTemplate := repository.getInvitationEmailTemplate(input)

	statusCode := repository.UtilsRepo.SendEmail(emailTemplate, nil)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Invitation mail has been sent successfully")

}

func (repository *NotificationRepository) getInvitationEmailTemplate(input NotificationData) (template *ses.SendEmailInput) {

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input, "invitation-email-template.html")

	subject := "Join Unity!"

	template = GetEmailTemplate(htmlBody, input, subject)

	return
}

func (repository *NotificationRepository) SendAcceptenceMail(c *gin.Context) {
	var input NotificationData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailTemplate := repository.getAcceptenceEmailTemplate(input)

	statusCode := repository.UtilsRepo.SendEmail(emailTemplate, nil)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Acceptence mail has been sent successfully")

}

func (repository *NotificationRepository) getAcceptenceEmailTemplate(input NotificationData) (template *ses.SendEmailInput) {

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input, "acceptence-email-template.html")

	subject := "Unity | Accepted the Invitation"

	template = GetEmailTemplate(htmlBody, input, subject)

	return
}

func (repository *NotificationRepository) getHTMLBodyOfEmailTemplate(input NotificationData, templateName string) string {

	var templateBuffer bytes.Buffer

	err := repository.templates.ExecuteTemplate(&templateBuffer, templateName, input)

	if err != nil {
		log.Error(err)
		return ""
	}

	return templateBuffer.String()
}

func GetEmailTemplate(htmlBody string, data NotificationData, subject string) (template *ses.SendEmailInput) {

	SENDER := "Unity <unity@canopas.com>"
	RECEIVER := data.Receiver
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
			},
			Subject: &ses.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(SENDER),
	}

	return
}
