package model

import (
	"fmt"
	"math"
	"sync"
)

type BitLevel uint16

const (
	Bit8 BitLevel = 256
	Bit7 BitLevel = 128
	Bit6 BitLevel = 64
	Bit5 BitLevel = 32
	Bit4 BitLevel = 16
	Bit3 BitLevel = 8
	Bit2 BitLevel = 4
	Bit1 BitLevel = 2
)

func getMaxValue(data []uint8) uint16 {
	var max uint16 = 0
	for i := 0; i < len(data); i++ {
		if uint16(data[i]) > max {
			max = uint16(data[i])
		}
	}
	return max
}

func DetectBitLevel(data []uint8) BitLevel {
	maxValue := getMaxValue(data)
	switch {
	case maxValue < 2:
		return Bit1
	case maxValue < 4:
		return Bit2
	case maxValue < 8:
		return Bit3
	case maxValue < 16:
		return Bit4
	case maxValue < 32:
		return Bit5
	case maxValue < 64:
		return Bit6
	case maxValue < 128:
		return Bit7
	default:
		return Bit8
	}
}

func NormalizeData(src []uint8, dst []float64, bitLevel BitLevel, start, length int) error {
	if src == nil {
		return fmt.Errorf("input array is nil")
	}
	if dst == nil {
		return fmt.Errorf("output array is nil")
	}

	for i := start; i < len(src) && i < start+length; i++ {
		dst[i] = float64(src[i]) / float64(bitLevel)
	}
	return nil
}

func QuantizeData(src []float64, dst []uint8, bitLevel BitLevel, start, length int) error {
	if src == nil {
		return fmt.Errorf("input array is nil")
	}
	if dst == nil {
		return fmt.Errorf("output array is nil")
	}

	for i := start; i < len(src) && i < start+length; i++ {
		dst[i] = uint8(math.Floor(float64(bitLevel) * src[i]))
	}
	return nil
}

func ParallelNormalize(src []uint8, dst []float64, bitLevel BitLevel, concurrencyCount int) error {
	if getMaxValue(src) > uint16(bitLevel) {
		return fmt.Errorf("maximum value of the input array is greater than the bit level")
	}

	var wg sync.WaitGroup
	wg.Add(concurrencyCount)

	chunkSize := len(src) / concurrencyCount
	for i := 0; i < concurrencyCount; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if i == concurrencyCount-1 {
				end = len(src)
			}
			NormalizeData(src, dst, bitLevel, start, end-start)
		}(i)
	}
	wg.Wait()

	return nil
}

func ParallelQuantize(src []float64, dst []uint8, bitLevel BitLevel, concurrencyCount int) error {
	var wg sync.WaitGroup
	wg.Add(concurrencyCount)

	chunkSize := len(src) / concurrencyCount
	for i := 0; i < concurrencyCount; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if i == concurrencyCount-1 {
				end = len(src)
			}
			QuantizeData(src, dst, bitLevel, start, end-start)
		}(i)
	}
	wg.Wait()

	return nil
}

func Resample(src []uint8, dst []uint8, bitLevel BitLevel) error {
	if src == nil {
		return fmt.Errorf("input array is nil")
	}
	if dst == nil {
		return fmt.Errorf("output array is nil")
	}

	if len(src) != len(dst) {
		return fmt.Errorf("input and output arrays have different lengths")
	}

	normalizedData := make([]float64, len(src))
	if err := NormalizeData(src, normalizedData, bitLevel, 0, len(src)); err != nil {
		return err
	}

	if err := QuantizeData(normalizedData, dst, bitLevel, 0, len(src)); err != nil {
		return err
	}

	return nil
}

func ParallelResample(src []uint8, dst []uint8, bitLevel BitLevel, concurrencyCount int) error {
	if getMaxValue(src) > uint16(bitLevel) {
		return fmt.Errorf("maximum value of the input array is greater than the bit level")
	}

	var wg sync.WaitGroup
	var err error
	var mu sync.Mutex

	wg.Add(concurrencyCount)
	chunkSize := len(src) / concurrencyCount

	for i := 0; i < concurrencyCount; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if i == concurrencyCount-1 {
				end = len(src)
			}
			if resampleErr := Resample(src[start:end], dst[start:end], bitLevel); resampleErr != nil {
				mu.Lock()
				err = resampleErr
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	return err
}
