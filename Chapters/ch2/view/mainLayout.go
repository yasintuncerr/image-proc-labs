package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type MainLayout struct {
	OriginalImageCanvas  *ImageCanvas
	QuantizedImageCanvas *ImageCanvas
	Sidebar              *Sidebar
	Container            *fyne.Container
}

func NewMainLayout(originalImg *ImageCanvas, quantizedImg *ImageCanvas, sidebar *Sidebar) *MainLayout {
	// Create the layout container with two columns for images and a sidebar on the right
	mainLayout := &MainLayout{
		OriginalImageCanvas:  originalImg,
		QuantizedImageCanvas: quantizedImg,
		Sidebar:              sidebar,
	}

	// Create containers for the original and quantized image side by side
	imageContainer := container.New(
		layout.NewHBoxLayout(),
		originalImg.ImgCanvas,  // Original Image on the left
		quantizedImg.ImgCanvas, // Quantized Image on the right
	)

	// The sidebar will be placed to the right of the images
	mainLayout.Container = container.New(
		layout.NewHBoxLayout(),
		imageContainer,         // Images (original and quantized) in the first column
		sidebar.GetContainer(), // Sidebar in the second column
	)

	return mainLayout
}

func (ml *MainLayout) GetContainer() *fyne.Container {
	return ml.Container
}
