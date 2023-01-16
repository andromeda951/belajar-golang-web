package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml",
	))

	// "layout" tidak pakai ekstensi .gohtml kerena nama templatenya sudah di define
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Andromeda",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
