package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func (s *Server) handleAppIndex() http.HandlerFunc {
	t := s.parseTemplate("templates/base.html", "templates/nav.html", "templates/index.html")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
	})
}

func (s *Server) handleAppAbout() http.HandlerFunc {
	t := s.parseTemplate("templates/base.html", "templates/nav.html", "templates/about.html")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "about.html", nil)
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
	})
}

type project struct {
	Title   string
	Summary string
	Version string
	Link    string
	Docs    *string
	Content template.HTML
}

var projects = []project{
	{
		Title:   "vscode-dbml",
		Summary: "Language support for the Database Markup Language",
		Version: "v0.3.4",
		Link:    "https://github.com/mattmeyers/vscode-dbml",
		Content: template.HTML(`
		<p>
			This Visual Studio Code extension provides language support for the Database Markup Language (<a target="_blank" href="https://dbml.org/">DBML</a>). Features include syntax highlighting, language snippets, and pass-through commands to the DBML CLI tool. These commands allow the user to convert between SQL and DBML from within VS Code.
		</p>
		<p>
			Future plans for this project includes integration with a custom DBML language server which is in the early stages of development.
		</p>
		`),
	},
}

func (s *Server) handleAppProjects() http.HandlerFunc {
	t := s.parseTemplate("templates/base.html", "templates/nav.html", "templates/projects.html")

	type params struct {
		Projects []project
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "projects.html", params{Projects: projects})
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
	})
}

func (s *Server) handleAppLogin() http.HandlerFunc {
	t := s.parseTemplate("templates/base.html", "templates/nav.html", "templates/login.html")

	type params struct {
		Projects []project
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
	})
}

func (s *Server) handleAppDashboard() http.HandlerFunc {
	t := s.parseTemplate(
		"templates/base.html",
		"templates/nav.html",
		"templates/dashboard.html",
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("heimdall_sessionToken")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		token := cookie.Value
        fmt.Println("SENDING TOKEN", token)
		body, err := json.Marshal(map[string]any{"token": token})
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}

		req, err := http.NewRequest(
            "POST",
			"http://localhost:8080/api/v1/auth/introspect",
			bytes.NewBuffer(body),
		)
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
        req.Header.Set("X-API-KEY", "0fea1f4c-0e8c-430c-bb04-504110cbe503:786bbd.BT6RCLHFa/n8ElPRqclLr1XmY474rrZIe0LOv2/ovCM=")

        res, err := http.DefaultClient.Do(req)
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
        }

		if res.StatusCode >= 300 {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		err = t.ExecuteTemplate(w, "dashboard.html", nil)
		if err != nil {
			s.respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}
	})
}

func (s *Server) parseTemplate(names ...string) *template.Template {
	return template.Must(template.ParseFS(s.Resources, names...))
}
