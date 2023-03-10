package go_web

import (
	"fmt"
	"github.com/maspratama/go-web/helper"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//implementasi http testing

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Test running unit test")

}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	helper.PanicIfErr(err)
	bodyString := string(body)

	fmt.Println(bodyString)

}
