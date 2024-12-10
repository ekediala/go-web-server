package httpio

import (
	"fmt"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) Handler

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if next := h(w, r); next != nil {
		next(w, r)
	}
}

func OK(w http.ResponseWriter, r *http.Request) Handler {
	return nil
}

func Text(s string) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, s)
		return OK
	}
}
