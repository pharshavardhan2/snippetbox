package main

import (
	"html/template"
	"path/filepath"

	"github.com/pharshavardhan2/snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	cache := map[string]*template.Template{}
	for _, page := range pages {
		name := filepath.Base(page)
		templateFiles := []string{
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			page,
		}

		ts, err := template.ParseFiles(templateFiles...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
