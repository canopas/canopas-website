package sitemap

import (
	"encoding/xml"
	"jobs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SitemapRepository struct {
	careerRepo *jobs.CareerRepository
}

func New(careerRepo *jobs.CareerRepository) *SitemapRepository {
	return &SitemapRepository{careerRepo: careerRepo}
}

type URL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	ChangeFreq string   `xml:"changefreq"`
	LastMod    string   `xml:"lastmod"`
	Priority   string   `xml:"priority"`
}

type URLset struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URL     []URL
}

func (repository *SitemapRepository) GenerateSitemap(c *gin.Context) {
	baseUrl := c.Query("baseUrl")
	jobsUrl := baseUrl + "/jobs"
	blogsUrl := "https://blog.canopas.com"

	sitemapUrls := []URL{
		{Loc: baseUrl, Priority: `1`},
		{Loc: jobsUrl, Priority: `1`},
		{Loc: baseUrl + `/contact`, Priority: `0.9`},
		{Loc: blogsUrl, Priority: `0.8`},
	}

	careers, err := repository.careerRepo.GetCareers()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for i := range careers {
		sitemapUrls = append(sitemapUrls, URL{Loc: jobsUrl + `/` + careers[i].UniqueId, Priority: `0.9`})
		sitemapUrls = append(sitemapUrls, URL{Loc: jobsUrl + `/apply/` + careers[i].UniqueId, Priority: `0.9`})
	}

	//get first day of current month
	year, month, _ := time.Now().Date()
	lastmod := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T00:00:00.000Z")

	for i := range sitemapUrls {
		sitemapUrls[i].ChangeFreq = "monthly"
		sitemapUrls[i].LastMod = lastmod
	}

	urlset := URLset{URL: sitemapUrls, XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	c.Header("Content-Type", "application/xml")

	c.XML(http.StatusOK, urlset)
}
