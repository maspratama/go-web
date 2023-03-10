package go_web

import (
	"fmt"
	"github.com/maspratama/go-web/helper"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, request.Method)
		helper.PanicIfErr(err)

		_, err = fmt.Fprintln(writer, request.RequestURI)
		helper.PanicIfErr(err)
	}

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
