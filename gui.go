package oneroll

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/andlabs/ui"
)

func parseNumRolls(s string) (int, error) {

	re := regexp.MustCompile("[0-9]+")

	var num int
	var numString string

	numString = re.FindString(s)
	num, err := strconv.Atoi(numString)
	if err != nil {
		num = 1
	}
	return num, err
}

// GUI renders a GUI for the die rolling app
func GUI() {
	err := ui.Main(func() {
		ndInput := ui.NewEntry()
		hdInput := ui.NewEntry()
		wdInput := ui.NewEntry()
		numInput := ui.NewEntry()
		button := ui.NewButton("Roll")
		results := ui.NewLabel("")
		box := ui.NewVerticalBox()
		rowD := ui.NewHorizontalBox()
		rowHD := ui.NewHorizontalBox()
		rowWD := ui.NewHorizontalBox()
		rowNum := ui.NewHorizontalBox()

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

		rowNum.Append(numInput, false)
		rowNum.Append(ui.NewLabel(" Number of Rolls"), false)

		box.Append(ui.NewHorizontalSeparator(), false)
		box.Append(rowNum, false)
		box.Append(button, false)
		box.Append(results, false)

		window := ui.NewWindow("ORE Die Roller", 300, 600, false)
		window.SetMargined(true)
		window.SetChild(box)

		button.OnClicked(func(*ui.Button) {

			var resultString string

			numRolls, err := parseNumRolls(numInput.Text())

			if err != nil {
				resultString += "Invalid number of rolls. Set to 1.\n\n"
			}

			for x := 1; x < numRolls+1; x++ {

				resultString += fmt.Sprintf("Roll #%d\n\n", x)

				roll := Roll{}
				text := fmt.Sprintf("%sd+%shd+%swd", ndInput.Text(), hdInput.Text(), wdInput.Text())

				r, err := roll.Resolve(text)

				if err != nil {
					resultString += fmt.Sprintf("%s", err)
				} else {
					resultString += fmt.Sprintf("%s", r)
				}
			}
			results.SetText(fmt.Sprintf("%s", resultString))

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
