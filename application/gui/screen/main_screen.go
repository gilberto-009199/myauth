package view

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MainView(app fyne.App) {

	window := app.NewWindow("Chaveiro")

	//fyne.LoadResourceFromPath("23")
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentPrintIcon(), func() {
			CaptureView(app)
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewList(
		func() int {
			return data.Length()
		},
		// Template
		func() fyne.CanvasObject {

			ctr_item := container.NewVBox()
			item := container.NewHBox()
			item_desc := container.NewVBox()

			item_icon := widget.NewIcon(theme.AccountIcon())

			item_title := canvas.NewText("Name", color.Black)
			item_title.TextSize = theme.TextSize() * 0.8
			item_title.Alignment = fyne.TextAlignLeading
			item_title.TextStyle = fyne.TextStyle{Italic: true}

			item_code := widget.NewLabel("Code")
			item_code.TextStyle = fyne.TextStyle{Bold: true}

			item_desc.Add(item_title)
			item_desc.Add(item_code)

			item.Add(item_icon)
			item.Add(item_desc)

			ctr_item.Add(item)

			progressbar := canvas.NewRectangle(theme.PrimaryColor())
			progressbar.Resize(fyne.NewSize(10, 2))
			ctr_item.Add(container.NewWithoutLayout(progressbar))

			return ctr_item
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			fmt.Sprintln(id)
			data_item, error := data.GetValue(id)
			if error != nil {
				panic("interate list ")
			} else {

				ctr_item := item.(*fyne.Container)

				item := ctr_item.Objects[0].(*fyne.Container)
				progressbar := ctr_item.Objects[1].(*fyne.Container).Objects[0].(*canvas.Rectangle)
				item_desc := item.Objects[1].(*fyne.Container)
				//item_icon := item.Objects[0].(*widget.Icon)

				item_title := item_desc.Objects[0].(*canvas.Text)
				item_title.Text = (data_item)

				item_code := item_desc.Objects[1].(*widget.Label)
				item_code.SetText("222 4342")

				//var second_OTP float64 = 15

				go func() {
					for {
						time.Sleep(time.Second / 4)
						// 90 ... 270 - horizontal
						//max_progress := ctr_item.Size().Width;
						//progress := progressbar.Size().Width;
						progressbar.Resize(fyne.NewSize(50, 2))
						//gradient.CenterOffsetY = 0

						//gradient.Angle = 90

						//gradient.Angle = 150

						//gradient.Angle += (270 - 90) / second_OTP
						//if gradient.Angle > 270 || gradient.Angle < 90 {
						//	gradient.Angle = 90
						//}

						ctr_item.Refresh()
					}
				}()

			}

		},
	)
	content := container.NewBorder(toolbar, widget.NewButton("Add", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
		list.Refresh()
	}), nil, nil, list)

	window.SetContent(content)

	//	window.SetContent(widget.NewLabel("Hello"))
	window.Resize(fyne.NewSize(300, 500))
	window.CenterOnScreen()
	//window.SetFullScreen(true)
	window.Show()

}
