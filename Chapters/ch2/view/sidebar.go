package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Sidebar struct {
	OpenFileBtn *widget.Button
	ExportBtn   *widget.Button
	Slider      *widget.Slider
	ApplyBtn    *widget.Button
	MAELabel    *widget.Label
	PSNRLabel   *widget.Label
	SliderValue *widget.Label
	Width       float32
	Container   *fyne.Container
}

func NewSidebar(width float32) *Sidebar {
	OpenFileBtn := widget.NewButton("Open File", func() {
		// Placeholder for the open file function
	})

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
		// Placeholder for the apply function
	})

	exportBtn := widget.NewButton("Export", func() {
		// Placeholder for the export function
	})

	sidebar := &Sidebar{
		OpenFileBtn: OpenFileBtn,
		ExportBtn:   exportBtn,
		Slider:      slider,
		ApplyBtn:    applyBtn,
		MAELabel:    maeLabel,
		PSNRLabel:   psnrLabel,
		SliderValue: sliderValueLabel,
		Width:       width,
	}

	sidebar.Container = container.NewVBox(
		OpenFileBtn,
		layout.NewSpacer(),
		exportBtn,
		layout.NewSpacer(),
		sliderValueLabel,
		slider,
		layout.NewSpacer(),
		applyBtn,
		layout.NewSpacer(),
		widget.NewLabel("Metrics"),
		maeLabel,
		psnrLabel,
		layout.NewSpacer(),
	)

	sidebar.Container = container.New(layout.NewVBoxLayout(), sidebar.Container)
	sidebar.Container.Resize(fyne.NewSize(sidebar.Width, 0))

	return sidebar
}

func (s *Sidebar) UpdateMAE(mae float64) {
	s.MAELabel.SetText(fmt.Sprintf("MAE: %.4f", mae))
}

func (s *Sidebar) UpdatePSNR(psnr float64) {
	s.PSNRLabel.SetText(fmt.Sprintf("PSNR: %.4f", psnr))
}

func (s *Sidebar) GetContainer() *fyne.Container {
	s.Container.Resize(fyne.NewSize(s.Width, 0))
	return s.Container
}
