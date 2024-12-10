package server

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/ekediala/template"
	"github.com/ekediala/template/httpio"
	"github.com/ekediala/template/store"
	"github.com/ekediala/template/templ/pages/health"
)

//go:embed public/*
var assets embed.FS

var (
	PathLanding = fmt.Sprintf("%s %s", http.MethodGet, template.RouteLanding)
	PathAssets  = fmt.Sprintf("%s /public/", http.MethodGet)
)

type Server struct {
	http.Handler
}

func New(store *store.Store) *Server {
	mux := http.NewServeMux()
	mux.Handle(PathLanding, Landing())
	mux.Handle(PathAssets, http.FileServer(http.FS(assets)))

	s := Server{
		Handler: mux,
	}
	return &s
}

func Landing() httpio.Handler {
	return func(w http.ResponseWriter, r *http.Request) httpio.Handler {
		return httpio.HTML(health.Health(), http.StatusOK)
	}
}
