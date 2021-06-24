package example_test

import (
	"embed"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

//go:embed testdata/embedTemplate/*.tmpl
var templateFS embed.FS

//go:embed testdata/embedHTTP/*.txt
var textFS embed.FS

// ExampleEmbedAndTextTemplate is an example using embed and text/template.
// You can use html/template with same way.
func ExampleEmbedAndTextTemplate() {
	functions := map[string]interface{}{
		"incr": func(n int) int { return n + 1 },
	}

	// (*template.Template).Funcs() must be called before parsing
	// https://pkg.go.dev/text/template#Template.Funcs
	templates, _ := template.New("").Funcs(functions).ParseFS(templateFS, "testdata/embedTemplate/*.tmpl")

	// basename is used for template name
	_ = templates.Lookup("foo.tmpl").Execute(os.Stdout, 10)

	// Output:
	// Incremented value of 10 is 11
}

// ExampleEmbedAndHttprouter is an example using embed and github.com/julienschmidt/httprouter.
func ExampleEmbedAndHttprouter() {
	r := httprouter.New()

	r.ServeFiles("/*filepath", http.FS(textFS))

	w := httptest.NewRecorder()
	// the path is relative from here
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/testdata/embedHTTP/text.txt", nil))

	io.Copy(os.Stdout, w.Result().Body)

	// Output:
	// HELLO
}
