package view

import (
	"image"

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

func NewMainLayout() *MainLayout {

	img := image.NewGray(image.Rect(0, 0, 512, 512))

	originalImg := NewImageCanvas(img, "Original Image", 800, 600)
	quantizedImg := NewImageCanvas(img, "Quantized Image", 800, 600)
	sidebar := NewSidebar(400)

	mainLayout := &MainLayout{
		OriginalImageCanvas:  originalImg,
		QuantizedImageCanvas: quantizedImg,
		Sidebar:              sidebar,
	}

	imageContainer := container.New(
		layout.NewHBoxLayout(),
		originalImg.ImgCanvas,
		quantizedImg.ImgCanvas,
	)

	mainLayout.Container = container.New(
		layout.NewHBoxLayout(),
		imageContainer,
		sidebar.GetContainer(),
	)

	return mainLayout
}

func (ml *MainLayout) GetContainer() *fyne.Container {
	return ml.Container
}

func (ml *MainLayout) GetWindow() fyne.Window {
	return fyne.CurrentApp().NewWindow("Image Processing Lab")
}
