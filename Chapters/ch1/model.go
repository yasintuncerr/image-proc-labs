package model

import (
	"image"
	"log"
)

type BrightnessModel struct {
	Img            *image.Gray
	Brightness     int
	IncrementStep  int
	DetectedLevels []int
}

func NewBrightnessModel(img *image.Gray) *BrightnessModel {
	return &BrightnessModel{
		Img:           img,
		Brightness:    0,
		IncrementStep: 1,
	}
}

func (bm *BrightnessModel) IncrementBrightness() error {
	err := proc.AdjustBrightness(bm.Img, bm.Brightness)
	if err != nil {
		return err
	}

	bm.Brightness += bm.IncrementStep

	return nil
}

func (bm *BrightnessModel) DecrementBrightness() error {
	err := proc.AdjustBrightness(bm.Img, -bm.Brightness)
	if err != nil {
		return err
	}

	bm.Brightness -= bm.IncrementStep
}
func (bm *BrightnessModel) SaveBrightnessLevel() {
	bm.DetectedLevels = append(bm.DetectedLevels, bm.Brightness)
	log.Println("Saved level: ", bm.Brightness)
}

func (bm *BrightnessModel) SetStep(step int) {
	bm.IncrementStep = step
}

func (bm *BrightnessModel) Reset() {

	proc.Fill(bm.Img.Pix, 0)
	bm.Brightness = 0
	bm.DetectedLevels = nil
}
