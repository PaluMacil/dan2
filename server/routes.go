package server

import (
	"fmt"
	"github.com/PaluMacil/dan2/www/projects"
	"log"
	"net/http"

	"github.com/PaluMacil/dan2/www"
	"github.com/a-h/templ"
)

func Serve() error {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(www.StaticFS)))
	projectConfig, err := projects.Get()
	if err != nil {
		return fmt.Errorf("getting project config: %w", err)
	}
	mux.Handle("/", templ.Handler(www.Base(projectConfig.ProjectList)))

	log.Print("Listening on http://localhost:3000")

	http.ListenAndServe(":3000", mux)

	return nil
}
