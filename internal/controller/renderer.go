package controller

import (
	"html/template"
	"io"

	"github.com/babattles/snoqualmie-crust-calculator/web"
	"github.com/labstack/echo/v4"
)

type Renderer struct {
	tmpl *template.Template
}

func NewRenderer() (*Renderer, error) {
	tmpl, err := template.ParseFS(web.FS, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return &Renderer{tmpl: tmpl}, nil
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.tmpl.ExecuteTemplate(w, name, data)
}
