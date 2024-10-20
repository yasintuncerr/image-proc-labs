package view

import (
	"image"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type BrightnessView struct {
	ImgCanvas       *canvas.Image
	BrightnessLabel *widget.Label
	StepLabel       *widget.Label
	StatusLabel     *widget.Label
	SavedLevelsText *widget.MultiLineEntry
}

func NewBrightnessView(img *image.Gray) *BrightnessView {
	imgCanvas := canvas.NewImageFromImage(img)
	imgCanvas.FillMode = canvas.ImageFillOriginal

	brightnessLabel := widget.NewLabel("Brightness: 0")
	stepLabel := widget.NewLabel("Step: 1")
	statusLabel := widget.NewLabel("")

	savedLevelsText := widget.NewMultiLineEntry()
	savedLevelsText.Disable()

	return &BrightnessView{
		ImgCanvas:       imgCanvas,
		BrightnessLabel: brightnessLabel,
		StepLabel:       stepLabel,
		StatusLabel:     statusLabel,
		SavedLevelsText: savedLevelsText,
	}
}

func (bv *BrightnessView) UpdateUI(brightness int, img *image.Gray) {
	bv.ImgCanvas.Image = img
	bv.BrightnessLabel.SetText("Brightness: " + strconv.Itoa(brightness))
	bv.StatusLabel.SetText("")
	bv.ImgCanvas.Refresh()
}

func (bv *BrightnessView) UpdateSavedLevels(detectedLevels []int) {
	var levelsStr []string
	for _, level := range detectedLevels {
		levelsStr = append(levelsStr, strconv.Itoa(level))
	}
	bv.SavedLevelsText.SetText(strings.Join(levelsStr, "\n"))
	bv.StatusLabel.SetText("Saved level!")
}
