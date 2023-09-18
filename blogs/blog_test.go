package blogs

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Unique_Int_Success(t *testing.T) {
	got := utils.Unique([]int{1, 1, 2, 2, 3, 3})
	expected := []int{1, 2, 3}
	assert.Equal(t, got, expected)
}

func Test_Unique_String_Fail(t *testing.T) {
	got := utils.Unique([]string{"apple", "has", "apple", "logo"})
	expected := []string{"apple", "has", "apple", "logo"}
	assert.NotEqual(t, got, expected)
}

func Test_Unique_Struct_Success(t *testing.T) {
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

func Test_ReadFile_Empty_Filepath(t *testing.T) {
	got := utils.ReadSliceFromFile("", []Item{})
	expected := []Item{}
	assert.Equal(t, got, expected)
}

func Test_ReadFile_Wrong_Filepath(t *testing.T) {
	got := utils.ReadSliceFromFile("test.json", []Item{})
	expected := []Item{}
	assert.Equal(t, got, expected)
}

func Test_GetBlogs_Integration(t *testing.T) {
	engine := gin.New()
	engine.GET("/api/blogs", Get)
	req, err := http.NewRequest("GET", "/api/blogs", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
