package server

import (
	"fmt"
	"github.com/PaluMacil/dan2/www/projects"
	"net/http"

	"github.com/PaluMacil/dan2/www"
	"github.com/a-h/templ"
)

func routes() (*http.ServeMux, error) {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(www.StaticFS)))
	projectConfig, err := projects.Get()
	if err != nil {
		return nil, fmt.Errorf("getting project config: %w", err)
	}
	mux.Handle("/", templ.Handler(www.Base(projectConfig.ProjectList)))

	return mux, nil
}
