package main

import (
	"image"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create the Fyne app and window
	myApp := app.New()
	w := myApp.NewWindow("Human Brightness Sensitivity Detector")

	// Initialize the image
	img := image.NewGray(image.Rect(0, 0, 512, 512))
	Fill(img.Pix, 0)

	// Initialize Model, View, Controller
	Model := NewBrightnessModel(img)
	View := NewBrightnessView(img)

	Controller := NewBrightnessController(Model, View)

	// Create buttons
	increaseButton := widget.NewButton("Increase", Controller.IncreaseBrightness)
	decreaseButton := widget.NewButton("Decrease", Controller.DecreaseBrightness)
	setStepButton := widget.NewButton("Set Step", Controller.SetStep)
	saveButton := widget.NewButton("Save", Controller.SaveBrightnessLevel)
	resetButton := widget.NewButton("Reset", Controller.Reset)

	msgButton := widget.NewButton("Calculate Human Sensitivity", func() {
		level, val, err := Controller.CalculateSensivity()
		if err != nil {
			dialog.ShowError(err, w)
			return // Return early if there is an error
		}

		levelStr := strconv.Itoa(level)                // Convert level to string
		valStr := strconv.FormatFloat(val, 'f', 2, 64) // Convert val to a formatted string with 2 decimal places

		message := "You detected the brightness level " + levelStr + " in an 8-bit grayscale image.\n" +
			"Your average detection step ratio is " + valStr + "."

		dialog.ShowInformation("Human Sensitivity", message, w)
	})

	leftLayout := container.NewVBox(
		View.ImgCanvas,   // Image at the top
		View.ProgressBar, // Use ProgressBar from View (already defined and initialized)
		container.NewGridWithColumns(4, // Buttons below the progress bar (horizontally)
			increaseButton, decreaseButton, saveButton, resetButton),
	)

	rightLayout := container.NewVBox(
		View.StepLabel,       // Step label
		setStepButton,        // Set step button
		View.BrightnessLabel, // Current brightness level
		View.SavedLevelsText, // Saved levels
		View.StatusLabel,     // Status message
		msgButton,            // Sensitivity calculation button
	)

	rightLayout.Size()

	splitLayout := container.NewHSplit(leftLayout, rightLayout)
	splitLayout.SetOffset(0.6) // Adjust the split ratio (60% left, 40% right)

	w.SetContent(splitLayout)

	w.ShowAndRun()
}
