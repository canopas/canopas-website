package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed vue-frontend/dist
var frontendStatic embed.FS

func main() {
    router := gin.Default()

	router.Use(static.Serve("/", EmbedFolder(frontendStatic, "vue-frontend/dist")))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	if inLambda() {
		log.Fatal(gateway.ListenAndServe(":8080", router))
	} else {
		router.Run()
	}
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
