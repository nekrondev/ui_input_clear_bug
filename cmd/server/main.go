// +build !wasm
package main

import (
	"log"
	"net/http"
	"ui_input_clear_bug/page"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// main starts here
func main() {
	// setup main page and routing
	app.Route("/", &page.Root{})
	http.Handle("/", &app.Handler{
		Name:        "Form Clear Bug",
		Description: "Contents of input form widgets is not cleared on update",
	})

	log.Println("Starting PWA server on port 8000.")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
