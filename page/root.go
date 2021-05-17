package page

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Root stores the current displayed page and is the main navbar navigation
type Root struct {
	app.Compo

	FormData   string // input form property value
	BoxEnabled bool   // checkbox enabled binding
	TextArea   string // text area text binding
	Selector   string // currently selected value from selector box
}

// Render will render the main navigation page
func (p *Root) Render() app.UI {
	app.Log("Render(): Root page rendered.")
	return app.Body().Body(
		app.Div().Body(
			app.Div().
				Text("CheckBox and TextArea widgets are not updated if struct bindigs for Checked() and Text() HTMLInput functions change."),
			app.Div().
				Body(
					app.Br(),
					app.Text("Plain text struct value: "+p.FormData),
					app.Br(),
					app.Text(fmt.Sprintf("CheckBox struct value: %t", p.BoxEnabled)),
					app.Br(),
					app.Text("Textarea text struct value: "+p.TextArea),
					app.Br(),
					app.Text("Selector struct value: "+p.Selector),
					app.Br(),
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
							app.Div().
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Checked(p.BoxEnabled).
												OnChange(func(ctx app.Context, e app.Event) { p.BoxEnabled = !p.BoxEnabled }),
											app.Text(" Enable Checkbox"),
										),
								),
						),
					app.Div().
						Body(
							app.Label().
								Text("Textbox data"),
							app.Div().
								Body(
									app.Textarea().
										Text(p.TextArea).
										OnChange(p.ValueTo(&p.TextArea)),
								),
						),
					app.Div().
						Body(
							app.Label().
								Text("Selector"),
							app.Div().
								Body(
									app.Select().
										Body(
											app.Option().
												Text("Foo").
												Selected(p.Selector == "Foo"),
											app.Option().
												Text("Bar").
												Selected(p.Selector == "Bar"),
										).
										OnChange(p.ValueTo(&p.Selector)),
								),
						),
					app.Br(),
					app.Div().
						Body(
							app.Button().
								Text("Clear Form, uncheck CheckBox and reset selector struct values").
								OnClick(p.onButtonClicked),
						),
				),
		),
	)
}

// onButtonClicked evemt will be triggered if button has been clicked
func (p *Root) onButtonClicked(ctx app.Context, e app.Event) {
	p.FormData = ""
	p.BoxEnabled = false
	p.TextArea = ""
	p.Selector = "Foo"
	app.Logf("onButtonClicked(): FormData struct variables reset.")

	// Workaround: Since "value" is a DOM node property it must be updated manually via DOM JS setter
	// go-app renders text content (which is not a node property and bound to DOM) correctly
	// but you have to take care if DOM node properties needs to be updated
	//app.Window().GetElementByID("formdata").Set("value", p.FormData)

	// NOTE: This issue had been fixed by go get -u -v github.com/maxence-charriere/go-app/v9@919dd2c
}
