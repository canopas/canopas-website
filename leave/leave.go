package leave

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

type LeaveData struct {
	Name        string `json:"name" form:"name"`
	Date        string `json:"date" form:"date"`
	Duration    string `json:"duration" form:"duration"`
	Status      int    `json:"status" form:"status"`
	Reason      string `json:"reason" form:"reason"`
	Receiver    string `json:"receiver" form:"receiver"`
	StatusValue string
}

type LeaveRepository struct {
	templates *template.Template
	UtilsRepo utils.UtilsRepository
}

func New(templateFs embed.FS, utilsRepo utils.UtilsRepository) *LeaveRepository {
	templates, _ := template.ParseFS(templateFs, "templates/*.html")

	return &LeaveRepository{
		templates: templates, UtilsRepo: utilsRepo,
	}
}

func (repository *LeaveRepository) SendLeaveRequest(c *gin.Context) {
	var input LeaveData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailTemplate := repository.getNewLeaveEmailTemplate(input)

	statusCode := repository.UtilsRepo.SendEmail(emailTemplate, nil)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Leave request has been sent successfully")

}

func (repository *LeaveRepository) getNewLeaveEmailTemplate(input LeaveData) (template *ses.SendEmailInput) {

	input.StatusValue = utils.GetStatusValue(input.Status)

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input, "new-leave-email-template.html")

	subject := "New leave request"

	template = GetEmailTemplate(htmlBody, input, subject)

	return
}

func (repository *LeaveRepository) SendUpdateLeaveMail(c *gin.Context) {
	var input LeaveData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailTemplate := repository.getUpdateLeaveEmailTemplate(input)

	statusCode := repository.UtilsRepo.SendEmail(emailTemplate, nil)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Update leave request has been sent successfully")

}

func (repository *LeaveRepository) getUpdateLeaveEmailTemplate(input LeaveData) (template *ses.SendEmailInput) {

	input.StatusValue = utils.GetStatusValue(input.Status)

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input, "update-leave-email-template.html")

	subject := "Leave request update"

	template = GetEmailTemplate(htmlBody, input, subject)

	return
}

func (repository *LeaveRepository) getHTMLBodyOfEmailTemplate(input LeaveData, templateName string) string {

	var templateBuffer bytes.Buffer

	err := repository.templates.ExecuteTemplate(&templateBuffer, templateName, input)

	if err != nil {
		log.Error(err)
		return ""
	}

	return templateBuffer.String()
}

func GetEmailTemplate(htmlBody string, data LeaveData, subject string) (template *ses.SendEmailInput) {

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
