package sitemap

import (
	"embed"
	"encoding/json"
	"jobs"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"utils"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

//go:embed templates/career-email-template.html
//go:embed templates/path.txt
var templateFS embed.FS

var repo *SitemapRepository
var careerRepo *jobs.CareerRepository
var err error
var testDB *sqlx.DB

func TestInit(t *testing.T) {
	testDB, err = utils.TestDB()
	if err != nil {
		t.Errorf("Error in initializing test DB: %v", err)
	}

	careerRepo = jobs.New(testDB, templateFS, &stubUtilsRepo{})
	repo = New(careerRepo, templateFS)
}

func TestSitemap(t *testing.T) {
	engine := gin.New()
	engine.GET("/api/sitemap", repo.GenerateSitemap)
	req, err := http.NewRequest("GET", "/api/sitemap?baseUrl=http://localhost:8080", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	var response []URL
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error decoding response XML: %v", err)
	}
	assert.Equal(t, expectedSitemapData(), response)
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

func expectedSitemapData() []URL {
	sitemapUrls := []URL{}
	sitemapUrls = append(sitemapUrls, expectedURLData()...)

	return sitemapUrls
}

func BeginningOfMonthDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func expectedURLData() []URL {
	t := time.Now()
	dateFormat := "2006-01-02"
	date := BeginningOfMonthDate(t).Format(dateFormat) + "T00:00:00.000"
	baseUrl := "http://localhost:8080"

	sitemapUrls := []URL{
		{Loc: baseUrl, Priority: `1`},
		{Loc: baseUrl + "/services", Priority: `0.9`},
		{Loc: baseUrl + "/portfolio", Priority: `0.9`},
		{Loc: baseUrl + `/contributions`, Priority: `0.9`},
		{Loc: baseUrl + "/resources", Priority: `0.9`},
		{Loc: baseUrl + "/blog", Priority: `0.9`},
		{Loc: baseUrl + "/about", Priority: `0.9`},
		{Loc: baseUrl + "/contact", Priority: `0.9`},
		{Loc: baseUrl + "/android-app-development", Priority: `0.9`},
		{Loc: baseUrl + "/ios-app-development", Priority: `0.9`},
		{Loc: baseUrl + `/mobile-app-development`, Priority: `0.9`},
		{Loc: baseUrl + `/thank-you`, Priority: `0.9`},
		{Loc: baseUrl + `/jobs/thank-you`, Priority: `0.9`},
		{Loc: baseUrl + `/unsubscribe`, Priority: `0.9`},
		{Loc: baseUrl + "/jobs", Priority: `1`},
		{Loc: baseUrl + "/jobs/ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d", Priority: `0.9`},
		{Loc: baseUrl + "/jobs/apply/ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d", Priority: `0.9`},
		{Loc: baseUrl + "/portfolio/luxeradio", Priority: `0.9`},
		{Loc: baseUrl + "/portfolio/togness", Priority: `0.9`},
		{Loc: baseUrl + "/portfolio/justly", Priority: `0.9`},
	}

	for i := range sitemapUrls {
		sitemapUrls[i].LastMod = date
	}

	return sitemapUrls
}
