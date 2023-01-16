package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates2 embed.FS

var myTemplate = template.Must(template.ParseFS(templates2, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
