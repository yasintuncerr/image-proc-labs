package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Sidebar struct {
	Slider      *widget.Slider
	ApplyBtn    *widget.Button
	MAELabel    *widget.Label
	PSNRLabel   *widget.Label
	SliderValue *widget.Label // New label to display the current slider value
	Width       float32
	Container   *fyne.Container
}

func NewSidebar(applyFunc func(level int), width float32) *Sidebar {
	slider := widget.NewSlider(1, 8)
	slider.Step = 1
	slider.Value = 1

	sliderValueLabel := widget.NewLabel(fmt.Sprintf("Level: %d", int(slider.Value)))

	slider.OnChanged = func(value float64) {
		sliderValueLabel.SetText(fmt.Sprintf("Level: %d", int(value)))
	}

	maeLabel := widget.NewLabel("MAE: Not Calculated")
	psnrLabel := widget.NewLabel("PSNR: Not Calculated")

	applyBtn := widget.NewButton("Apply", func() {
		level := int(slider.Value)
		applyFunc(level) // Call the apply function with the current slider value
	})

	sidebar := &Sidebar{
		Slider:      slider,
		ApplyBtn:    applyBtn,
		MAELabel:    maeLabel,
		PSNRLabel:   psnrLabel,
		SliderValue: sliderValueLabel,
		Width:       width, // Store the width in the struct
	}

	sidebar.Container = container.NewVBox(
		widget.NewLabel("Level"),
		slider,           // Slider in the middle
		sliderValueLabel, // Display the current slider value
		applyBtn,
		widget.NewLabel("Metrics"),
		maeLabel,
		psnrLabel,
		layout.NewSpacer(), // Push everything to the top
	)

	sidebar.Container = container.New(layout.NewVBoxLayout(), sidebar.Container)
	sidebar.Container.Resize(fyne.NewSize(sidebar.Width, 0)) // Use the width from the struct

	return sidebar
}

func (s *Sidebar) UpdateMAE(mae float64) {
	s.MAELabel.SetText(fmt.Sprintf("MAE: %.4f", mae))
}

func (s *Sidebar) UpdatePSNR(psnr float64) {
	s.PSNRLabel.SetText(fmt.Sprintf("PSNR: %.4f", psnr))
}

func (s *Sidebar) GetContainer() *fyne.Container {
	s.Container.Resize(fyne.NewSize(s.Width, 0)) // Ensure the width is applied
	return s.Container
}
