package controller

import (
	"image"
	"image/color"
	"image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	model "github.com/yasintuncerr/image-proc-labs/Chapters/ch2/model"
	view "github.com/yasintuncerr/image-proc-labs/Chapters/ch2/view"
)

func IsImageEmpty(img *image.Gray) bool {
	if len(img.Pix) == 0 {
		return true
	}

	firstPixel := img.Pix[0]

	for _, pix := range img.Pix {
		if pix != firstPixel {
			return false
		}
	}

	return true
}

type Controller struct {
	MainLayout *view.MainLayout
	Model      *model.Model
	Window     fyne.Window
}

func NewController(mainLayout *view.MainLayout, model *model.Model, window fyne.Window) *Controller {
	return &Controller{
		MainLayout: mainLayout,
		Model:      model,
		Window:     window,
	}
}

func (c *Controller) Init() {

	// Set the Sidebar ApplyBtn OnTapped function
	c.MainLayout.Sidebar.ApplyBtn.OnTapped = func() {

		if c.Model.GetOriginalImage() != nil && !IsImageEmpty(c.Model.GetOriginalImage()) {
			level := int(c.MainLayout.Sidebar.Slider.Value)
			c.Model.Resample(level)
			c.MainLayout.QuantizedImageCanvas.UpdateImage(c.Model.GetResampledImage())
			mae := c.Model.CalculateMAE()
			psnr := c.Model.CalculatePSNR()
			c.MainLayout.Sidebar.MAELabel.SetText("MAE: " + mae)
			c.MainLayout.Sidebar.PSNRLabel.SetText("PSNR: " + psnr)

			c.MainLayout.QuantizedImageCanvas.UpdateImage(c.Model.GetResampledImage())

		}
	}

	// Set the Sidebar OpenFileBtn OnTapped function
	c.MainLayout.Sidebar.OpenFileBtn.OnTapped = func() {
		filedialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, c.Window)
				return
			}
			if reader == nil {
				return
			}

			img, _, err := image.Decode(reader)

			if err != nil {
				dialog.ShowError(err, c.Window)
				return
			}

			grayImg, ok := img.(*image.Gray)
			if !ok {
				grayImg = image.NewGray(img.Bounds())
			}

			for y := 0; y < img.Bounds().Dy(); y++ {
				for x := 0; x < img.Bounds().Dx(); x++ {
					originalColor := img.At(x, y)
					grayColor := color.GrayModel.Convert(originalColor).(color.Gray)
					grayImg.Set(x, y, grayColor)
				}
			}

			c.Model.SetImage(grayImg)
			c.MainLayout.OriginalImageCanvas.UpdateImage(c.Model.GetOriginalImage())
			c.MainLayout.QuantizedImageCanvas.UpdateImage(c.Model.GetResampledImage())
		}, c.Window)
		filedialog.Show()
	}

	c.MainLayout.Sidebar.ExportBtn.OnTapped = func() {
		filedialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, c.Window)
				return
			}
			if writer == nil {
				return
			}

			img := c.Model.GetResampledImage()

			err = png.Encode(writer, img)
			if err != nil {
				dialog.ShowError(err, c.Window)
				writer.Close()
				return
			}

			writer.Close()

		}, c.Window)

		filedialog.Show()
	}
}
