package go_web

import (
	"fmt"
	"github.com/maspratama/go-web/helper"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// mengambil semua value dengan menggunkan query[key]

func MultipleValuesQueryParameter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	_, err := fmt.Fprint(w, strings.Join(names, " "))
	helper.PanicIfErr(err)
}

func TestMultipleValuesQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090?name=Mas&name=Singgih&name=Pratama", nil)
	recorder := httptest.NewRecorder()

	MultipleValuesQueryParameter(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	helper.PanicIfErr(err)

	fmt.Println(string(body))

}
