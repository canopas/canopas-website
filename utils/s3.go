package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

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
