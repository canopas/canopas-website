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

	err := os.Setenv("DB_ENV", "test")
	if err != nil {
		return nil, err
	}

	err = os.Setenv("DB_NAME", "website_admin_test")
	if err != nil {
		return nil, err
	}

	DB := db.NewSql()
	return DB, nil
}

func CreateTables(db *sqlx.DB) error {

	jobs := "CREATE TABLE IF NOT EXISTS jobs (id SERIAL PRIMARY KEY, title varchar, summary text, description text, button_name varchar, qualification varchar, employment_type varchar, base_salary int4, experience varchar, is_active bool DEFAULT false, created_at timestamptz DEFAULT CURRENT_TIMESTAMP, updated_at timestamptz DEFAULT CURRENT_TIMESTAMP, skills text, total_openings int4 DEFAULT 1, responsibilities text, icon_name int4, unique_id varchar NOT NULL, seo_title varchar, seo_description text, apply_seo_title varchar, apply_seo_description text, index int4 NOT NULL DEFAULT 0)"

	_, err := db.Exec(jobs)

	if err != nil {
		return err
	}

	jobApplications := "CREATE TABLE job_applications (id SERIAL PRIMARY KEY, name varchar NOT NULL, email varchar NOT NULL, phone varchar NOT NULL, place varchar, reference varchar, resumeurl varchar NOT NULL, position varchar NOT NULL, message text, status int4 NOT NULL DEFAULT 1, created_at timestamptz DEFAULT CURRENT_TIMESTAMP, updated_at timestamptz DEFAULT CURRENT_TIMESTAMP )"

	_, err = db.Exec(jobApplications)

	if err != nil {
		return err
	}

	jobsApplicantStatus := "CREATE TABLE job_applicant_statuses (id SERIAL PRIMARY KEY,applicant_id int4 NOT NULL,status int4 NOT NULL DEFAULT 1,index int4 NOT NULL DEFAULT 0,rejection_with_mail bool DEFAULT false,created_at timestamptz DEFAULT CURRENT_TIMESTAMP,updated_at timestamptz DEFAULT CURRENT_TIMESTAMP)"

	_, err = db.Exec(jobsApplicantStatus)

	if err != nil {
		return err
	}

	jobsIcons := "CREATE TABLE job_icons ( id SERIAL PRIMARY KEY, name varchar, value varchar, createdAt timestamptz, updatedAt timestamptz)"

	_, err = db.Exec(jobsIcons)


	lifeatCanopas := "CREATE TABLE lifeatcanopas (id SERIAL PRIMARY KEY, image_urls varchar, index int4 NOT NULL DEFAULT 0, created_at timestamptz DEFAULT CURRENT_TIMESTAMP,updated_at timestamptz DEFAULT CURRENT_TIMESTAMP)"

	_, err = db.Exec(lifeatCanopas)

	if err != nil {
		return err
	}

	perkImages := "CREATE TABLE perks (id SERIAL PRIMARY KEY, image_urls varchar, index int4 NOT NULL DEFAULT 0, created_at timestamptz DEFAULT CURRENT_TIMESTAMP,updated_at timestamptz DEFAULT CURRENT_TIMESTAMP)"

	_, err = db.Exec(perkImages)

	if err != nil {
		return err
	}

	return nil
}

func TruncateTables(db *sqlx.DB) {
	db.MustExec("DROP TABLE IF EXISTS jobs")
	db.MustExec("DROP TABLE IF EXISTS job_applications")
	db.MustExec("DROP TABLE IF EXISTS job_applicant_statuses")
	db.MustExec("DROP TABLE IF EXISTS job_icons")
	db.MustExec("DROP TABLE IF EXISTS lifeatcanopas")
	db.MustExec("DROP TABLE IF EXISTS perks")
}

func PrepareTablesData(db *sqlx.DB) {
	db.MustExec("INSERT INTO jobs(id, title, summary, description, button_name,qualification, employment_type, base_salary, experience, is_active, skills, total_openings, responsibilities, icon_name, unique_id, seo_title, seo_description, apply_seo_title, apply_seo_description, created_at, updated_at) VALUES(1, 'IOS Developer', 'We have an opening for a passionate iOS developer who can develop and maintain applications for iOS devices. As an iOS developer you will be a part of a highly agile team tasked with developing new features. If you love to tackle new challenges, we would love to meet you!', '<p><span style=\"color:#16a085;\"><span style=\"font-family:Arial,Helvetica,sans-serif;\">&nbsp;Develop mobile applications (iOS: Swift)</span></span></p>\r\n', 'Apply', 'B.E/B.Tech/BCA/MCA/MSc IT degree in Computer Science, Engineering, or a related subject ', 'full_time', 15000 , '0-3 years', true, '', 1, '',1, 'ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d', 'iOS developer jobs in surat at Canopas, Mobile developer jobs', 'Find iOS developer job in surat. If we hire you as an iOS developer, you will be responsible for designing and coding the base application, ensuring the quality, fixing bugs, maintaining the code, and implementing app updates.', 'Apply for iOS developer job in surat or mobile developer jobs in surat at canopas software', 'Apply for iOS developer job at Canopas and be part of a dynamic and versatile iOS app development team.',  current_timestamp, current_timestamp)")
	db.MustExec("INSERT INTO job_applications(id, name, email, phone, place,reference, resumeURL, position, message, status, created_at, updated_at) VALUES(1, 'New Web Developer', 'developer@gmail.com', '1234567890', 'surat', 'From canopas', '', 'Web Developer from testing', 'I m a very good programer', '1',current_timestamp, current_timestamp)")
	db.MustExec("INSERT INTO job_applicant_statuses(id, applicant_id, status, index, rejection_with_mail, created_at, updated_at) VALUES(1, 1, 1, 0, false, current_timestamp,current_timestamp )")
	db.MustExec("INSERT INTO job_icons (id,name,value,createdAt,updatedAt) VALUES ('1','iOS','fab fa-apple',NULL,NULL);")
    db.MustExec("INSERT INTO lifeatcanopas (id,image_urls,created_at,updated_at) VALUES(1,'https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-400.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-800.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-1600.webp',current_timestamp, current_timestamp)")
	db.MustExec("INSERT INTO perks (id,image_urls,created_at,updated_at) VALUES(1,'https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-400.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-800.webp',current_timestamp, current_timestamp)")
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
