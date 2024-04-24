package post

import (
	"time"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
)

type Repository struct {
	Db *sqlx.DB
}

type Post struct {
	Id               int               `json:"id"`
	Title            string            `json:"title"`
	Content          string            `json:"content,omitempty"`
	Slug             string            `json:"slug"`
	PublishedOn      time.Time         `json:"published_on,omitempty"`
	IsFeatured       bool              `json:"is_featured"`
	CreatedAt        string            `json:"created_at,omitempty"`
	UpdatedAt        string            `json:"updated_at,omitempty"`
	PublishedAt      null.String       `json:"published_at"`
	Summary          string            `json:"summary,omitempty"`
	BlogContent      string            `json:"blog_content,omitempty"`
	MetaDescription  string            `json:"meta_description,omitempty"`
	Toc              string            `json:"toc,omitempty"`
	Tag              string            `json:"tag,omitempty"`
	Tags             []Tag             `json:"tags,omitempty"`
	IsPublished      bool              `json:"is_published"`
	Keywords         *null.String      `json:"keywords,omitempty"`
	NewContent       *null.String      `json:"new_content,omitempty"`
	NewToc           *null.String      `json:"new_toc,omitempty"`
	NewBlogContent   *null.String      `json:"new_blog_content,omitempty"`
	IsResource       bool              `json:"is_resource"`
	ReadingTime      *null.Int         `json:"reading_time,omitempty"`
	Author           Author            `json:"author"`
	Image            Image             `json:"image"`
	Cta              Cta               `json:"cta,omitempty"`
	RecommendedPosts []RecommendedPost `json:"recommended_posts,omitempty"`
}

type Tag struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Author struct {
	Id        int    `json:"id"`
	RelatedId int    `json:"related_id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Bio       string `json:"bio"`
	Image     Image  `json:"image"`
}

type Image struct {
	Id               int         `json:"id"`
	RelatedId        int         `json:"related_id"`
	Name             string      `json:"name"`
	AlternativeText  null.String `json:"alternative_text"`
	Caption          null.String `json:"caption"`
	Height           int         `json:"height"`
	Width            int         `json:"width"`
	Format           string      `json:"format,omitempty"`
	Formats          Formats     `json:"formats"`
	Hash             string      `json:"hash"`
	Ext              string      `json:"ext"`
	Mime             string      `json:"mime"`
	Size             float32     `json:"size"`
	Url              string      `json:"url"`
	PreviewUrl       null.String `json:"preview_url"`
	Provider         string      `json:"provider"`
	ProviderMetadata null.String `json:"provider_metadata"`
	FolderPath       string      `json:"folder_path"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
}

type Formats struct {
	Large     *Format `json:"large,omitempty"`
	Small     *Format `json:"small,omitempty"`
	Medium    *Format `json:"medium,omitempty"`
	Thumbnail Format  `json:"thumbnail"`
}

type Format struct {
	Ext    string      `json:"ext"`
	Url    string      `json:"url"`
	Hash   string      `json:"hash"`
	Mime   string      `json:"mime"`
	Name   string      `json:"name"`
	Path   null.String `json:"path"`
	Size   float32     `json:"size"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
}

type Cta struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	ComponentName string `json:"component_name"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type RecommendedPost struct {
	Id              int         `json:"id"`
	Title           string      `json:"title"`
	Content         string      `json:"content"`
	Slug            string      `json:"slug"`
	PublishedOn     time.Time   `json:"published_on"`
	IsFeatured      bool        `json:"is_featured"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	PublishedAt     null.String `json:"published_at"`
	Summary         string      `json:"summary"`
	BlogContent     string      `json:"blog_content"`
	MetaDescription string      `json:"meta_description"`
	Toc             string      `json:"toc"`
	Tag             string      `json:"tag,omitempty"`
	Tags            []Tag       `json:"tags"`
	IsPublished     bool        `json:"is_published"`
	Keywords        null.String `json:"keywords"`
	NewContent      null.String `json:"new_content"`
	NewToc          null.String `json:"new_toc"`
	NewBlogContent  null.String `json:"new_blog_content"`
	IsResource      bool        `json:"is_resource"`
	ReadingTime     null.Int    `json:"reading_time"`
	Author          Author      `json:"author"`
	Image           Image       `json:"image"`
}
