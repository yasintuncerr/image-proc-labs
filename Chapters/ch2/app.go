package main

import (
	"fyne.io/fyne/v2/app"
	Controller "github.com/yasintuncerr/image-proc-labs/Chapters/ch2/controller"
	Model "github.com/yasintuncerr/image-proc-labs/Chapters/ch2/model"
	View "github.com/yasintuncerr/image-proc-labs/Chapters/ch2/view"
)

func main() {
	// Initialize the app and window
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Processing Application")

	// Create the model
	model := Model.NewModel()
	view := View.NewMainLayout()

	// Create the controller
	controller := Controller.NewController(view, model, myWindow)
	controller.Init()

	// Set the window content
	myWindow.SetContent(view.GetContainer())

	// Show the window
	myWindow.ShowAndRun()

}
