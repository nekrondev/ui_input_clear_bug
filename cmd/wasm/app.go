// +build wasm

package main

import (
	"ui_input_clear_bug/page"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &page.Root{})
	app.RunWhenOnBrowser()
}
