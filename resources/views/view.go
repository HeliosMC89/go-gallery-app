package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	templateDir       = "resources/views/"
	layoutDir         = "resources/views/layouts/"
	templateExtension = ".gohtml"
)

// NewView function to create a new view by parsing passed templates.
// when function is first letter uppercase it's already exported, if we don't want to export it, we name it normally.
func NewView(files ...string) *View {
	files = append(addTemplateExtensionToFile(addTemplatePath(files)),
		layoutFiles(layoutDir)...,
	)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		// render the layout template we have at our master .gohtml
		Layout: "layout",
	}
}

// Render function to render the recently created views.
func (v *View) Render(response http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(response, v.Layout, data)
}

// View struct, to create a new object file and parse the view files.
type View struct {
	Template *template.Template
	Layout   string
}

// layout files returns a slice of strings representing.
// the layout files used in our application.
func layoutFiles(layoutDir string) []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExtension)
	if err != nil {
		panic(err)
	}
	return files
}

// addTemplatePath takes a slice of strings representing files path for templates,
// and it prepends the TemplateDir directory to each string in the slice
// Eg: the input {"home"} would be result in the output
// {"resources/views/home"} if TemplateDir == "resources/views/"
func addTemplatePath(files []string) []string {
	for i, f := range files {
		files[i] = templateDir + f
	}
	return files
}

func addTemplateExtensionToFile(files []string) []string {
	for i, f := range files {
		files[i] = f + templateExtension
	}
	return files
}
