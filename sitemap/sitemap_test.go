package sitemap

import (
	"bytes"
	"embed"
	"encoding/json"
	"encoding/xml"
	"jobs"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

//go:embed templates/career-email-template.html
var templateFS embed.FS

var repo *SitemapRepository
var careerRepo *jobs.CareerRepository
var err error
var testDB *sqlx.DB
var testRequests []utils.TestRequest

const (
	GET_SITEMAP_API_URL = "/sitemap?baseUrl=http://localhost:8080"
)

func Test_init(t *testing.T) {
	repo, err = initializeRepo()
	assert.Nil(t, err)
	testRequests = []utils.TestRequest{
		{
			Url:               GET_SITEMAP_API_URL,
			Method:            "GET",
			Headers:           nil,
			Body:              nil,
			ResponseCode:      http.StatusOK,
			ResponseTypeArray: false,
			ExpectedData:      expectedSitemapData(),
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

		if testData.Body != nil {
			requestByte, _ := json.Marshal(testData.Body)
			reqBodyData := bytes.NewReader(requestByte)
			req, err = http.NewRequest(testData.Method, testData.Url, reqBodyData)
		} else {
			req, err = http.NewRequest(testData.Method, testData.Url, nil)
		}

		asserts.NoError(err)

		engine.ServeHTTP(w, req)
		assert.Equal(t, testData.ResponseCode, w.Code)

		var urlset URLset

		err = xml.Unmarshal(w.Body.Bytes(), &urlset)

		assert.Equal(t, testData.ExpectedData, urlset)

	}
}

func initializeRepo() (*SitemapRepository, error) {

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

	careerRepo = jobs.New(testDB, templateFS, nil)
	repo = New(careerRepo)

	return repo, err
}

func setUpRouter(engine *gin.Engine) {
	engine.GET("/sitemap", repo.GenerateSitemap)
}

func expectedSitemapData() URLset {
	sitemap := URLset{}

	sitemap.XMLName.Space = XMLNS
	sitemap.XMLName.Local = "urlset"
	sitemap.XMLNS = XMLNS
	sitemap.URL = append(sitemap.URL, expectedURLData()...)

	return sitemap
}

func BeginningOfMonthDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func expectedURLData() []URL {
	t := time.Now()
	dateFormat := "2006-01-02"
	date := BeginningOfMonthDate(t).Format(dateFormat) + "T00:00:00.000Z"
	baseUrl := "http://localhost:8080"

	sitemapUrls := []URL{
		{Loc: baseUrl, Priority: `1`},
		{Loc: baseUrl + "/contact", Priority: `0.9`},
		{Loc: BLOG_URL, Priority: `0.8`},
		{Loc: baseUrl + "/portfolio", Priority: `0.9`},
		{Loc: baseUrl + "/portfolio/luxeradio", Priority: `0.9`},
		{Loc: baseUrl + "/portfolio/togness", Priority: `0.9`},
		{Loc: baseUrl + "/portfolio/nolonely", Priority: `0.9`},
		{Loc: baseUrl + "/jobs", Priority: `1`},
		{Loc: baseUrl + "/jobs/ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d", Priority: `0.9`},
		{Loc: baseUrl + "/jobs/apply/ios-developer-a9b45f34-a1a5-419f-b536-b7c290925d6d", Priority: `0.9`},
	}

	for i := range sitemapUrls {
		sitemapUrls[i].XMLName.Space = XMLNS
		sitemapUrls[i].XMLName.Local = "url"
		sitemapUrls[i].ChangeFreq = "monthly"
		sitemapUrls[i].LastMod = date
	}

	return sitemapUrls
}
