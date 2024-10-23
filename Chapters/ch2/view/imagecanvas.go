package view

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type ImageCanvas struct {
	ImgCanvas  *canvas.Image
	TitleLabel *widget.Label
	width      float32
	height     float32
}

func NewImageCanvas(img *image.Gray, title string, width, height float32) *ImageCanvas {
	imgCanvas := canvas.NewImageFromImage(img)
	imgCanvas.FillMode = canvas.ImageFillContain
	imgCanvas.SetMinSize(fyne.NewSize(width, height))

	titleLabel := widget.NewLabel(title)
	return &ImageCanvas{
		ImgCanvas:  imgCanvas,
		TitleLabel: titleLabel,
	}
}

func (ic *ImageCanvas) UpdateImage(img *image.Gray) {
	ic.ImgCanvas.Image = img
	ic.ImgCanvas.Refresh()
}

func (ic *ImageCanvas) UpdateTitle(title string) {
	ic.TitleLabel.SetText(title)
}
