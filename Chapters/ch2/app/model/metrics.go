package model

import (
	"fmt"
	"math"
)

func MeanSquaredError(x, y []uint8) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("The length of the two arrays is different")
	}

	var sum float64
	for i := 0; i < len(x); i++ {
		sum += math.Pow(float64(x[i]-y[i]), 2)
	}

	return sum / float64(len(x)), nil
}

func MeanAbsoluteError(x, y []uint8) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("The length of the two arrays is different")
	}

	var sum float64
	for i := 0; i < len(x); i++ {
		sum += math.Abs(float64(x[i] - y[i]))
	}

	return sum / float64(len(x)), nil
}

func PSNR(x, y []uint8) (float64, error) {
	mse, err := MeanSquaredError(x, y)
	if err != nil {
		return 0, err
	}

	if mse == 0 {
		return math.Inf(1), nil
	}

	return 10 * math.Log10(math.Pow(255, 2)/mse), nil
}
