package page

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Root stores the current displayed page and is the main navbar navigation
type Root struct {
	app.Compo

	FormData string // input form property value
}

// Render will render the main navigation page
func (p *Root) Render() app.UI {
	app.Log("Render(): Root page rendered.")
	return app.Body().Body(
		app.Div().Body(
			app.Div().
				Text("Example shows a go-app workaround clearing an input field 'value' property."),
			app.Div().
				Text("Text entered into input field is stored inside components exported 'FormData' variable using input fields OnChanged() event and reading DOM node 'value' property."),
			app.Div().
				Text("However if another event like OnClick() from a button is clearing the exported 'FormData' variable bound to input field DOM node 'value' property and updating the component the 'value' property inside the input field DOM node is not cleared."),
			app.Div().
				Body(
					app.Br(),
					app.Text("Plain text (rendered by go-app): "+p.FormData),
					app.Br(),
					app.Label().
						Text("Input field with DOM node 'value' property"),
					app.Div().
						Body(
							app.Input().
								ID("formdata").
								Value(p.FormData).
								OnChange(p.ValueTo(&p.FormData)),
						),
					app.Div().
						Body(
							app.Button().
								Text("Clear Form").
								OnClick(p.onButtonClicked),
						),
				),
		),
	)
}

// onButtonClicked evemt will be triggered if button has been clicked
func (p *Root) onButtonClicked(ctx app.Context, e app.Event) {
	p.FormData = ""
	app.Logf("onButtonClicked(): FormData variable cleared.")

	// Workaround: Since "value" is a DOM node property it must be updated manually via DOM JS setter
	// go-app renders text content (which is not a node property and bound to DOM) correctly
	// but you have to take care if DOM node properties needs to be updated
	//app.Window().GetElementByID("formdata").Set("value", p.FormData)

	// NOTE: This issue had been fixed by go get -u -v github.com/maxence-charriere/go-app/v9@919dd2c
}
