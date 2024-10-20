package main

import (
	"fmt"
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
	err := AdjustBrightness(bm.Img, bm.IncrementStep)
	if err != nil {
		return err
	}

	if bm.Brightness == 255 {
		return fmt.Errorf("Brightness level cannot be greater than 255")
	}

	bm.Brightness += bm.IncrementStep

	return nil
}

func (bm *BrightnessModel) DecrementBrightness() error {
	err := AdjustBrightness(bm.Img, -bm.IncrementStep)
	if err != nil {
		return err
	}
	if bm.Brightness == 0 {
		return fmt.Errorf("Brightness level cannot be less than 0")
	}

	bm.Brightness -= bm.IncrementStep
	return nil
}

func (bm *BrightnessModel) SaveBrightnessLevel() bool {
	for _, level := range bm.DetectedLevels {
		if level == bm.Brightness {
			log.Println("Level already saved: ", bm.Brightness)
			return false
		}
	}

	bm.DetectedLevels = append(bm.DetectedLevels, bm.Brightness)
	log.Println("Saved level: ", bm.Brightness)
	return true
}

func (bm *BrightnessModel) SetStep(step int) bool {
	if step < 1 {
		log.Println("Step value must be greater than 0")
		return false
	} else if step > 255 {
		log.Println("Step value must be less than 255")
		return false
	}

	bm.IncrementStep = step
	log.Println("Step value set to ", step)
	return true
}

func (bm *BrightnessModel) Reset() {

	Fill(bm.Img.Pix, 0)
	bm.Brightness = 0
	bm.DetectedLevels = nil
}
