package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (repository *Repository) GetPostsByTag(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		log.Warn("Slug not found in query paramete")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	isPublished, err := strconv.ParseBool(c.DefaultQuery("is_published", "false"))
	if err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	query := ""
	if isPublished {
		query = "is_published = true AND published_at IS NOT NULL AND"
	}

	posts := []Post{}
	err = repository.Db.Select(&posts, `SELECT id, title, slug, tags as tag, is_featured, is_published, published_on, published_at, reading_time, summary 
										FROM posts
										WHERE `+query+` tags @> $1`, `[{"slug": "`+slug+`"}]`)
	if err != nil {
		log.Error("Error while fetching posts: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	new_posts := []Post{}
	for _, post := range posts {

		var tags []Tag

		if err := json.Unmarshal([]byte(post.Tag), &tags); err != nil {
			log.Error("Error while unmarshal tags: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		post.Tag = ""
		post.Tags = tags

		new_posts = append(new_posts, post)
	}

	new_posts = repository.PreparePosts(new_posts)

	c.JSON(http.StatusOK, new_posts)
}
