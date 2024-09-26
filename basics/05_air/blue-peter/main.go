package main

import (
	"05_air/libs/ui"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := ui.Hello("Underworld!")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}