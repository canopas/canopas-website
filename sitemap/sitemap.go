package sitemap

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"jobs"
	"net/http"
	"os"
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

// resources structs
type Attributes struct {
	Slug string `json:"slug"`
}

type Data struct {
	Attributes Attributes `json:"attributes"`
}

type Resources struct {
	Data []Data `json:"data"`
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
	}

	// add portfolios
	sitemapUrls = addPortfolios(baseUrl, sitemapUrls)

	// add all jobs
	sitemapUrls, err := repository.addJobs(baseUrl, sitemapUrls)

	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// add published resource blogs
	sitemapUrls, err = addPublishedResources(baseUrl, sitemapUrls)

	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//add contribution
	sitemapUrls = addcontributions(baseUrl, sitemapUrls)

	//add service
	sitemapUrls = addServices(baseUrl, sitemapUrls)

	//add client-thankyou
	sitemapUrls = addClientThankyou(baseUrl, sitemapUrls)

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
		sitemapUrls = append(sitemapUrls, URL{Loc: jobsUrl + `/thank-you`, Priority: `0.9`})
	}

	return sitemapUrls, nil
}
func addcontributions(baseUrl string, sitemapUrls []URL) []URL {
	contributionUrl := baseUrl + "/contributions"

	sitemapUrls = append(sitemapUrls, URL{Loc: contributionUrl, Priority: `0.9`})

	return sitemapUrls
}
func addServices(baseUrl string, sitemapUrls []URL) []URL {
	serviceUrl := baseUrl + "/services"

	sitemapUrls = append(sitemapUrls, URL{Loc: serviceUrl, Priority: `0.9`})

	return sitemapUrls
}
func addClientThankyou(baseUrl string, sitemapUrls []URL) []URL {
	clientThankyouUrl := baseUrl + "/thank-you"

	sitemapUrls = append(sitemapUrls, URL{Loc: clientThankyouUrl, Priority: `0.9`})

	return sitemapUrls
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

func addPublishedResources(baseUrl string, sitemapUrls []URL) ([]URL, error) {
	resourcesUrl := baseUrl + "/resources"
	sitemapUrls = append(sitemapUrls, URL{Loc: resourcesUrl, Priority: `0.9`})

	resourceUrl := os.Getenv("RESOURCES_URL")

	if resourceUrl == "" {
		return sitemapUrls, nil
	}

	req, err := http.NewRequest("GET", resourceUrl+"/v1/posts?populate=deep&publicationState=live", nil)
	if err != nil {
		log.Error(err)
		return sitemapUrls, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		return sitemapUrls, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return sitemapUrls, err
	}

	var resources Resources
	json.Unmarshal(responseData, &resources)

	for i := range resources.Data {
		data := resources.Data[i]
		sitemapUrls = append(sitemapUrls, URL{Loc: resourcesUrl + `/` + data.Attributes.Slug, Priority: `0.9`})
	}

	return sitemapUrls, nil
}
