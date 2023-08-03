package notification

import (
	"embed"
	"net/http"
	"utils"

	"github.com/canopas/go-reusables/email"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	templates embed.FS
	UtilsRepo utils.UtilsRepository
}

func New(templateFs embed.FS, utilsRepo utils.UtilsRepository) *NotificationRepository {
	return &NotificationRepository{
		templates: templateFs, UtilsRepo: utilsRepo,
	}
}

func (repository *NotificationRepository) SendInvitationMail(c *gin.Context) {
	var input NotificationData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	statusCode := repository.sendEmail(input, "Join Unity!", "invitation-email-template.html")
	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Invitation mail has been sent successfully")

}

func (repository *NotificationRepository) SendAcceptenceMail(c *gin.Context) {
	var input NotificationData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	statusCode := repository.sendEmail(input, "Unity | Accepted the Invitation", "acceptence-email-template.html")

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Acceptence mail has been sent successfully")

}

func (repository *NotificationRepository) sendEmail(data NotificationData, title string, templateName string) int {

	emailData := &email.EmailData{
		Title:            title,
		Sender:           "Unity <unity@canopas.com>",
		Receiver:         data.Receiver,
		Charset:          CHARSET,
		TemplateFs:       repository.templates,
		TemplatePatterns: "templates/*.html",
		TemplateName:     templateName,
		Input:            data,
	}

	return repository.UtilsRepo.SendEmail(emailData)
}
