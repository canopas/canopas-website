package contact

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ContactDetails struct {
	Name             string            `json:"name"`
	Designation      string            `json:"designation"`
	DesignationInfo  string            `json:"designation_info"`
	SocialMediaLinks map[string]string `json:"social_media_links"`
	Idea             string            `json:"idea"`
	Reason           string            `json:"reason"`
	Email            string            `json:"email"`
	Message          string            `json:"message"`
	ContactType      string            `json:"contact_type"`
}

const CHARSET = "utf-8"

var TemplateFs embed.FS

type Template struct {
	templates *template.Template
}

func New() *Template {
	templates, _ := template.ParseFS(TemplateFs, "templates/email-template.html")
	return &Template{
		templates: templates,
	}
}

func (repository *Template) ContactDetail(c *gin.Context) {
	var input ContactDetails

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailTemplate := repository.getEmailTemplate(input)

	statusCode := SendEmail(emailTemplate)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, input)
}

func (repository *Template) getEmailTemplate(input ContactDetails) (template *ses.SendEmailInput) {
	SENDER := os.Getenv("SENDER")
	RECEIVER := os.Getenv("RECEIVER")

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input)

	fmt.Println(htmlBody)

	title := "Canopas Website - Contact Information (By " + input.ContactType + ")"

	template = GetEmailTemplate(htmlBody, input, title, SENDER, RECEIVER)

	return
}

func (repository *Template) getHTMLBodyOfEmailTemplate(input ContactDetails) string {

	var templateBuffer bytes.Buffer

	err := repository.templates.ExecuteTemplate(&templateBuffer, "email-template.html", input)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return templateBuffer.String()
}

func SendEmail(emailTemplate *ses.SendEmailInput) int {

	sess, err := GetAWSIAMUserSession()

	if err != nil {
		log.Fatal(err)
		return http.StatusInternalServerError
	}

	service := ses.New(sess)

	// Attempt to send the email.
	_, err = service.SendEmail(emailTemplate)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.Fatal(aerr.Error())
			return http.StatusInternalServerError
		} else {
			log.Fatal(err)
			return http.StatusBadRequest
		}
	}

	return 0
}

func GetAWSIAMUserSession() (*session.Session, error) {
	awsRegion := os.Getenv("REGION_AWS")
	awsAccessKeyId := os.Getenv("ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, ""),
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return sess, err

}

func GetEmailTemplate(htmlBody string, data ContactDetails, title string, sender string, receiver string) (template *ses.SendEmailInput) {
	contactData := ConcatenateContactData(data)

	template = &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(receiver),
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
					Data:    aws.String(contactData),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(title),
			},
		},
		Source: aws.String(sender),
	}
	return
}

func ConcatenateContactData(data ContactDetails) (contactData string) {
	values := []string{}
	values = append(values, data.Name, data.Designation, data.DesignationInfo, data.Idea, data.Email, data.Message)

	contactData = strings.Join(values, "")

	return contactData
}
