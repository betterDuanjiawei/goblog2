package view

import (
	"goblog2/pkg/logger"
	"goblog2/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, name string, data interface{}) {
	viewDir := "resources/views/"
	name = strings.Replace(name, ".", "/", -1)

	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	newFiles := append(files, viewDir+name+".gohtml")

	tmpl, err := template.New(name + ".gohtml").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(newFiles...)
	logger.LogError(err)
	tmpl.ExecuteTemplate(w, "app", data)
}
