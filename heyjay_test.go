package heyjay_test

import (
	"fmt"
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
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		heyjay.JSON(w, user)
	}

	req := httptest.NewRequest("GET", "http://localhost:3030", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Output:
	// 200
	// application/json; charset=utf-8
}
