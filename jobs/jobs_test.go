package jobs

import (
	"bytes"
	"embed"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/tj/assert"
)

//go:embed templates/career-email-template.html
var templateFS embed.FS

var contentType string
var repo *CareerRepository
var err error
var testDB *sqlx.DB
var testRequests []utils.TestRequest

const (
	GET_ALL_JOBS_API_URL   = "/api/careers"
	GET_JOBS_BY_ID_API_URL = "/api/careers/ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d"
	SEND_CAREER_MAIL       = "/api/send-career-mail"
)

func Test_init(t *testing.T) {
	repo, err = initializeRepo()
	assert.Nil(t, err)
	testRequests = []utils.TestRequest{
		{
			Url:               GET_ALL_JOBS_API_URL,
			Method:            "GET",
			Headers:           nil,
			Body:              nil,
			ResponseCode:      http.StatusOK,
			ResponseTypeArray: true,
			ExpectedData:      []interface{}{expectedJobsData()},
		},
		{
			Url:               GET_JOBS_BY_ID_API_URL,
			Method:            "GET",
			Headers:           nil,
			Body:              nil,
			ResponseCode:      http.StatusOK,
			ResponseTypeArray: false,
			ExpectedData:      expectedJobsData(),
		},
		{
			Url:               SEND_CAREER_MAIL,
			Method:            "POST",
			Headers:           map[string]interface{}{"Content-Type": contentType},
			Body:              prepareRequestBody(),
			ResponseCode:      http.StatusOK,
			ResponseTypeArray: false,
			ExpectedData:      expectedMailData(),
		},
	}
}

func TestAllAPIs(t *testing.T) {

	asserts := assert.New(t)
	engine := gin.New()

	setUpRouter(engine)

	for _, testData := range testRequests {

		w := httptest.NewRecorder()
		var req *http.Request
		var got interface{}

		if testData.Body != nil {
			req, err = http.NewRequest(testData.Method, testData.Url, testData.Body.(io.Reader))
		} else {
			req, err = http.NewRequest(testData.Method, testData.Url, nil)
		}

		asserts.NoError(err)

		// set headers if present in request
		utils.SetRequestHeaders(req, testData.Headers, contentType)

		engine.ServeHTTP(w, req)
		assert.Equal(t, testData.ResponseCode, w.Code)
		if testData.ResponseTypeArray {
			got = utils.GotArrayData(w, t)
		} else {
			got = utils.GotData(w, t)
		}

		assert.Equal(t, testData.ExpectedData, got)
	}
}

func initializeRepo() (*CareerRepository, error) {

	testDB, err = utils.TestDB()
	if err != nil {
		return nil, err
	}

	err = utils.CreateTables(testDB)
	if err != nil {
		return nil, err
	}

	utils.TruncateTables(testDB)
	utils.PrepareTablesData(testDB)

	repo = New(testDB, templateFS)

	return repo, err
}

// configure api you want to test
func setUpRouter(engine *gin.Engine) {
	engine.GET(GET_ALL_JOBS_API_URL, repo.Careers)
	engine.GET("/api/careers/:id", repo.CareerById)
	engine.POST(SEND_CAREER_MAIL, repo.SendCareerMail)
}

func expectedJobsData() map[string]interface{} {
	jobs := make(map[string]interface{})
	jobs["id"] = 1.0
	jobs["title"] = "IOS Developer"
	jobs["summary"] = "We have an opening for a passionate iOS developer who can develop and maintain applications for iOS devices. As an iOS developer you will be a part of a highly agile team tasked with developing new features. If you love to tackle new challenges, we would love to meet you!"
	jobs["description"] = "<p><span style=\"color:#16a085;\"><span style=\"font-family:Arial,Helvetica,sans-serif;\">&nbsp;Develop mobile applications (iOS: Swift)</span></span></p>\r\n"
	jobs["button_name"] = "Apply"
	jobs["qualification"] = "B.E/B.Tech/BCA/MCA/MSc IT degree in Computer Science, Engineering, or a related subject "
	jobs["employment_type"] = "full_time"
	jobs["base_salary"] = 15000.0
	jobs["experience"] = "0-3 years"
	jobs["is_active"] = true
	jobs["skills"] = ""
	jobs["total_openings"] = 1.0
	jobs["responsibilities"] = ""
	jobs["icon_name"] = "fab fa-apple"
	jobs["unique_id"] = "ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d"
	jobs["seo_title"] = "iOS developer jobs in surat at Canopas, Mobile developer jobs"
	jobs["seo_description"] = "Find iOS developer job in surat. If we hire you as an iOS developer, you will be responsible for designing and coding the base application, ensuring the quality, fixing bugs, maintaining the code, and implementing app updates."
	jobs["apply_seo_title"] = "Apply for iOS developer job in surat or mobile developer jobs in surat at canopas software"
	jobs["apply_seo_description"] = "Apply for iOS developer job at Canopas and be part of a dynamic and versatile iOS app development team."
	return jobs
}
func expectedMailData() map[string]interface{} {
	career := make(map[string]interface{})

	career["job_title"] = "Web Developer from testing"
	career["name"] = "New Web Developer"
	career["email"] = "developer@gmail.com"
	career["phone"] = "1234567890"
	career["place"] = "surat"
	career["references"] = "From canopas"
	career["message"] = "I'm a very good programer"
	return career
}

func prepareRequestBody() *bytes.Buffer {

	content_type, body := utils.PrepareTextFileFormData()

	contentType = content_type

	return body
}
