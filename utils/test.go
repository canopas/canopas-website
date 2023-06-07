package utils

import (
	"bytes"
	"db"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TestRequest struct {
	Url               string
	Method            string
	Headers           map[string]interface{}
	Body              interface{}
	ResponseCode      int
	ResponseTypeArray bool
	ExpectedData      interface{}
}

const PDF_FILE_PATH = "../test_assets/test_text_file.pdf"

func GotData(w *httptest.ResponseRecorder, t *testing.T) interface{} {
	var got interface{}
	if len(w.Body.Bytes()) != 0 {
		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
	}
	return got
}

func GotArrayData(w *httptest.ResponseRecorder, t *testing.T) []interface{} {
	var got []interface{}
	if len(w.Body.Bytes()) != 0 {
		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
	}
	return got
}

func SetRequestHeaders(req *http.Request, headers map[string]interface{}, contentType string) {
	if len(headers) > 0 {
		for key, value := range headers {
			if key == "Content-Type" {
				req.Header.Add(key, contentType)
			} else {
				req.Header.Add(key, value.(string))
			}
		}
	}
}

func TestDB() (*sqlx.DB, error) {

	err := os.Setenv("DB_NAME", "website_admin_test")
	if err != nil {
		return nil, err
	}

	DB := db.NewSql()
	return DB, nil
}

func CreateTables(db *sqlx.DB) error {

	jobs := "CREATE TABLE IF NOT EXISTS `jobs` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`title` varchar(255) NULL DEFAULT NULL,`summary` varchar(500) NULL DEFAULT NULL,`description` text NULL DEFAULT NULL,`button_name` varchar(225) NULL DEFAULT NULL,`qualification` varchar(225) NULL DEFAULT NULL,`employment_type` varchar(225) NULL DEFAULT NULL,`base_salary` int(11) NULL DEFAULT NULL,`experience` varchar(225) NULL DEFAULT NULL,`is_active` tinyint(1) NULL DEFAULT 0,`skills` text NULL DEFAULT NULL,`total_openings` int(11) NULL DEFAULT 1,`responsibilities` text NULL DEFAULT NULL, `icon_name` varchar(255) NULL DEFAULT NULL , `unique_id` varchar(255) NULL DEFAULT NULL, `seo_title` varchar(255) NULL DEFAULT NULL, `seo_description` text NULL DEFAULT NULL, `apply_seo_title` varchar(255) NULL DEFAULT NULL, `apply_seo_description` text NULL DEFAULT NULL, `index` int(10) NULL DEFAULT 0, `created_at` timestamp NULL DEFAULT NULL,`updated_at` timestamp NULL DEFAULT NULL, PRIMARY KEY (`id`))"

	_, err := db.Exec(jobs)

	if err != nil {
		return err
	}
	jobApplications := "CREATE TABLE IF NOT EXISTS `job_applications` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`name` varchar(255) NULL DEFAULT NULL,`email` varchar(255) NULL DEFAULT NULL,`phone` varchar(255) NULL DEFAULT NULL,`place` varchar(255) NULL DEFAULT NULL,`reference` varchar(255) NULL DEFAULT NULL,`resumeURL` varchar(255) NULL DEFAULT NULL,`position` varchar(255) NULL DEFAULT NULL,`message` text NULL DEFAULT NULL,`status` int(10) NULL DEFAULT 1, `created_at` timestamp NULL DEFAULT NULL,`updated_at` timestamp NULL DEFAULT NULL, PRIMARY KEY (`id`))"

	_, err = db.Exec(jobApplications)

	if err != nil {
		return err
	}
	jobsApplicantStatus := "CREATE TABLE IF NOT EXISTS `jobs_applicant_status` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT, `applicant_id` int(10) unsigned NOT NULL DEFAULT 0, `status` int(10) NOT NULL DEFAULT 1, `index` int(10) NOT NULL DEFAULT 0, `rejection_with_mail` tinyint(1) NULL DEFAULT 0, `created_at` timestamp NULL DEFAULT NULL,`updated_at` timestamp NULL DEFAULT NULL, PRIMARY KEY (`id`) ,FOREIGN KEY (`applicant_id`) REFERENCES `job_applications` (`id`))"

	_, err = db.Exec(jobsApplicantStatus)

	if err != nil {
		return err
	}

	return nil
}

func TruncateTables(db *sqlx.DB) {
	db.MustExec("SET FOREIGN_KEY_CHECKS = 0") // Disable foreign key checks

	db.MustExec("TRUNCATE TABLE jobs")
	db.MustExec("TRUNCATE TABLE job_applications")
	db.MustExec("TRUNCATE TABLE jobs_applicant_status")

	db.MustExec("SET FOREIGN_KEY_CHECKS = 1") // Enable foreign key checks
}

func PrepareTablesData(db *sqlx.DB) {
	db.MustExec("INSERT INTO jobs(`id`, `title`, `summary`, `description`, `button_name`,`qualification`, `employment_type`, `base_salary`, `experience`, `is_active`, `skills`, `total_openings`, `responsibilities`, `icon_name`, `unique_id`, `seo_title`, `seo_description`, `apply_seo_title`, `apply_seo_description`, `created_at`, `updated_at`) VALUES(1, 'IOS Developer', 'We have an opening for a passionate iOS developer who can develop and maintain applications for iOS devices. As an iOS developer you will be a part of a highly agile team tasked with developing new features. If you love to tackle new challenges, we would love to meet you!', '<p><span style=\"color:#16a085;\"><span style=\"font-family:Arial,Helvetica,sans-serif;\">&nbsp;Develop mobile applications (iOS: Swift)</span></span></p>\r\n', 'Apply', 'B.E/B.Tech/BCA/MCA/MSc IT degree in Computer Science, Engineering, or a related subject ', 'full_time', 15000 , '0-3 years', true, '', 1, '','fab fa-apple', 'ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d', 'iOS developer jobs in surat at Canopas, Mobile developer jobs', 'Find iOS developer job in surat. If we hire you as an iOS developer, you will be responsible for designing and coding the base application, ensuring the quality, fixing bugs, maintaining the code, and implementing app updates.', 'Apply for iOS developer job in surat or mobile developer jobs in surat at canopas software', 'Apply for iOS developer job at Canopas and be part of a dynamic and versatile iOS app development team.',  UTC_TIMESTAMP(), UTC_TIMESTAMP())")
	db.MustExec("INSERT INTO job_applications(`id`, `name`, `email`, `phone`, `place`,`reference`, `resumeURL`, `position`, `message`, `status`, `created_at`, `updated_at`) VALUES(1, 'New Web Developer', 'developer@gmail.com', '1234567890', 'surat', 'From canopas', '', 'Web Developer from testing', 'I m a very good programer', '1',UTC_TIMESTAMP(), UTC_TIMESTAMP())")
	db.MustExec("INSERT INTO jobs_applicant_status(`id`, `applicant_id`, `status`, `index`, `rejection_with_mail`, `created_at`, `updated_at`) VALUES(1, 1, 1, 0, 0, UTC_TIMESTAMP(), UTC_TIMESTAMP())")
}

func PrepareTextFileFormData() (string, *bytes.Buffer) {

	var (
		file *os.File
		err  error
	)

	file, err = os.Open(PDF_FILE_PATH)

	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body) // create fileWriter to write bytes in body

	filePart, err := writer.CreateFormFile("file", filepath.Base(file.Name())) // create filePart

	if err != nil {
		log.Error(err)
	}

	io.Copy(filePart, file) // copy fileBytes to file

	contentType := writer.FormDataContentType() // get contentType of file

	writer.WriteField("job_title", "Web Developer from testing")
	writer.WriteField("name", "New Web Developer")
	writer.WriteField("email", "developer@gmail.com")
	writer.WriteField("phone", "1234567890")
	writer.WriteField("place", "surat")
	writer.WriteField("references", "From canopas")
	writer.WriteField("message", "I'm a very good programer")
	writer.WriteField("token", "xyz123")
	writer.Close()

	return contentType, body
}
