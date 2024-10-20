package main

import (
	"image"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/yasintuncerr/image-proc-labs/Chapters/ch1/controller"
	"github.com/yasintuncerr/image-proc-labs/Chapters/ch1/imageprocessing"
	"github.com/yasintuncerr/image-proc-labs/Chapters/ch1/model"
	"github.com/yasintuncerr/image-proc-labs/Chapters/ch1/view"
)

func main() {
	// Create a new image
	img := image.NewGray(image.Rect(0, 0, 512, 512))

	imageprocessing.Fill(img.Pix, 0)

	// Create a new model
	Model := model.NewBrightnessModel(img)

	// Create a new view
	View := model.NewBrightnessView(img)

	// Create a new controller
	Controller := controller.NewBrightnessController(Model, View)

	increaseButton := widget.NewButton("Increase", Controller.IncreaseBrightness)

	decreaseButton := widget.NewButton("Decrease", Controller.DecreaseBrightness)

	saveButton := widget.NewButton("Save", Controller.SaveBrightnessLevel)

	setStepButton := widget.NewButton("Set Step", Controller.SetStep)

	// Create a new app
	app := app.New()

	// Create a new window
	w := app.NewWindow("Brightness Controller")

	// Set the window content
	w.SetContent(container.NewVBox(
		view.ImgCanvas,
		view.BrightnessLabel,
		view.StepLabel,
		view.StatusLabel,
		increaseButton,
		decreaseButton,
		saveButton,
		view.SavedLevelsText,
		setStepButton,
	))

	// Show the window
	w.ShowAndRun()
}
