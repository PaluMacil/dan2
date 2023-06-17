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
	// TODO: use template composition and children
	// - https://github.com/a-h/templ/blob/acea959ba5020adc62f477c0fc8610b17c7aa253/examples/blog/main.go
	// -
	mux.Handle("/", templ.Handler(www.Base(projectConfig.ProjectList)))

	return mux, nil
}
