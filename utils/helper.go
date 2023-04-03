package utils

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"

	b64 "encoding/base64"
	"encoding/json"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/ses"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
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

func (repo *utilsRepository) SaveJobsToSpreadSheet(records []string) {
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

	// create the batch request
	batchUpdateRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AppendCells: &sheets.AppendCellsRequest{
					Fields:  "*",
					Rows:    populateCells(records),
					SheetId: int64(sheetId),
				},
			},
		},
	}

	// execute the request
	res, err := srv.Spreadsheets.BatchUpdate(spreadsheetId, &batchUpdateRequest).Context(ctx).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		log.Error(err)
		return
	}
}

func populateCells(records []string) []*sheets.RowData {
	cells := []*sheets.CellData{}
	for i := range records {
		data := &sheets.CellData{
			UserEnteredValue: &sheets.ExtendedValue{
				StringValue: &(records[i]),
			},
			UserEnteredFormat: &sheets.CellFormat{
				BackgroundColor: &sheets.Color{ // white background
					Alpha: 1,
					Blue:  1,
					Red:   1,
					Green: 1,
				},
			},
		}
		cells = append(cells, data)
	}

	return []*sheets.RowData{
		{Values: cells},
	}
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

func ReadSliceFromFile[T comparable](filePath string, items []T) []T {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Warn(err)
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Warn(err)
	}

	return items
}

func WriteSliceToFile[T comparable](filePath string, items []T) {

	data, err := json.Marshal(items)

	if err != nil {
		log.Warn(err)
	}

	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		log.Warn(err)
	}
}

func DownloadFileFromS3(fileName string, sess *session.Session) {
	downloader := s3manager.NewDownloader(sess)
	file, err := os.Create(fileName)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String("canopas-blogs"),
			Key:    aws.String(fileName),
		})

	if err != nil {
		log.Error("Unable to download item %q, %v", fileName, err)
	}

	return
}

func UploadFileToS3(fileName string, sess *session.Session) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	file.Read(fileBuffer)

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String("canopas-blogs"),
		Key:                aws.String("/" + fileName),
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(fileBuffer),
		ContentLength:      aws.Int64(fileSize),
		ContentType:        aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		log.Error(err)
	}

	return
}
