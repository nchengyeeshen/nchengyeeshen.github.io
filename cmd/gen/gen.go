package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Generator generates the static website.
type Generator struct {
	Assets     fs.FS
	ReadFile   func(name string) ([]byte, error)
	CreateFile func(name string) (*os.File, error)
	MkdirAll   func(path string, perm os.FileMode) error
}

func (g Generator) Run(outDir string) error {
	outputs := []struct {
		Path         string
		TemplateName string
		Data         any
	}{
		{
			Path:         "/",
			TemplateName: "pages/home.tmpl",
			Data:         newTemplateData(),
		},
	}

	for _, blogEntry := range blogEntries {
		contents, err := g.ReadFile(blogEntry.ContentPath)
		if err != nil {
			return fmt.Errorf("read file %s: %v", blogEntry.ContentPath, err)
		}

		outputs = append(outputs, struct {
			Path         string
			TemplateName string
			Data         any
		}{
			Path:         fmt.Sprintf("/blog/%s", blogEntry.Slug),
			TemplateName: "pages/blog.tmpl",
			Data: func() any {
				d := newTemplateData()
				d["Page"] = map[string]any{
					"Title":         blogEntry.Title,
					"PublishedDate": blogEntry.PublishedDate,
					"Content":       string(contents),
				}
				return d
			}(),
		})
	}

	for _, output := range outputs {
		dirPath := filepath.Join(outDir, output.Path)
		if err := g.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("make directory %s: %v", dirPath, err)
		}

		indexFile := filepath.Join(dirPath, "index.html")
		f, err := g.CreateFile(indexFile)
		if err != nil {
			return fmt.Errorf("create file %s: %v", indexFile, err)
		}

		if err := g.renderTemplate(f, output.TemplateName, output.Data); err != nil {
			return fmt.Errorf("write template to %s: %v", indexFile, err)
		}

		if err := f.Close(); err != nil {
			return fmt.Errorf("close %s: %v", indexFile, err)
		}
	}

	return nil
}

func (g Generator) renderTemplate(w io.Writer, templateName string, data any) error {
	ts, err := template.New("").Funcs(TemplateFuncs).ParseFS(
		g.Assets,
		"templates/base.tmpl",
		"templates/partials/*.tmpl",
		"templates/"+templateName,
	)
	if err != nil {
		return fmt.Errorf("parse templates: %v", err)
	}

	buf := new(bytes.Buffer)

	if err := ts.ExecuteTemplate(buf, "base", data); err != nil {
		return fmt.Errorf("execute template: %v", err)
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("write: %v", err)
	}

	return nil
}
