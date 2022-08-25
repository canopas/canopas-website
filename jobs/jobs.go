package jobs

import (
	"bytes"
	"database/sql"
	"embed"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"utils"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gopkg.in/gomail.v2"
	"gopkg.in/guregu/null.v3"
)

const (
	CHARSET = "utf-8"
)

var extension string

type Career struct {
	Id                  int         `json:"id"`
	Title               string      `json:"title"`
	Summary             string      `json:"summary"`
	Description         string      `json:"description"`
	ButtonName          string      `json:"button_name"`
	Qualification       string      `json:"qualification"`
	EmploymentType      string      `json:"employment_type"`
	BaseSalary          int         `json:"base_salary"`
	Experience          string      `json:"experience"`
	IsActive            bool        `json:"is_active"`
	Skills              null.String `json:"skills"`
	TotalOpenings       int         `json:"total_openings"`
	Responsibilities    null.String `json:"responsibilities"`
	IconName            string      `json:"icon_name"`
	UniqueId            string      `json:"unique_id"`
	SEOTitle            null.String `json:"seo_title"`
	SEODescription      null.String `json:"seo_description"`
	ApplySEOTitle       null.String `json:"apply_seo_title"`
	ApplySEODescription null.String `json:"apply_seo_description"`
	Index               int         `json:"index"`
}

type CareerDetails struct {
	JobTitle   string `json:"job_title" form:"job_title"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Phone      string `json:"phone" form:"phone"`
	Place      string `json:"place" form:"place"`
	References string `json:"references" form:"references"`
	Message    string `json:"message" form:"message"`
}

type CareerRepository struct {
	Db        *sqlx.DB
	templates *template.Template
	EmailRepo utils.EmailRepository
}

func New(db *sqlx.DB, templateFs embed.FS, emailRepo utils.EmailRepository) *CareerRepository {
	templates, _ := template.ParseFS(templateFs, "templates/career-email-template.html")
	return &CareerRepository{Db: db, templates: templates, EmailRepo: emailRepo}
}

func (repository *CareerRepository) Careers(c *gin.Context) {

	careerList, err := repository.GetCareers()

	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, careerList)

}

func (repository *CareerRepository) GetCareers() ([]Career, error) {

	careers := []Career{}

	err := repository.Db.Select(&careers, "SELECT id, title, summary, description, button_name, qualification, employment_type,"+
		"base_salary, experience, is_active, skills, total_openings,"+
		"responsibilities, icon_name, unique_id, seo_title, seo_description,"+
		"apply_seo_title, apply_seo_description, `index` "+
		"FROM jobs WHERE is_active = 1 ORDER BY `index`")

	return careers, err
}

func (repository *CareerRepository) CareerById(c *gin.Context) {

	career := Career{}

	id := c.Param("unique_id")

	err := repository.Db.Get(&career, `SELECT id, title, summary, description, button_name, qualification, 
							   		   employment_type, base_salary, experience, is_active, skills, 
									   total_openings, responsibilities, icon_name, 
									   unique_id, seo_title, seo_description,
									   apply_seo_title, apply_seo_description
									   FROM jobs
									   WHERE unique_id = ?
									   AND is_active = 1`, id)

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, career)
}

func (repository *CareerRepository) SendCareerMail(c *gin.Context) {
	input := CareerDetails{}

	err := c.Bind(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	htmlBody := repository.getHTMLBodyOfEmailTemplate(input)

	title := input.JobTitle + " job application - Canopas Website"

	attachmentBytes, err := getFileBytes(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailTemplate := GetEmailTemplate(htmlBody, title, attachmentBytes)

	statusCode := repository.EmailRepo.SendEmail(nil, emailTemplate)

	if statusCode != 0 {
		c.AbortWithStatus(statusCode)
		return
	}

	c.JSON(http.StatusOK, input)
}

func getFileBytes(c *gin.Context) (*bytes.Buffer, error) {
	fileHeader, err := c.FormFile("file")

	extension = filepath.Ext(fileHeader.Filename)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	attachmentBytes := bytes.NewBuffer(nil)
	if _, err := io.Copy(attachmentBytes, file); err != nil {
		log.Error(err)
		return nil, err
	}

	return attachmentBytes, nil
}

func GetEmailTemplate(htmlBody string, title string, attachmentBytes *bytes.Buffer) (template *ses.SendRawEmailInput) {

	SENDER := os.Getenv("SENDER")
	RECEIVER := os.Getenv("JOBS_RECEIVER")

	msg := gomail.NewMessage()
	msg.SetHeader("From", SENDER)
	msg.SetHeader("To", RECEIVER)
	msg.SetHeader("Subject", title)
	msg.SetBody("text/html", htmlBody)

	msg.Attach("resume"+extension, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(attachmentBytes.Bytes())
		return err
	}))

	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)

	message := ses.RawMessage{
		Data: emailRaw.Bytes(),
	}

	template = &ses.SendRawEmailInput{
		RawMessage: &message,
		Destinations: []*string{
			aws.String(RECEIVER),
		},
		Source: aws.String(SENDER),
	}

	return template
}

func (repository *CareerRepository) getHTMLBodyOfEmailTemplate(input CareerDetails) string {

	var templateBuffer bytes.Buffer

	err := repository.templates.ExecuteTemplate(&templateBuffer, "career-email-template.html", input)

	if err != nil {
		log.Error(err)
		return ""
	}

	return templateBuffer.String()
}
