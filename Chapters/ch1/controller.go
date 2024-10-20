package main

import (
	"errors"
	"strconv"
)

type BrightnessController struct {
	Model *BrightnessModel
	View  *BrightnessView
}

func NewBrightnessController(model *BrightnessModel, view *BrightnessView) *BrightnessController {
	return &BrightnessController{
		Model: model,
		View:  view,
	}
}

func (bc *BrightnessController) IncreaseBrightness() {
	err := bc.Model.IncrementBrightness()
	if err != nil {
		bc.View.StatusLabel.SetText(err.Error())
		return
	}
	bc.View.ProgressBar.SetValue(float64(bc.Model.Brightness))
	bc.View.UpdateUI(bc.Model.Brightness, bc.Model.Img)
}

func (bc *BrightnessController) DecreaseBrightness() {
	err := bc.Model.DecrementBrightness()
	if err != nil {
		bc.View.StatusLabel.SetText(err.Error())
		return
	}
	bc.View.ProgressBar.SetValue(float64(bc.Model.Brightness))
	bc.View.UpdateUI(bc.Model.Brightness, bc.Model.Img)
}

func (bc *BrightnessController) SaveBrightnessLevel() {
	if bc.Model.SaveBrightnessLevel() {
		bc.View.UpdateSavedLevels(bc.Model.DetectedLevels)
	}
}

func (bc *BrightnessController) SetStep() {
	step_value := bc.View.StepLabel.Text
	step, err := strconv.Atoi(step_value)
	if err != nil {
		bc.View.StatusLabel.SetText("Invalid step value")
		bc.View.StepLabel.SetText("")
		return
	}
	if step < 1 {
		bc.View.StatusLabel.SetText("Step value must be greater than 0")
		bc.View.StepLabel.SetText("")
		return
	}

	bc.Model.SetStep(step)
	bc.View.StatusLabel.SetText("Step value set to " + step_value)

}

func (bc *BrightnessController) Reset() {
	bc.Model.Reset()
	bc.View.ProgressBar.SetValue(0)
	bc.View.UpdateUI(bc.Model.Brightness, bc.Model.Img)
	bc.View.UpdateSavedLevels(bc.Model.DetectedLevels)
	bc.View.StatusLabel.SetText("")
}

func (bc *BrightnessController) CalculateSensivity() (int, float64, error) {
	if len(bc.Model.DetectedLevels) == 0 {
		return 0, 0.0, errors.New("No detected levels yet")
	}

	sum := float64(bc.Model.DetectedLevels[0])

	for i := 1; i < len(bc.Model.DetectedLevels); i++ {
		sum += float64(bc.Model.DetectedLevels[i]) - float64(bc.Model.DetectedLevels[i-1])
	}

	max_level := bc.Model.DetectedLevels[len(bc.Model.DetectedLevels)-1]

	founded_levels := int(255/max_level) * len(bc.Model.DetectedLevels)

	return founded_levels, sum / float64(len(bc.Model.DetectedLevels)), nil
}
