package page

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// root stores the current displayed page and is the main navbar navigation
type Root struct {
	app.Compo

	FormData string // input form property value
	initDone bool   // init was done
}

// init initializes the root page
func (p *Root) init() {
	if p.initDone {
		return
	}
	p.FormData = ""
	p.initDone = true
}

// Render will render the main navigation page
func (p *Root) Render() app.UI {
	app.Log("Render(): Root page rendered.")
	p.init()
	return app.Body().Class("has-navbar-fixed-top").Body(
		app.Div().Body(
			app.Div().
				Text("Example shows a go-app bug clearing an input field."),
			app.Div().
				Text("Text entered into input field is stored inside components exported 'FormData' variable using input fields OnChanged() event."),
			app.Div().
				Text("However if another event like OnClick() from a button is clearing the exported 'FormData' variable bound to input field and updating the component the text inside the input field is not cleared."),
			app.Div().
				Body(
					app.Br(),
					app.Label().
						Text("Input field"),
					app.Div().
						Body(
							app.Input().
								Value(p.FormData).
								OnChange(p.onInputChanged),
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

// onInputChanged event is triggered when input field content has changed
func (p *Root) onInputChanged(ctx app.Context, e app.Event) {
	p.FormData = ctx.JSSrc.Get("value").String()
	app.Logf("onInputChanged(): FormData variable new value => '%s'.", p.FormData)
	p.Update()
}

// onButtonClicked evemt will be triggered if button has been clicked
func (p *Root) onButtonClicked(ctx app.Context, e app.Event) {
	p.FormData = ""
	app.Logf("onButtonClicked(): FormData variable cleared.")
	p.Update()
}

// OnMount will mount the root page
func (p *Root) OnMount(ctx app.Context) {
	app.Log("OnMount(): Root page mounted.")
}
