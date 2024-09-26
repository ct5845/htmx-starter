package main

import (
	"04_go_templ/libs/ui"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := ui.Hello("World!")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}