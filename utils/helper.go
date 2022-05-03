package utils

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type emailRepository struct{}
type EmailRepository interface {
	SendEmail(*ses.SendEmailInput, *ses.SendRawEmailInput) int
}

func NewEmail() *emailRepository {
	return &emailRepository{}
}

func GetAWSIAMUserSession() (*session.Session, error) {
	awsRegion := os.Getenv("REGION")
	awsAccessKeyId := os.Getenv("ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, ""),
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return sess, err

}

func (repo *emailRepository) SendEmail(emailTemplate *ses.SendEmailInput, jobsTemplate *ses.SendRawEmailInput) int {

	sess, err := GetAWSIAMUserSession()

	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError
	}

	service := ses.New(sess)

	// Attempt to send the email.
	if emailTemplate != nil {
		_, err = service.SendEmail(emailTemplate)
	} else {
		_, err = service.SendRawEmail(jobsTemplate)

	}

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.Error(aerr.Error())
			return http.StatusInternalServerError
		} else {
			log.Error(err)
			return http.StatusBadRequest
		}
	}

	return 0
}
