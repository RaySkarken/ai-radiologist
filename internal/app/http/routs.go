package http

import (
	"ai-radiologist/internal/app/http/handlers"
	"ai-radiologist/internal/views"
	"errors"
	"net/http"
)

func NewRouter(handlers *handlers.Handlers) (*http.ServeMux, error) {
	if handlers == nil {
		return nil, errors.New("handler is nil")
	}
	mux := http.NewServeMux()

	// auth
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err := views.Index().Render(r.Context(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	return mux, nil
}
