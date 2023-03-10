package go_web

import (
	"fmt"
	"github.com/maspratama/go-web/helper"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")

	_, err := fmt.Fprintf(w, "Hello my name is %s %s", firstName, lastName)
	helper.PanicIfErr(err)
}

func TestMultipleQueryParamater(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090?firstName=Singgih&lastName=Pratama", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	helper.PanicIfErr(err)
	bodyString := string(body)

	fmt.Println(bodyString)
}
