package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Andromeda",
		"Address": map[string]interface{}{
			"Street": "Jalan belum ada",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Andromeda",
		Address: Address{
			Street: "Jalan belum ada",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
