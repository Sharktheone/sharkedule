package web

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed dist
	embedFS embed.FS
)

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

func Serve(r *gin.Engine) {
	subFs, err := fs.Sub(embedFS, "dist")
	if err != nil {
		log.Fatalf("Failed to get subFS: %v", err)
	}

	FS := &embedFileSystem{
		FileSystem: http.FS(subFs),
	}

	r.Use(static.Serve("/", FS))
}
