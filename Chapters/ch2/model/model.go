package model

import (
	"fmt"
	"image"
)

type Model struct {
	OriginalImage  *image.Gray
	ResampledImage *image.Gray
}

func NewModel() *Model {
	return &Model{
		OriginalImage:  nil,
		ResampledImage: nil,
	}
}

func (m *Model) GetOriginalImage() *image.Gray {
	return m.OriginalImage
}

func (m *Model) GetResampledImage() *image.Gray {
	return m.ResampledImage
}

func (m *Model) Resample(level int) {
	if m.OriginalImage == nil {
		fmt.Println("Original image is nil.")
		return
	}

	fmt.Println("Resampling with level:", level)

	ResampleImage(m.OriginalImage, m.ResampledImage, 8, level)
}

func (m *Model) SetImage(img *image.Gray) {
	m.OriginalImage = img
	m.ResampledImage = image.NewGray(img.Rect)

}

func (m *Model) CalculateMAE() string {

	mae, err := MeanAbsoluteError(m.OriginalImage.Pix, m.ResampledImage.Pix)

	if err != nil {
		fmt.Println("Error calculating MAE:", err)
		return "Error"
	}
	return fmt.Sprintf("%.2f", mae)
}

func (m *Model) CalculatePSNR() string {
	psnr, err := PSNR(m.OriginalImage.Pix, m.ResampledImage.Pix)
	if err != nil {
		fmt.Println("Error calculating PSNR:", err)
		return "Error"

	}
	return fmt.Sprintf("%.2f", psnr)
}
