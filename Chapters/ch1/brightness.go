package imageprocessing

import (
	"fmt"
	"image"
)

const (
	maxUint8  = 255
	minUint8  = 0
	maxUint16 = 65535
	minUint16 = 0
	maxUint32 = 4294967295
	minUint32 = 0
)

func getMinMax[T uint8 | uint16 | uint32]() (min, max int) {
	switch any(*new(T)).(type) {
	case uint8:
		return minUint8, maxUint8
	case uint16:
		return minUint16, maxUint16
	case uint32:
		return minUint32, maxUint32
	}
	return 0, 0
}

func sum[T uint8 | uint16 | uint32](a []T, length, stride, val int) error {
	minVal, maxVal := getMinMax[T]()

	for i := 0; i < length; i += stride {
		newVal := int(a[i]) + val

		if newVal > maxVal {
			return fmt.Errorf("overflow: %d exceeds max for type", newVal)
		}
		if newVal < minVal {
			return fmt.Errorf("underflow: %d below min for type", newVal)
		}

		a[i] = T(newVal)
	}
	return nil
}

func Fill[T any](a []T, val T) {
	for i := range a {
		a[i] = val
	}
}

func AdjustBrightness[T image.Gray | image.RGBA | image.Gray16 | image.RGBA64](img *T, val int) error {
	switch castedImage := any(img).(type) {
	case *image.Gray:
		return sum(castedImage.Pix, len(castedImage.Pix), castedImage.Stride, val)
	case *image.Gray16:
		return sum(castedImage.Pix, len(castedImage.Pix), castedImage.Stride, val)
	case *image.RGBA:
		channels := make(chan error, 3)
		for i := 0; i < 3; i++ {
			go func(i int) {
				channels <- sum(castedImage.Pix[i:], len(castedImage.Pix), 4, val)
			}(i)
		}
		for i := 0; i < 3; i++ {
			if err := <-channels; err != nil {
				return err
			}
		}
		return nil
	case *image.RGBA64:
		channels := make(chan error, 3)
		for i := 0; i < 3; i++ {
			go func(i int) {
				channels <- sum(castedImage.Pix[i:], len(castedImage.Pix), 8, val)
			}(i)
		}
		for i := 0; i < 3; i++ {
			if err := <-channels; err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("unsupported image type %T", img)

}
