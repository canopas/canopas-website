package contact

import (
	"embed"
	"net/http"
	"os"
	"utils"

	"github.com/canopas/go-reusables/email"

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
	templateFs embed.FS
	UtilsRepo  utils.UtilsRepository
}

func New(templateFs embed.FS, utilsRepo utils.UtilsRepository) *Template {
	return &Template{
		templateFs: templateFs, UtilsRepo: utilsRepo,
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

	repository.sendEmail(input, "Canopas Website - Contact Information", os.Getenv("RECEIVER"), "contact-email-template.html")

	if input.SendMailToClient {

		repository.sendEmail(input, "Thank you "+input.Name+" for choosing Canopas!", input.Email, "client-thankyou-email-template.html")
	}

	c.JSON(http.StatusOK, "Contact mail sent successfully")
}

func (repository *Template) sendEmail(input ContactDetails, title string, receiver string, templateName string) int {

	data := &email.EmailData{
		Title:            title,
		Sender:           os.Getenv("SENDER"),
		Receiver:         receiver,
		Charset:          CHARSET,
		TemplateFs:       repository.templateFs,
		TemplatePatterns: "templates/*.html",
		TemplateName:     templateName,
		Input:            input,
	}

	return repository.UtilsRepo.SendEmail(data)
}
