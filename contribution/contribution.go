package contribution

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Items []Item `json:"items"`
}

type Item struct {
	Name  string `json:"name"`
	Stars int    `json:"stargazers_count"`
}

func GetStargazers(c *gin.Context) {
	req, err := http.NewRequest("GET", "https://api.github.com/search/repositories?q=org:canopas+stars:>150", nil)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("PERSONAL_ACCESS_TOKEN_GITHUB"))
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var repos Response
	json.Unmarshal(responseData, &repos)

	c.JSON(http.StatusOK, repos)
}
