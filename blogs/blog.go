package blogs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
	"utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const YYYYMMDD = "2006-01-02 15:04:05"
const MAX_BLOG_NUMBER = 3

type Item struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	PubDate   string `json:"pubDate"`
	Link      string `json:"link"`
	GUID      string `json:"guid"`
	Thumbnail string `json:"thumbnail"`
}

type Blog struct {
	Items []Item `json:"items"`
}

func Get(c *gin.Context) {

	// get blogs from API
	response, err := http.Get("https://api.rss2json.com/v1/api.json?rss_url=https://medium.com/feed/canopas")

	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// bind data to struct
	var blogs Blog
	json.Unmarshal(responseData, &blogs)

	filteredItems := []Item{}
	fileName := "./blogs.json"

	// filter weekly and newletteres
	for _, item := range blogs.Items {
		if !strings.Contains(strings.ToLower(item.Title), "weekly") && !strings.Contains(strings.ToLower(item.Title), "newsletter") {
			filteredItems = append(filteredItems, item)
		}
	}

	// read and append blog from file, if blogs are less then 3
	if len(filteredItems) < MAX_BLOG_NUMBER {
		filteredItems = append(filteredItems, utils.ReadSliceFromFile(fileName, []Item{})...)

		// make blogs unique and get only 3 blogs
		filteredItems = utils.Unique(filteredItems)
	}

	if len(filteredItems) > MAX_BLOG_NUMBER {
		filteredItems = filteredItems[0:MAX_BLOG_NUMBER]
	}

	// sort blogs by published date
	sortBlogs(filteredItems)

	// write blogs to file
	utils.WriteSliceToFile(fileName, filteredItems)

	c.JSON(http.StatusOK, filteredItems)
}

func sortBlogs(blogs []Item) {
	sort.Slice(blogs, func(i, j int) bool {
		dateI, err := time.Parse(YYYYMMDD, blogs[i].PubDate)
		if err != nil {
			log.Warn(err)
		}
		dateJ, err := time.Parse(YYYYMMDD, blogs[j].PubDate)
		if err != nil {
			log.Warn(err)
		}
		return dateI.After(dateJ)
	})
}
