package controller

import (
	"strconv"

	"github.com/yasintuncerr/image-proc-labs/Chapters/ch1/model"
	"github.com/yasintuncerr/image-proc-labs/Chapters/ch1/view"
)

type BrightnessController struct {
	Model *model.BrightnessModel
	View  *view.BrightnessView
}

func NewBrightnessController(model *model.BrightnessModel, view *view.BrightnessView) *BrightnessController {
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

	bc.View.UpdateUI(bc.Model.Brightness, bc.Model.Img)
}

func (bc *BrightnessController) DecreaseBrightness() {
	err := bc.Model.DecrementBrightness()
	if err != nil {
		bc.View.StatusLabel.SetText(err.Error())
		return
	}

	bc.View.UpdateUI(bc.Model.Brightness, bc.Model.Img)
}

func (bc *BrightnessController) SaveBrightnessLevel() {
	bc.Model.SaveBrightnessLevel()
	bc.View.UpdateSavedLevels(bc.Model.DetectedLevels)
}

func (bc *BrightnessController) SetStep(step int) {
	step_value := bc.View.StepEntry.Text
	step, err := strconv.Atoi(step_value)
	if err != nil {
		bc.View.StatusLabel.SetText("Invalid step value")
		bc.View.StepEntry.SetText("")
		return
	}
	if step < 1 {
		bc.View.StatusLabel.SetText("Step value must be greater than 0")
		bc.View.StepEntry.SetText("")
		return
	}

	bc.Model.SetStep(step)
	bc.View.StatusLabel.SetText("Step value set to " + step_value)

}
