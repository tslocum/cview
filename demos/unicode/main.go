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
	form.AddDropDownSimple("称谓", 0, nil, "先生", "女士", "博士", "老师", "师傅").
		AddInputField("姓名", "", 20, nil, nil).
		AddPasswordField("密码", "", 10, '*', nil).
		AddCheckBox("", "年龄 18+", false, nil).
		AddButton("保存", func() {
			_, option := form.GetFormItem(0).(*cview.DropDown).GetCurrentOption()
			userName := form.GetFormItem(1).(*cview.InputField).GetText()

			alert(pages, "alert-dialog", fmt.Sprintf("保存成功，%s %s！", userName, option.GetText()))
		}).
		AddButton("退出", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("输入一些内容").SetTitleAlign(cview.AlignLeft)
	pages.AddPage("base", form, true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

// alert shows a confirmation dialog.
func alert(pages *cview.Pages, id string, message string) *cview.Pages {
	return pages.AddPage(
		id,
		cview.NewModal().
			SetText(message).
			AddButtons([]string{"确定"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.HidePage(id).RemovePage(id)
			}),
		false,
		true,
	)
}
