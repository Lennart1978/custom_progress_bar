package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/lennart1978/custom_progress_bar/tempWidget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Temperature Widget")
	label1 := widget.NewLabel("Temperature widget 1:")
	label1.Alignment = fyne.TextAlignCenter
	label2 := widget.NewLabel("Temperature widget 2:")
	label2.Alignment = fyne.TextAlignCenter

	tempWidget1 := tempWidget.NewTemperatureWidget()
	tempWidget1.Temperature = -25.5
	tempWidget2 := tempWidget.NewTemperatureWidget()
	tempWidget2.Temperature = 35.8
	myWindow.SetContent(container.NewVBox(label1, tempWidget1, label2, tempWidget2))
	myWindow.ShowAndRun()
}
