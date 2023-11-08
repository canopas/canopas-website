package jobs

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var l_repo *LifeandPerksRepository
var l_err error
var l_testDB *sqlx.DB

func TestInitLifePerksImages(t *testing.T) {
	l_testDB, l_err = utils.TestDB()
	if l_err != nil {
		t.Errorf("Error in initializing test DB: %v", l_err)
	}

	l_repo = NewLifePerksImages(l_testDB)
}

func TestGetAllLifeImages(t *testing.T) {
	utils.TruncateTables(l_testDB)
	l_err = utils.CreateTables(l_testDB)
	if l_err != nil {
		t.Errorf("Error in initializing test DB: %v", l_err)
	}
	utils.PrepareTablesData(l_testDB)
	engine := gin.New()
	engine.GET("/api/lifeimages", l_repo.LifeImages)
	req, l_err := http.NewRequest("GET", "/api/lifeimages", nil)
	if l_err != nil {
		t.Errorf("Error in creating request: %v", l_err)
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotArrayData(w, t)
	expected := []interface{}{expectedLifeImages()}
	assert.Equal(t, expected, got)
}

func TestGetAllPerksImages(t *testing.T) {
	utils.TruncateTables(l_testDB)
	l_err = utils.CreateTables(l_testDB)
	if l_err != nil {
		t.Errorf("Error in initializing test DB: %v", l_err)
	}
	utils.PrepareTablesData(l_testDB)
	engine := gin.New()
	engine.GET("/api/perkimages", l_repo.PerksImages)
	req, l_err := http.NewRequest("GET", "/api/perkimages", nil)
	if l_err != nil {
		t.Errorf("Error in creating request: %v", l_err)
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotArrayData(w, t)
	expected := []interface{}{expectedPerkImages()}
	assert.Equal(t, expected, got)
}
func expectedLifeImages() map[string]interface{} {
	lifeImages := make(map[string]interface{})
	lifeImages["id"] = 1.0
	lifeImages["image_urls"] = "https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-400.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-800.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-1600.webp"
	lifeImages["index"] = 0.0
	return lifeImages
}
func expectedPerkImages() map[string]interface{} {
	perkImages := make(map[string]interface{})
	perkImages["id"] = 1.0
	perkImages["image_urls"] = "https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-400.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-800.webp"
	perkImages["index"] = 0.0
	return perkImages
}
