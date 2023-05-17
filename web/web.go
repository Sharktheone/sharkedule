package web

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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

func Serve(r *fiber.App) {
	subFs, err := fs.Sub(embedFS, "dist")
	if err != nil {
		log.Fatalf("Failed to get subFS: %v", err)
	}

	FS := &embedFileSystem{
		FileSystem: http.FS(subFs),
	}

	static := filesystem.New(filesystem.Config{
		Root:   FS,
		Index:  "index.html",
		Browse: true,
		MaxAge: 3600,
	})

	r.Use(static)
	r.Get("*", func(c *fiber.Ctx) error {
		return filesystem.SendFile(c, FS, "index.html")
	})
}
