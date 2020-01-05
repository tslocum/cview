package main

import (
	"git.sr.ht/~tslocum/cview"
)

const form = `[green]package[white] main

[green]import[white] (
    [red]"git.sr.ht/~tslocum/cview"[white]
)

[green]func[white] [yellow]main[white]() {
    form := cview.[yellow]NewForm[white]().
        [yellow]AddInputField[white]([red]"First name:"[white], [red]""[white], [red]20[white], nil, nil).
        [yellow]AddInputField[white]([red]"Last name:"[white], [red]""[white], [red]20[white], nil, nil).
        [yellow]AddDropDown[white]([red]"Role:"[white], [][green]string[white]{
            [red]"Engineer"[white],
            [red]"Manager"[white],
            [red]"Administration"[white],
        }, [red]0[white], nil).
        [yellow]AddCheckbox[white]([red]"On vacation:"[white], false, nil).
        [yellow]AddPasswordField[white]([red]"Password:"[white], [red]""[white], [red]10[white], [red]'*'[white], nil).
        [yellow]AddButton[white]([red]"Save"[white], [yellow]func[white]() { [blue]/* Save data */[white] }).
        [yellow]AddButton[white]([red]"Cancel"[white], [yellow]func[white]() { [blue]/* Cancel */[white] })
    cview.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](form, true).
        [yellow]Run[white]()
}`

// Form demonstrates forms.
func Form(nextSlide func()) (title string, content cview.Primitive) {
	f := cview.NewForm().
		AddInputField("First name:", "", 20, nil, nil).
		AddInputField("Last name:", "", 20, nil, nil).
		AddDropDown("Role:", []string{"Engineer", "Manager", "Administration"}, 0, nil).
		AddPasswordField("Password:", "", 10, '*', nil).
		AddCheckbox("", "On vacation", false, nil).
		AddButton("Save", nextSlide).
		AddButton("Cancel", nextSlide)
	f.SetBorder(true).SetTitle("Employee Information")
	return "Forms", Code(f, 36, 15, form)
}
