### Go-App Input Field Update Bug ###
This repository is a quick and dirty example on some render / UI update bug I found in Go-App framework.

### Details ###
An input field text value bound to some public (exported) component structure variable is not updating the UI text value if internally modified and component rendered.

In this example a input field text property is bound to component exported `FormData` variable. The `FormData` variable is updated via input fields `OnChanged()` event. 

A button is bound to `OnClick()` event that is clearing the `FormData` variable and calling the component's `update()` function.

As the content of the `FormData` variable had been cleared I would expect that after rendering the input field the text should be cleared on the UI, too. However the previous entered text stays visible and is not updated on the UI.

This bug prevents UIs that must clear input forms for entering new data.

A workaround might be to separate the input fields as dedicated components and re-initializing the input field components in case the form must be cleared but that is cubersome if
you only have a simple form.


