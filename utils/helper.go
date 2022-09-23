package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"

	b64 "encoding/base64"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	recaptchapb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
)

type utilsRepository struct{}
type UtilsRepository interface {
	SendEmail(*ses.SendEmailInput, *ses.SendRawEmailInput) int
	VerifyRecaptcha(string) (bool, error)
	SaveJobsToSpreadSheet([][]interface{})
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

func (repo *utilsRepository) SaveJobsToSpreadSheet(records [][]interface{}) {
	ctx := context.Background()

	credBytes, err := b64.StdEncoding.DecodeString(os.Getenv("RECAPTCHA_CONFIG_JSON_BASE64"))
	if err != nil {
		log.Error(err)
		return
	}

	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Error(err)
		return
	}

	client := config.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Error(err)
		return
	}

	/**
	spreadsheet url = https://docs.google.com/spreadsheets/d/<SPREADSHEETID>/edit#gid=<SHEETID>
	*/
	sheetId, err := strconv.Atoi(os.Getenv("JOBS_SHEET_ID"))
	if err != nil {
		log.Error(err)
		return
	}
	spreadsheetId := os.Getenv("JOBS_SPREADSHEET_ID")

	// Convert sheet ID to sheet name.
	response1, err := srv.Spreadsheets.Get(spreadsheetId).Fields("sheets(properties(sheetId,title))").Do()
	if err != nil || response1.HTTPStatusCode != 200 {
		log.Error(err)
		return
	}

	sheetName := ""
	for _, v := range response1.Sheets {
		prop := v.Properties
		if prop.SheetId == int64(sheetId) {
			sheetName = prop.Title
			break
		}
	}

	//Append value to the sheet.
	row := &sheets.ValueRange{
		Values: records,
	}
	response2, err := srv.Spreadsheets.Values.Append(spreadsheetId, sheetName, row).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(ctx).Do()
	if err != nil || response2.HTTPStatusCode != 200 {
		log.Error(err)
		return
	}
}
