package utils

import (
	"net/http"

	"github.com/canopas/go-reusables/email"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func (repo *utilsRepository) SendEmail(data *email.EmailData) int {

	sess, err := GetAWSIAMUserSession()

	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError
	}

	// Attempt to send the email.
	_, err = email.SendAWSSESEmail(sess, data)

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
