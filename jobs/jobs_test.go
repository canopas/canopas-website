package jobs

import (
	"embed"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

//go:embed templates/career-email-template.html
var templateFS embed.FS
var repo *CareerRepository
var err error
var testDB *sqlx.DB

func TestInit(t *testing.T) {
	testDB, err = utils.TestDB()
	if err != nil {
		t.Errorf("Error in initializing test DB: %v", err)
	}

	repo = New(testDB, templateFS, &stubUtilsRepo{})
}

func TestGetAllJobs(t *testing.T) {
	utils.TruncateTables(testDB)
	err = utils.CreateTables(testDB)
	if err != nil {
		t.Errorf("Error in initializing test DB: %v", err)
	}
	utils.PrepareTablesData(testDB)
	engine := gin.New()
	engine.GET("/api/careers", repo.Careers)
	req, err := http.NewRequest("GET", "/api/careers", nil)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotArrayData(w, t)
	expected := []interface{}{expectedJobsData()}
	assert.Equal(t, expected, got)
}

func TestGetJobById(t *testing.T) {
	engine := gin.New()
	engine.GET("/api/careers/:unique_id", repo.CareerById)
	req, err := http.NewRequest("GET", "/api/careers/ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d", nil)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotData(w, t)
	assert.Equal(t, expectedJobsData(), got)
}

func TestSaveApplicationsData(t *testing.T) {
	engine := gin.New()
	engine.POST("/api/send-jobs-applications", repo.SaveApplicationsData)
	contentType, fileBody := utils.PrepareTextFileFormData()
	req, err := http.NewRequest("POST", "/api/send-jobs-applications", io.NopCloser(fileBody))
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}
	req.Header.Set("Content-Type", contentType)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotData(w, t)

	expected := "Job application received successfully"

	assert.Equal(t, expected, got)
}

// stubUtilsRepo is a mock Utils Service Interface
type stubUtilsRepo struct{}

func (faker *stubUtilsRepo) SendEmail(emailInput *ses.SendEmailInput, jobsInput *ses.SendRawEmailInput) int {
	return 0
}

func (faker *stubUtilsRepo) VerifyRecaptcha(token string) (bool, error) {
	return true, nil
}

func (faker *stubUtilsRepo) SaveJobsToSpreadSheet(input []string) {
	/** this is stub method for adding jobs details in google spreadsheet */
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
	jobs["base_salary"] = "15000"
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
	jobs["index"] = 0.0
	return jobs
}
