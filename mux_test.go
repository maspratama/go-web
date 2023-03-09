package go_web

import (
	"fmt"
	"github.com/maspratama/go-web/helper"
	"net/http"
	"testing"
)

func TestServerMux(t *testing.T) {
	mux := http.ServeMux{}
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello world!")
		helper.PanicIfErr(err)
	})

	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hi my name is Tama")
		helper.PanicIfErr(err)
	})

	mux.HandleFunc("/images", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Images")
		helper.PanicIfErr(err)
	})

	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Thumbnail")
		helper.PanicIfErr(err)
	})

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: &mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
