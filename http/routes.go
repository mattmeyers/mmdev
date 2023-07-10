package http

import (
	"io"
	"net/http"

	"github.com/mattmeyers/mmdev/app"
)

func (s *Server) loadRoutes() {
	s.Router.Get("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/css")
		f, err := s.Resources.Open("style.css")
		if err != nil {
			panic(err)
		}

		io.Copy(w, f)
	})
	s.Router.Get("/", s.handleAppIndex())
	s.Router.Get("/about", s.handleAppAbout())
	s.Router.Get("/projects", s.handleAppProjects())

	s.Router.Handle("/*", http.FileServer(http.FS(app.Favicons)))
}
