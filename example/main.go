package main

import (
	"log"
	"net/http"

	"github.com/whaangbuu/heyjay"

	"github.com/julienschmidt/httprouter"
)

// Chair ...
type Chair struct {
	Width  int
	Height int
	Color  string
}

func main() {
	router := httprouter.New()
	c := Chair{Width: 10, Height: 12, Color: "Pink"}
	// GET
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		heyjay.JSON(w, c)
	})
	log.Fatal(http.ListenAndServe(":3030", router))
}
