### Go-App UI Update Bugs ###
I found three more issues with UI updates for rendered elements.

### Details ###

1. Input fields that are checkboxes are not updated if struct value is reset to unchecked or vice versa

- Tick checkbox
- Struct value `p.BoxEnabled` is set to true, checkbox shows ticked off square in UI
- Press reset form buton. The internal `p.BoxEnabled` value is set back to `false`.
- Checkbox UI square is still ticked off even if internal struct value is `false`.

2. TextArea text field is not updated if struct value is cleared

- Write some text into textarea widget
- Struct value `p.TextArea` is updated with textarea value on change
- Press reset form button. The internal `p.TextArea` value is cleared
- TextArea UI still contains text

3. Selector option not reset

- Select `Bar` value from selector
- Struct value `p.Selector` is updated with selected value `Bar`
- Press reset form button. The internal `p.Selector` value is reset to `Foo`.
- Selector UI still shows `Bar` as selected value

