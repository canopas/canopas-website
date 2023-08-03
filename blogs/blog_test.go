package blogs

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUniqueIntSuccess(t *testing.T) {
	got := utils.Unique([]int{1, 1, 2, 2, 3, 3})
	expected := []int{1, 2, 3}
	assert.Equal(t, got, expected)
}

func TestUniqueStringFail(t *testing.T) {
	got := utils.Unique([]string{"apple", "has", "apple", "logo"})
	expected := []string{"apple", "has", "apple", "logo"}
	assert.NotEqual(t, got, expected)
}

func TestUniqueStructSuccess(t *testing.T) {
	got := utils.Unique([]Item{{
		Title: "test1",
	}, {
		Title: "test1",
	}, {
		Title: "test2",
	}})
	expected := []Item{{
		Title: "test1",
	}, {
		Title: "test2",
	}}
	assert.Equal(t, got, expected)
}

func TestReadFileEmptyFilepath(t *testing.T) {
	got := utils.ReadSliceFromFile("", []Item{})
	expected := []Item{}
	assert.Equal(t, got, expected)
}

func TestReadFileWrongFilepath(t *testing.T) {
	got := utils.ReadSliceFromFile("test.json", []Item{})
	expected := []Item{}
	assert.Equal(t, got, expected)
}

func TestGetBlogsIntegration(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/blogs", nil)

	assert.NoError(t, err)

	engine := gin.New()
	setUpRouter(engine)
	engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func setUpRouter(engine *gin.Engine) {
	engine.GET("/api/blogs", Get)
}
