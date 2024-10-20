package main

import (
	"image"
	"testing"
)

func TestFillUint8(t *testing.T) {
	a := make([]uint8, 10)
	Fill(a, 1)

	for i := 0; i < len(a); i++ {
		if a[i] != 1 {
			t.Errorf("Expected 1, got %d", a[i])
		}
	}
}

func TestFillUint16(t *testing.T) {
	a := make([]uint16, 10)
	Fill(a, 1)

	for i := 0; i < len(a); i++ {
		if a[i] != 1 {
			t.Errorf("Expected 1, got %d", a[i])
		}
	}
}

func TestFillUint32(t *testing.T) {
	a := make([]uint32, 10)
	Fill(a, 1)

	for i := 0; i < len(a); i++ {
		if a[i] != 1 {
			t.Errorf("Expected 1, got %d", a[i])
		}
	}
}

func TestsumUint8(t *testing.T) {
	a := make([]uint8, 10)
	Fill(a, 0)
	err := sum(a, len(a), 1, 1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != 1 {
			t.Errorf("Expected 1, got %d", a[i])
		}
	}
}

func TestsumUint16(t *testing.T) {
	a := make([]uint16, 10)
	Fill(a, 0)
	err := sum(a, len(a), 1, 1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != 1 {
			t.Errorf("Expected 1, got %d", a[i])
		}
	}
}

func TestsumUint32(t *testing.T) {
	a := make([]uint32, 10)
	Fill(a, 0)
	err := sum(a, len(a), 1, 1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != 1 {
			t.Errorf("Expected 1, got %d", a[i])
		}
	}
}

func TestAdjustBrightnessGray(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 2, 2))

	Fill(img.Pix, 0)

	err := AdjustBrightness(img, 1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	for i := 0; i < len(img.Pix); i++ {
		if img.Pix[i] != 1 {
			t.Errorf("Expected 1, got %d", img.Pix[i])
		}
	}
}

func TestAdjustBrightnessGray16(t *testing.T) {
	img := image.NewGray16(image.Rect(0, 0, 2, 2))

	Fill(img.Pix, 0)

	err := AdjustBrightness(img, 1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	for i := 0; i < len(img.Pix); i++ {
		if img.Pix[i] != 1 {
			t.Errorf("Expected 1, got %d", img.Pix[i])
		}
	}
}

func TestAdjustBrightnessRGBA(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	Fill(img.Pix, 0)

	err := AdjustBrightness(img, 1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	for i := 0; i < len(img.Pix); i++ {
		if (i+1)%4 == 0 {
			continue
		}

		if img.Pix[i] != 1 {
			t.Errorf("Expected 1, got %d", img.Pix[i])
		}
	}
}

func TestAdjustBrightnessOverflow(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 2, 2))

	Fill(img.Pix, 255)

	err := AdjustBrightness(img, 1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAdjustBrightnessUnderflow(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 2, 2))

	Fill(img.Pix, 0)

	err := AdjustBrightness(img, -1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAdjustBrightnessRGBAOverflow(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	Fill(img.Pix, 255)

	err := AdjustBrightness(img, 1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAdjustBrightnessRGBAUnderflow(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	Fill(img.Pix, 0)

	err := AdjustBrightness(img, -1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
