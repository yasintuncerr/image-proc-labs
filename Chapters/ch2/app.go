package main

import (
	"image"
	"image/color"

	"github.com/yasintuncerr/image-proc-labs/Chapters/ch2/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// Initialize the app and window
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Processing Application")

	// Create a sample original image (Gray image)
	originalImg := image.NewGray(image.Rect(0, 0, 256, 256))
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			originalImg.Set(x, y, color.Gray{uint8(x ^ y)})
		}
	}

	// Create a sample quantized image (Gray image)
	quantizedImg := image.NewGray(image.Rect(0, 0, 256, 256))
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			quantizedImg.Set(x, y, color.Gray{uint8((x ^ y) / 2)})
		}
	}

	// Create the ImageCanvas for original and quantized images
	originalImageCanvas := view.NewImageCanvas(originalImg, "Original Image")
	quantizedImageCanvas := view.NewImageCanvas(quantizedImg, "Quantized Image")

	// Declare the Sidebar variable
	var sidebar *view.Sidebar

	// Initialize the Sidebar
	sidebar = view.NewSidebar(func(level int) {
		// Example: Perform quantization and update MSE/PSNR (place your actual logic here)
		mae := 0.1234  // Dummy MSE value
		psnr := 35.678 // Dummy PSNR value

		sidebar.UpdateMAE(mae)
		sidebar.UpdatePSNR(psnr)
	}, 200) // Sidebar width is 200

	// Create the Main Layout
	mainLayout := view.NewMainLayout(originalImageCanvas, quantizedImageCanvas, sidebar)

	// Set the content of the window to the main layout container
	myWindow.SetContent(mainLayout.GetContainer())
	myWindow.Resize(fyne.NewSize(800, 600)) // Set the window size
	myWindow.ShowAndRun()
}
