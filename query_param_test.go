package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello anonymous")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestQueyParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Joko&last_name=Malik", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joko&name=Hariyadi", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
