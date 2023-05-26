package utils

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	b64 "encoding/base64"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"google.golang.org/api/option"
	recaptchapb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
)

type utilsRepository struct{}
type UtilsRepository interface {
	SendEmail(*ses.SendEmailInput, *ses.SendRawEmailInput) int
	VerifyRecaptcha(string) (bool, error)
	SaveJobsToSpreadSheet([]string)
}

func NewEmail() *utilsRepository {
	return &utilsRepository{}
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

func (repo *utilsRepository) VerifyRecaptcha(token string) (bool, error) {

	//create recaptcha assessment and verify token
	ctx := context.Background()

	credBytes, err := b64.StdEncoding.DecodeString(os.Getenv("RECAPTCHA_CONFIG_JSON_BASE64"))
	if err != nil {
		log.Error(err)
		return false, err
	}

	client, err := recaptcha.NewClient(ctx, option.WithCredentialsJSON(credBytes))
	if err != nil {
		log.Error(err)
		return false, err
	}
	defer client.Close()

	// Build the assessment request
	request := &recaptchapb.CreateAssessmentRequest{
		Assessment: &recaptchapb.Assessment{
			Event: &recaptchapb.Event{
				Token:   token,
				SiteKey: os.Getenv("RECAPTCHA_SITE_KEY"),
			},
		},
		Parent: fmt.Sprintf("projects/%s", os.Getenv("RECAPTCHA_PROJECT_ID")),
	}

	response, err := client.CreateAssessment(ctx, request)

	if err != nil {
		log.Error(err)
		return false, err
	}

	// Interpret and verify assessment response
	if response.TokenProperties.Action == "verify" && response.TokenProperties.Valid && response.RiskAnalysis.Score >= 0.9 {
		return true, nil
	}

	return false, nil
}

// make array or slice unique
func Unique[T comparable](arr []T) []T {
	occurred := make(map[T]bool)
	var result []T
	for _, e := range arr {
		if !occurred[e] {
			occurred[e] = true
			result = append(result, e)
		}
	}
	return result
}

func GenerateUniqueFileName(filename string) string {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	name = strings.Replace(name, " ", "", -1)
	return fmt.Sprintf("%s_%d%s", name, time.Now().Unix(), ext)
}
