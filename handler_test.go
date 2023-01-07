package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(write http.ResponseWriter, request *http.Request) {
		fmt.Fprint(write, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(write http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(write, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "Hi")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "Images")
	})

	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "Thumbnails")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}