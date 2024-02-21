package sitemap

import (
	"bytes"
	"embed"
	"encoding/json"
	"io"
	"jobs"
	"net/http"
	"os"
	"strings"
	"text/template"
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
	templates  *template.Template
}

func New(careerRepo *jobs.CareerRepository, templateFs embed.FS) *SitemapRepository {
	templates := template.Must(template.ParseFS(templateFs, "templates/*.txt"))
	return &SitemapRepository{careerRepo: careerRepo, templates: templates}
}

type URL struct {
	Loc      string `json:"loc"`
	LastMod  string `json:"lastmod"`
	Priority string `json:"priority"`
}

// resources structs
type Post struct {
	Slug        string `json:"slug"`
	PublishedOn string `json:"published_on"`
}

type Attributes struct {
	Posts     []Post `json:"posts"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"createdAt"`
	Username  string `json:"username"`
}

type Data struct {
	Attributes Attributes `json:"attributes"`
}

type Resources struct {
	Data Data `json:"data"`
}

type TagAndAuthors struct {
	Data []Data `json:"data"`
}

func (repository *SitemapRepository) GenerateSitemap(c *gin.Context) {
	baseUrl := c.Query("baseUrl")

	sitemapUrls := []URL{
		{Loc: baseUrl, Priority: `1`},
	}

	filePath := "templates/path.txt"
	sitemapUrls, err := repository.addPages(baseUrl, sitemapUrls, filePath)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//get first day of current month
	year, month, _ := time.Now().Date()
	lastmod := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T00:00:00.000")

	for i := range sitemapUrls {
		sitemapUrls[i].LastMod = lastmod
	}

	// add published resource blogs
	sitemapUrls, err = addPublishedResources(baseUrl, sitemapUrls)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, sitemapUrls)
}

func (repository *SitemapRepository) addPages(baseUrl string, sitemapUrls []URL, file string) ([]URL, error) {
	var templateBuffer bytes.Buffer
	var input string

	err := repository.templates.ExecuteTemplate(&templateBuffer, "path.txt", input)

	if err != nil {
		log.Error(err)
		return sitemapUrls, nil
	}

	fileContent := string(templateBuffer.String())
	lines := strings.Split(fileContent, "\n")

	// append all pages
	for _, fileName := range lines {
		sitemapUrls = append(sitemapUrls, URL{Loc: baseUrl + "/" + fileName, Priority: "0.9"})
	}

	// append jobs
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

	// append portfolios
	portfolioUrl := baseUrl + "/portfolio"
	portfolios := []string{"luxeradio", "togness", "justly"}
	for i := range portfolios {
		sitemapUrls = append(sitemapUrls, URL{Loc: portfolioUrl + `/` + portfolios[i], Priority: `0.9`})
	}

	return sitemapUrls, nil
}

func addPublishedResources(baseUrl string, sitemapUrls []URL) ([]URL, error) {

	resourceUrl := os.Getenv("RESOURCES_URL")

	if resourceUrl == "" {
		return sitemapUrls, nil
	}

	responseData, err := doRequest(resourceUrl + "/v1/posts?populate=deep&publicationState=live&fields[0]=slug&fields[1]=published_on")

	if err != nil {
		log.Error(err)
		return sitemapUrls, err
	}

	var resources Resources
	json.Unmarshal(responseData, &resources)
	for i := range resources.Data.Attributes.Posts {
		data := resources.Data.Attributes.Posts[i]
		sitemapUrls = append(sitemapUrls, URL{Loc: baseUrl + `/` + data.Slug, Priority: `0.9`, LastMod: data.PublishedOn})
	}

	// get tags
	responseTags, err := doRequest(resourceUrl + "/v1/tags?fields[0]=slug&fields[1]=createdAt")

	if err != nil {
		log.Error(err)
		return sitemapUrls, err
	}

	var tags TagAndAuthors
	json.Unmarshal(responseTags, &tags)

	for i := range tags.Data {
		data := tags.Data[i]
		sitemapUrls = append(sitemapUrls, URL{Loc: baseUrl + `/tag/` + data.Attributes.Slug, Priority: `0.8`, LastMod: data.Attributes.CreatedAt})
	}

	// get authors
	responseAuthors, err := doRequest(resourceUrl + "/v1/authors?fields[0]=username")

	if err != nil {
		log.Error(err)
		return sitemapUrls, err
	}

	var authors TagAndAuthors
	json.Unmarshal(responseAuthors, &authors)

	for i := range authors.Data {
		data := authors.Data[i]
		sitemapUrls = append(sitemapUrls, URL{Loc: baseUrl + `/author/` + data.Attributes.Username, Priority: `0.8`})
	}

	return sitemapUrls, nil
}

func doRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return responseData, nil
}
