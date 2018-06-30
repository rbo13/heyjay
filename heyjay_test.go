package heyjay_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/whaangbuu/heyjay"
)

type User struct {
	Fullname string
	Age      int
	Address  string
}

func TestHeyJayOnCLI(t *testing.T) {
	user := User{Fullname: "Richard Burk", Age: 22, Address: "Cebu City"}
	err := heyjay.JSON(os.Stdout, user)

	if err != nil {
		t.Errorf("Error due to: %v", err)
	}
}

func TestHeyJayOnWebServer(t *testing.T) {
	user := User{Fullname: "Richard Burk", Age: 22, Address: "Cebu City"}
	expectedContentType := "application/json"
	expectedStatusCode := 200

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		heyjay.JSON(w, user)
	}

	req := httptest.NewRequest("GET", "http://localhost:3030", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	actualContentType := resp.Header.Get("Content-Type")
	actualStatusCode := resp.StatusCode

	if expectedContentType != actualContentType {
		t.Errorf("Expected to be '%s', but got '%s' instead.", expectedContentType, actualContentType)
	}

	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected to be '%d', but got '%d' instead.", expectedStatusCode, actualStatusCode)
	}

	t.Log(string(body))
}
