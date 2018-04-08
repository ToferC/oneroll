package main

import (
	"fmt"

	"github.com/andlabs/ui"
)

func GUI() {
	err := ui.Main(func() {
		ndInput := ui.NewEntry()
		hdInput := ui.NewEntry()
		wdInput := ui.NewEntry()
		button := ui.NewButton("Roll")
		results := ui.NewLabel("")
		box := ui.NewVerticalBox()
		rowD := ui.NewHorizontalBox()
		rowHD := ui.NewHorizontalBox()
		rowWD := ui.NewHorizontalBox()
		box.Append(ui.NewLabel("Set your ORE Dice Pool\n"), false)
		box.Append(ui.NewHorizontalSeparator(), false)
		rowD.Append(ndInput, false)
		rowD.Append(ui.NewLabel(" Normal Dice"), false)
		rowHD.Append(hdInput, false)
		rowHD.Append(ui.NewLabel(" Hard Dice"), false)
		rowWD.Append(wdInput, false)
		rowWD.Append(ui.NewLabel(" Wiggle Dice"), false)
		box.Append(rowD, false)
		box.Append(rowHD, false)
		box.Append(rowWD, false)
		box.Append(ui.NewHorizontalSeparator(), false)
		box.Append(ui.NewLabel(""), false)
		box.Append(button, false)
		box.Append(results, false)
		window := ui.NewWindow("ORE Die Roller", 300, 600, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			roll := Roll{}
			text := fmt.Sprintf("%sd+%shd+%swd", ndInput.Text(), hdInput.Text(), wdInput.Text())

			r, err := roll.Resolve(text)

			if err != nil {
				results.SetText(fmt.Sprintf("%s", err))
			} else {
				results.SetText(fmt.Sprintf("%s", r))
			}
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
