package jobs

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type LifePerksImages struct {
	Id        int    `json:"id"`
	ImageUrls string `json:"image_urls"`
	Index     int    `json:"index"`
}

type LifeandPerksRepository struct {
	Db *sqlx.DB
}

func NewLifePerksImages(db *sqlx.DB) *LifeandPerksRepository {
	return &LifeandPerksRepository{Db: db}
}

func (repository *LifeandPerksRepository) LifeImages(c *gin.Context) {
	lifeImages := []LifePerksImages{}
	err := repository.Db.Select(&lifeImages, "SELECT l.id, l.image_urls, l.index FROM lifeatcanopas l")
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, lifeImages)
}

func (repository *LifeandPerksRepository) PerksImages(c *gin.Context) {
	perksImages := []LifePerksImages{}
	err := repository.Db.Select(&perksImages, "SELECT p.id, p.image_urls, p.index FROM perks p")
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, perksImages)
}
