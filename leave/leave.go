package leave

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

type LeaveData struct {
	Name        string `json:"name" form:"name"`
	Date        string `json:"date" form:"date"`
	Status      int    `json:"status" form:"status"`
	Reason      string `json:"reason" form:"reason"`
	Receiver    string `json:"receiver" form:"receiver"`
	StatusValue string
}

type LeaveRepository struct {
	templates embed.FS
	UtilsRepo utils.UtilsRepository
}

func New(templateFs embed.FS, utilsRepo utils.UtilsRepository) *LeaveRepository {
	return &LeaveRepository{
		templates: templateFs, UtilsRepo: utilsRepo,
	}
}

func (repository *LeaveRepository) SendLeaveRequest(c *gin.Context) {
	var input LeaveData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	input.StatusValue = utils.GetStatusValue(input.Status)

	statusCode := repository.sendEmail(input, "New leave request", "new-leave-email-template.html")

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Leave request has been sent successfully")
}

func (repository *LeaveRepository) SendUpdateLeaveMail(c *gin.Context) {
	var input LeaveData

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	input.StatusValue = utils.GetStatusValue(input.Status)

	statusCode := repository.sendEmail(input, "Leave request update", "update-leave-email-template.html")

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, "Update leave request has been sent successfully")
}

func (repository *LeaveRepository) sendEmail(data LeaveData, title string, templateName string) int {

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
