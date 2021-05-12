### Go-App Input Field Value Update ~~Bug~~ ###
~~This repository is a quick and dirty example on some render / UI update bug I found in Go-App framework.~~

This repository shows a solution to clear input field values that are not automagically updated by go-app.

### Details ###
An input field DOM node `value` property bound to some public (exported) component structure variable is not updating the UI text if internally modified and component is rendered.

In this example a input field text DOM node `value` property is bound to component exported `FormData` variable. The `FormData` variable is updated via input fields `OnChanged()` event. 

A button is bound to `OnClick()` event that is clearing the `FormData` variable and calling the component's `update()` function.

As the content of the `FormData` variable had been cleared I would expect that after rendering the input field the text should be cleared on the UI, too. However the previous entered text stays visible and is not updated in the UI because DOM node `value` property of the input field is not updated by rendering.

~~This bug prevents UIs that must clear input forms for entering new data.~~

~~A workaround might be to separate the input fields as dedicated components and re-initializing the input field components in case the form must be cleared but that is cubersome if you only have a simple form.~~

The workaround (or in this case works as expected) is to clear the inputs fields DOM node `value` property manually at the `OnClick()` handler function.


