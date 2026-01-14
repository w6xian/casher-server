package frontend

import (
	"context"
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed dist/*
var embeddedFiles embed.FS

type FrontendService struct {
	Path string
}

func NewFrontendService(path string) *FrontendService {
	return &FrontendService{
		Path: "/" + path + "/",
	}
}

func (fe *FrontendService) Serve(ctx context.Context, r *mux.Router) error {
	// 这个是嵌入
	prefix := fe.Path

	r.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(getFileSystem("dist"))))
	// // 这个f需要目录
	// r.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(http.Dir("./dist"))))
	return nil
}

// Register healthz endpoint.

func getFileSystem(path string) http.FileSystem {
	fs, err := fs.Sub(embeddedFiles, path)
	if err != nil {
		panic(err)
	}
	return http.FS(fs)
}
