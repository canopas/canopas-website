package post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

func New(db *sqlx.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Get(c *gin.Context) {
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

	isResource, err := strconv.ParseBool(c.DefaultQuery("is_resource", "false"))
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

	var PostData struct {
		Posts         []Post `json:"posts"`
		FeaturedPosts []Post `json:"featuredPosts"`
		Count         int    `json:"count"`
	}

	new_posts := PostData
	err, new_posts.Posts = repository.GetPosts(false, isResource, isPublished, limit, skip)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if isResource && skip == 0 {
		err, new_posts.FeaturedPosts = repository.GetPosts(true, isResource, isPublished, 6, 0)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	resourceQuery := ""
	if isResource {
		resourceQuery = "is_resource = true AND"
	}

	publishQuery := ""
	if isPublished {
		publishQuery = "AND is_published = true"
	}

	err = repository.Db.Get(&new_posts.Count, `SELECT COUNT(id) FROM posts WHERE `+resourceQuery+` is_featured = false `+publishQuery)

	c.JSON(http.StatusOK, new_posts)
}

func (repository *Repository) GetPosts(isFeatured, isResource, isPublished bool, limit, skip int) (error, []Post) {

	featuredQuery := ""
	if isFeatured {
		featuredQuery = " AND is_featured = true"
	}

	publishQuery := ""
	if isPublished {
		publishQuery = " AND is_published = true"
	}

	posts := []Post{}
	err := repository.Db.Select(&posts, `SELECT id, title, slug, is_featured, is_published, published_on, published_at, reading_time, summary  
										FROM posts 
										WHERE is_resource = $1`+publishQuery+featuredQuery+` 
										ORDER BY published_on DESC
										LIMIT $2 OFFSET $3`, isResource, limit, skip)
	if err != nil {
		log.Error(err)
		return err, nil
	}

	new_posts := []Post{}
	new_posts = repository.PreparePosts(posts)

	return nil, new_posts
}

func (repository *Repository) Show(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		log.Warn("Slug not found in query paramete")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	post := Post{}
	err := repository.Db.Get(&post, `SELECT id, title, content, slug, published_on, is_featured, created_at, updated_at, published_at, 
									summary, blog_content, meta_description, toc, tags as tag, is_published, keywords, new_content, new_toc, new_blog_content, 
									is_resource, reading_time 
									FROM posts 
									WHERE slug = $1`, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Error("Error while fetching post: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var tags []Tag

	if err := json.Unmarshal([]byte(post.Tag), &tags); err != nil {
		log.Error("Error while unmarshal tags: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	post.Tag = ""
	post.Tags = tags

	err = repository.Db.Get(&post.Cta, `SELECT c.id, c.name, c.component_name, c.created_at, c.updated_at 
									FROM ctas c
									JOIN posts_cta_links pc ON pc.post_id = $1 AND pc.cta_id = c.id`, post.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("No CTA found")
			post.Cta = Cta{}
		} else {
			log.Error("Error while fetching cta: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	if post.IsResource {
		posts := []RecommendedPost{}
		err = repository.Db.Select(&posts, `SELECT id, title, content, slug, published_on, is_featured, created_at, updated_at, published_at,
											summary, blog_content, meta_description, toc, tags as tag, is_published, keywords, new_content, new_toc, new_blog_content,
											is_resource, reading_time 
											FROM posts 
											WHERE slug != $1`, slug)
		if err != nil {
			log.Error("Error while fetching recommended post: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		post.RecommendedPosts = repository.FilterPostsByTags(posts, repository.ExtractTagNames(post.Tags))
	}

	new_posts := repository.PreparePosts([]Post{post})

	c.JSON(http.StatusOK, new_posts[0])
}

func (repository *Repository) PreparePosts(posts []Post) []Post {

	new_posts := repository.SetPostImageInPosts(posts)
	new_posts = repository.SetAuthorWithImageInPosts(posts)

	return new_posts
}

func (repository *Repository) SetPostImageInPosts(posts []Post) []Post {

	postIds := []string{}
	for _, post := range posts {
		postIds = append(postIds, strconv.Itoa(post.Id))
	}

	new_post_ids := strings.Join(postIds, ", ")

	images := []Image{}
	postImageQuery := fmt.Sprintf(`SELECT f.id, frm.related_id as related_id, f.name, f.alternative_text, f.caption, f.width, f.height, f.formats as format, f.hash, f.ext, f.mime, 
						f.size, f.url, f.preview_url, f.provider, f.provider_metadata, f.folder_path, f.created_at, f.updated_at 
						FROM files f 
						JOIN files_related_morphs frm ON frm.related_id IN (%s) AND frm.related_type = 'api::post.post' AND frm.file_id = f.id`, new_post_ids)

	err := repository.Db.Select(&images, postImageQuery)
	if err != nil {
		log.Error("Error while fetching post images: ", err)
		return posts
	}

	for i, post := range posts {
		for _, image := range images {
			if post.Id == image.RelatedId {
				// set image in post
				posts[i].Image = image

				// set format of image if exists
				if posts[i].Image.Format != "" {

					var postFormats Formats
					if err := json.Unmarshal([]byte(posts[i].Image.Format), &postFormats); err != nil {
						log.Error("Error while unmarshal post image formats: ", err)
						return posts
					}

					posts[i].Image.Format = ""
					posts[i].Image.Formats = postFormats
				}
				break
			}
		}
	}

	return posts
}

// Filter posts based on tags
func (repository *Repository) FilterPostsByTags(posts []RecommendedPost, tags []string) []RecommendedPost {

	filteredPosts := []RecommendedPost{}
	count := 0
	var err error

	for _, post := range posts {
		if count >= 3 {
			break
		}

		var postTags []Tag

		if err = json.Unmarshal([]byte(post.Tag), &postTags); err != nil {
			log.Error("Error while unmarshal tags: ", err)
			return nil
		}

		post.Tags = postTags

		for _, tag := range post.Tags {
			if lo.Contains(tags, tag.Name) {

				err, post.Author = repository.GetAuthorByPostId(post.Id)
				if err != nil {
					return nil
				}

				filteredPosts = append(filteredPosts, post)
				count++

				break
			}
		}
	}

	return filteredPosts
}

// Helper function to extract tag names from a slice of tags
func (repository *Repository) ExtractTagNames(tags []Tag) []string {
	tagNames := make([]string, len(tags))
	for i, tag := range tags {
		tagNames[i] = tag.Name
	}
	return tagNames
}
