package api

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type TemplateHandler struct {
	Filename string
	once     sync.Once
	templ    *template.Template
}

func (t *TemplateHandler) parseTemplates(path string) {
	// Reads template files from disk and compiles them
	t.once.Do(func() {
		t.templ = template.Must(
			template.ParseFiles(
				//filepath.Join("views", t.filename),
				filepath.Join(path, t.Filename),
			),
		)
	})
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.parseTemplates("views")
	t.templ.Execute(w, r)
}
