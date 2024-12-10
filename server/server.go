package server

import (
	"fmt"
	"net/http"

	"github.com/ekediala/expensix"
	"github.com/ekediala/expensix/httpio"
	"github.com/ekediala/expensix/store"
)

var (
	PathLanding = fmt.Sprintf("%s %s", http.MethodGet, expensix.RouteLanding)
)

type Server struct {
	http.Handler
}

func New(store *store.Store) *Server {
	mux := http.NewServeMux()
	mux.Handle(PathLanding, Landing())

	s := Server{
		Handler: mux,
	}
	return &s
}

func Landing() httpio.Handler {
	return func(w http.ResponseWriter, r *http.Request) httpio.Handler {
		return httpio.Text(http.StatusText(http.StatusOK))
	}
}
