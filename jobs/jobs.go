package jobs

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
)

type Career struct {
	Id               int         `json:"id"`
	Title            string      `json:"title"`
	Summary          string      `json:"summary"`
	Description      string      `json:"description"`
	ButtonName       string      `json:"button_name"`
	Qualification    string      `json:"qualification"`
	EmploymentType   string      `json:"employment_type"`
	BaseSalary       int         `json:"base_salary"`
	Experience       string      `json:"experience"`
	IsActive         bool        `json:"is_active"`
	Skills           null.String `json:"skills"`
	TotalOpenings    int         `json:"total_openings"`
	Responsibilities null.String `json:"responsibilities"`
	IconName         string      `json:"icon_name"`
}

type CareerRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *CareerRepository {
	return &CareerRepository{Db: db}
}

func (repository *CareerRepository) Careers(c *gin.Context) {
	var careersList []Career

	err := repository.Db.Select(&careersList, `SELECT id, title, summary, description, button_name, qualification, employment_type, base_salary, experience, is_active, skills, total_openings, responsibilities, icon_name FROM jobs WHERE is_active = 1 `)

	if err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, careersList)

}

func (repository *CareerRepository) CareerById(c *gin.Context) {

	career := Career{}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = repository.Db.Get(&career, `SELECT id, title, summary, description, button_name, qualification, employment_type, base_salary, experience, is_active, skills, total_openings, responsibilities, icon_name FROM jobs
										WHERE id = ?
										AND is_active = 1`, id)

	if err != nil {
		log.Print(err)
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, career)
}
