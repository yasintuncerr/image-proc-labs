package view

import (
	"image"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type ImageCanvas struct {
	ImgCanvas  *canvas.Image
	TitleLabel *widget.Label
}

func NewImageCanvas(img *image.Gray, title string) *ImageCanvas {
	imgCanvas := canvas.NewImageFromImage(img)
	imgCanvas.FillMode = canvas.ImageFillOriginal
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
