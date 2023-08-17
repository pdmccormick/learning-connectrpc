package frontend

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
)

//go:embed all:build
var RawBuildAssets embed.FS

var (
	BuildAssets, _       = fs.Sub(RawBuildAssets, "build")
	_              fs.FS = BuildAssets
)

// Handler
type Handler struct {
	PathOverride string

	h http.Handler
}

func (fe *Handler) Init() error {
	var fsys = BuildAssets

	if v := fe.PathOverride; v != "" {
		fsys = os.DirFS(v)
	}

	var h = http.FileServer(http.FS(fsys))

	*fe = Handler{
		PathOverride: fe.PathOverride,

		h: h,
	}

	return nil
}

func (fe *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) { fe.h.ServeHTTP(w, r) }
