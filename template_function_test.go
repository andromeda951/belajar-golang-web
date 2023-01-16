package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", my name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Budi" }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Andromeda",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))

	t.ExecuteTemplate(writer, "FUNCTION", map[string]interface{}{
		"Name": "Belajar Go-Lang",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateCreateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Andromeda",
	})
}

func TestTemplateCreateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateCreateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Andromeda",
	})
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
