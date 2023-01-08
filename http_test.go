package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

	// ini sama dengan yang atas
	// var handler http.HandlerFunc = func(write http.ResponseWriter, request *http.Request) {
	// 	fmt.Fprint(write, "Hello Test test")
	// }

	// handler(recorder, request)

	// response = recorder.Result()
	// body, _ = io.ReadAll(recorder.Body)

	// fmt.Println(response.StatusCode)
	// fmt.Println(response.Status)
	// fmt.Println(string(body))
}
