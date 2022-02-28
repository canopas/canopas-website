package sitemap

import (
	"jobs"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pahanini/go-sitemap-generator"
)

type SitemapRepository struct {
	careerRepo *jobs.CareerRepository
}

func New(careerRepo *jobs.CareerRepository) *SitemapRepository {
	return &SitemapRepository{careerRepo: careerRepo}
}

func (repository *SitemapRepository) GenerateSitemap(c *gin.Context) {
	baseUrl := c.Query("baseUrl")
	jobsUrl := baseUrl + "/jobs"
	blogsUrl := "https://blog.canopas.com"

	sitemapUrls := []sitemap.URL{
		{Loc: baseUrl, Priority: `1`},
		{Loc: jobsUrl, Priority: `1`},
		{Loc: baseUrl + `/contact`, Priority: `0.9`},
		{Loc: blogsUrl, Priority: `0.8`},
	}

	if strings.Contains(baseUrl, "dev-stack") {
		careers, err := repository.careerRepo.GetCareers()

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		for i := range careers {
			sitemapUrls = append(sitemapUrls, sitemap.URL{Loc: jobsUrl + `/` + careers[i].UniqueId, Priority: `0.9`})
			sitemapUrls = append(sitemapUrls, sitemap.URL{Loc: jobsUrl + `/apply/` + careers[i].UniqueId, Priority: `0.9`})
		}
	}

	//get first day of current month
	year, month, _ := time.Now().Date()
	lastmod := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T00:00:00.000Z")

	// generate sitemap
	sitemapXML := sitemap.New(sitemap.Options{
		Dir:     "vue-frontend/public",
		BaseURL: baseUrl,
	})

	sitemapXML.Open()
	for i := range sitemapUrls {
		sitemapUrls[i].ChangeFreq = "monthly"
		sitemapUrls[i].LastMod = lastmod
		sitemapXML.Add(sitemapUrls[i])
	}
	sitemapXML.Close()
}
