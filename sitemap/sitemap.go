package sitemap

import (
	"encoding/xml"
	"jobs"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

const (
	XMLNS       = "http://www.sitemaps.org/schemas/sitemap/0.9"
	XMLNS_VIDEO = "http://www.google.com/schemas/sitemap-video/1.1"
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
	Video      []Video  `xml:"video:video;omitempty"`
}

type Video struct {
	XMLName      xml.Name `xml:"video:video"`
	Title        string   `xml:"video:title"`
	ContentLoc   string   `xml:"video:content_loc"`
	ThumbnailLoc string   `xml:"video:thumbnail_loc"`
}

type Portfolio struct {
	Name   string
	Videos []Video
}

type URLset struct {
	XMLName    xml.Name `xml:"urlset"`
	XMLNS      string   `xml:"xmlns,attr"`
	XMLNSVideo string   `xml:"xmlns:video,attr"`
	URL        []URL    `xml:"url"`
}

func (repository *SitemapRepository) GenerateSitemap(c *gin.Context) {
	baseUrl := c.Query("baseUrl")

	sitemapUrls := []URL{
		{Loc: baseUrl, Priority: `1`},
		{Loc: baseUrl + `/contact`, Priority: `0.9`},
		{Loc: baseUrl + `/about`, Priority: `0.9`},
		{Loc: baseUrl + `/resources`, Priority: `0.9`},
	}

	sitemapUrls = addPortfolios(baseUrl, sitemapUrls)

	sitemapUrls, err := repository.addJobs(baseUrl, sitemapUrls)

	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//get first day of current month
	year, month, _ := time.Now().Date()
	lastmod := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T00:00:00.000Z")

	for i := range sitemapUrls {
		sitemapUrls[i].ChangeFreq = "monthly"
		sitemapUrls[i].LastMod = lastmod
	}

	urlset := URLset{
		URL:        sitemapUrls,
		XMLNS:      XMLNS,
		XMLNSVideo: XMLNS_VIDEO,
	}

	c.Header("Content-Type", "application/xml")

	c.XML(http.StatusOK, urlset)
}

func (repository *SitemapRepository) addJobs(baseUrl string, sitemapUrls []URL) ([]URL, error) {
	jobsUrl := baseUrl + "/jobs"
	sitemapUrls = append(sitemapUrls, URL{Loc: jobsUrl, Priority: `1`})

	careers, err := repository.careerRepo.GetCareers()
	if err != nil {
		return sitemapUrls, err
	}

	for i := range careers {
		sitemapUrls = append(sitemapUrls, URL{Loc: jobsUrl + `/` + careers[i].UniqueId, Priority: `0.9`})
		sitemapUrls = append(sitemapUrls, URL{Loc: jobsUrl + `/apply/` + careers[i].UniqueId, Priority: `0.9`})
	}

	return sitemapUrls, nil
}

func addPortfolios(baseUrl string, sitemapUrls []URL) []URL {
	portfolioUrl := baseUrl + "/portfolio"

	portfolios := []Portfolio{
		{
			Name: "luxeradio",
			Videos: []Video{
				{
					Title:        "luxeradio video",
					ContentLoc:   baseUrl + "/videos/luxeradio_video.mp4",
					ThumbnailLoc: baseUrl + "/videos/luxeradio_video_thumbnail.webp",
				},
			},
		},
		{
			Name: "togness",
		},
		{
			Name: "justly",
		},
	}

	sitemapUrls = append(sitemapUrls, URL{Loc: portfolioUrl, Priority: `0.9`})

	for i := range portfolios {
		sitemapUrls = append(sitemapUrls, URL{Loc: portfolioUrl + `/` + portfolios[i].Name, Priority: `0.9`, Video: portfolios[i].Videos})
	}

	return sitemapUrls
}
