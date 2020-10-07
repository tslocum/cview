// Demo code for unicode support (demonstrates wide Chinese characters).
package main

import (
	"fmt"

	"gitlab.com/tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	pages := cview.NewPages()

	form := cview.NewForm()
	form.AddDropDownSimple("称谓", 0, nil, "先生", "女士", "博士", "老师", "师傅")
	form.AddInputField("姓名", "", 20, nil, nil)
	form.AddPasswordField("密码", "", 10, '*', nil)
	form.AddCheckBox("", "年龄 18+", false, nil)
	form.AddButton("保存", func() {
		_, option := form.GetFormItem(0).(*cview.DropDown).GetCurrentOption()
		userName := form.GetFormItem(1).(*cview.InputField).GetText()

		alert(pages, "alert-dialog", fmt.Sprintf("保存成功，%s %s！", userName, option.GetText()))
	})
	form.AddButton("退出", func() {
		app.Stop()
	})
	form.SetBorder(true)
	form.SetTitle("输入一些内容")
	form.SetTitleAlign(cview.AlignLeft)
	pages.AddPage("base", form, true, true)

	app.SetRoot(pages, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

// alert shows a confirmation dialog.
func alert(pages *cview.Pages, id string, message string) {
	modal := cview.NewModal()
	modal.SetText(message)
	modal.AddButtons([]string{"确定"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage(id)
		pages.RemovePage(id)
	})

	pages.AddPage(id, modal, false, true)
}
