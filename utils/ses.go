package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
)

func (repo *utilsRepository) SendEmail(emailTemplate *ses.SendEmailInput, jobsTemplate *ses.SendRawEmailInput) int {
	
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
