package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (repository *Repository) GetPostsByAuthor(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		log.Warn("Username not found in query parameter")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	skip, err := strconv.Atoi(c.DefaultQuery("skip", "0"))
	if err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		log.Warn(err)
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
		query = " WHERE p.is_published = true AND published_at IS NOT NULL "
	}

	posts := []Post{}
	err = repository.Db.Select(&posts, `SELECT p.id, p.title, p.slug, p.tags as tag, p.is_featured, p.is_published, p.published_on, p.published_at, p.reading_time, p.summary 
										FROM posts p
										JOIN posts_author_links pa ON p.id = pa.post_id
										JOIN authors a ON a.id = pa.author_id AND a.username = $1`+query+`
										ORDER BY p.published_on DESC
										LIMIT $2 OFFSET $3`, username, limit, skip)
	if err != nil {
		log.Error("Error while fetching posts: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var PostData struct {
		Posts []Post `json:"posts"`
		Count int    `json:"count"`
	}

	new_post := PostData
	for _, post := range posts {

		var tags []Tag

		if err := json.Unmarshal([]byte(post.Tag), &tags); err != nil {
			log.Error("Error while unmarshal tags: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		post.Tag = ""
		post.Tags = tags

		new_post.Posts = append(new_post.Posts, post)
	}

	new_post.Posts = repository.PreparePosts(new_post.Posts)

	if skip == 0 {
		publishQuery := ""
		if isPublished {
			publishQuery = "WHERE p.is_published = true AND published_at IS NOT NULL"
		}

		err = repository.Db.Get(&new_post.Count, `SELECT COUNT(p.id) 
												FROM posts p 
												JOIN posts_author_links pa ON p.id = pa.post_id 
												JOIN authors a ON a.id = pa.author_id AND a.username = $1 `+publishQuery, username)
	}

	c.JSON(http.StatusOK, new_post)
}

func (repository *Repository) GetAuthorByPostId(id int) (error, Author) {
	author := Author{}

	err := repository.Db.Get(&author, `SELECT a.id, a.username, a.name, a.email, a.created_at, a.updated_at, a.bio 
									FROM authors a
									JOIN posts_author_links pa ON pa.post_id = $1 AND pa.author_id = a.id`, id)
	if err != nil {
		log.Error("Error while fetching author: ", err)
		return err, author
	}

	err = repository.Db.Get(&author.Image, `SELECT f.id, f.name, f.alternative_text, f.caption, f.width, f.height, f.formats as format, f.hash, f.ext, f.mime, 
												f.size, f.url, f.preview_url, f.provider, f.provider_metadata, f.folder_path, f.created_at, f.updated_at 
												FROM files f 
												JOIN files_related_morphs frm ON frm.related_id = $1 AND frm.related_type = 'api::author.author' AND frm.file_id = f.id`, author.Id)
	if err != nil {
		log.Error("Error while fetching author image: ", err)
		return err, author
	}

	var authorFormats Formats

	if err := json.Unmarshal([]byte(author.Image.Format), &authorFormats); err != nil {
		log.Error("Error while unmarshal author image formats: ", err)
		return err, author
	}

	author.Image.Format = ""
	author.Image.Formats = authorFormats

	return nil, author
}

func (repository *Repository) SetAuthorWithImageInPosts(posts []Post) []Post {

	postIds := []string{}
	for _, post := range posts {
		postIds = append(postIds, strconv.Itoa(post.Id))
	}

	new_post_ids := strings.Join(postIds, ", ")

	authors := []Author{}
	authorQuery := fmt.Sprintf(`SELECT a.id, pa.post_id as related_id, a.username, a.name, a.email, a.created_at, a.updated_at, a.bio 
									FROM authors a
									JOIN posts_author_links pa ON pa.post_id IN (%s) AND pa.author_id = a.id`, new_post_ids)

	err := repository.Db.Select(&authors, authorQuery)
	if err != nil {
		log.Error("Error while fetching author: ", err)
		return posts
	}

	authorIds := []string{}
	for _, author := range authors {
		authorIds = append(authorIds, strconv.Itoa(author.Id))
	}

	new_author_ids := strings.Join(authorIds, ", ")

	images := []Image{}
	authorImageQuery := fmt.Sprintf(`SELECT f.id, frm.related_id as related_id, f.name, f.alternative_text, f.caption, f.width, f.height, f.formats as format, f.hash, f.ext, f.mime, 
						f.size, f.url, f.preview_url, f.provider, f.provider_metadata, f.folder_path, f.created_at, f.updated_at 
						FROM files f 
						JOIN files_related_morphs frm ON frm.related_id IN (%s) AND frm.related_type = 'api::author.author' AND frm.file_id = f.id`, new_author_ids)

	err = repository.Db.Select(&images, authorImageQuery)
	if err != nil {
		log.Error("Error while fetching post images: ", err)
		return posts
	}

	// set author image in author
	for i, author := range authors {
		for _, image := range images {
			if author.Id == image.RelatedId {

				authors[i].Image = image

				// set format of author image if exists
				if authors[i].Image.Format != "" {

					var authorFormats Formats
					if err := json.Unmarshal([]byte(authors[i].Image.Format), &authorFormats); err != nil {
						log.Error("Error while unmarshal post image formats: ", err)
						return posts
					}

					authors[i].Image.Format = ""
					authors[i].Image.Formats = authorFormats
				}
				break
			}
		}
	}

	// set author with image in post
	for i, post := range posts {
		for _, author := range authors {
			if post.Id == author.RelatedId {
				posts[i].Author = author
				break
			}
		}
	}

	return posts
}
