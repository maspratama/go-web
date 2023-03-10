package go_web

import (
	"fmt"
	"github.com/maspratama/go-web/helper"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//mengambil queryparameter

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		_, err := fmt.Fprint(writer, "hello")
		helper.PanicIfErr(err)
	} else {
		_, err := fmt.Fprintf(writer, "Hello %s", name)
		helper.PanicIfErr(err)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090?name=Tama", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	helper.PanicIfErr(err)
	bodyString := string(body)

	fmt.Println(bodyString)
}
