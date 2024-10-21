package model

import (
	"fmt"
	"math"
	"sync"
)

func MeanSquaredError(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("The length of the two arrays is different")
	}

	var sum float64
	for i := 0; i < len(x); i++ {
		sum += math.Pow(x[i]-y[i], 2)
	}

	return sum / float64(len(x)), nil
}

func MeanAbsoluteError(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("The length of the two arrays is different")
	}

	var sum float64
	for i := 0; i < len(x); i++ {
		sum += math.Abs(x[i] - y[i])
	}

	return sum / float64(len(x)), nil
}

func PSNR(x, y []float64) (float64, error) {
	mse, err := MeanSquaredError(x, y)
	if err != nil {
		return 0, err
	}

	if mse == 0 {
		return 0, fmt.Errorf("The Mean Squared Error is zero")
	}

	return 10 * math.Log10(math.Pow(255, 2)/mse), nil
}

func ParallelMeanSquaredError(x, y []float64, concurrencyCount int) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("The length of the two arrays is different")
	}

	if concurrencyCount <= 0 {
		return 0, fmt.Errorf("The concurrency count must be greater than 0")
	}

	var sum float64
	var mu sync.Mutex // Mutex to avoid race conditions when summing
	var wg sync.WaitGroup
	wg.Add(concurrencyCount)

	chunkSize := len(x) / concurrencyCount

	for i := 0; i < concurrencyCount; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if i == concurrencyCount-1 {
				end = len(x) // Handle the last chunk to cover the remaining elements
			}
			localSum := 0.0
			for j := start; j < end; j++ {
				localSum += math.Pow(x[j]-y[j], 2)
			}
			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	return sum / float64(len(x)), nil
}

func ParallelMeanAbsoluteError(x, y []float64, concurrencyCount int) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("The length of the two arrays is different")
	}

	if concurrencyCount <= 0 {
		return 0, fmt.Errorf("The concurrency count must be greater than 0")
	}

	var sum float64
	var mu sync.Mutex // Mutex to avoid race conditions when summing
	var wg sync.WaitGroup
	wg.Add(concurrencyCount)

	chunkSize := len(x) / concurrencyCount

	for i := 0; i < concurrencyCount; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if i == concurrencyCount-1 {
				end = len(x) // Handle the last chunk to cover the remaining elements
			}
			localSum := 0.0
			for j := start; j < end; j++ {
				localSum += math.Abs(x[j] - y[j])
			}
			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	return sum / float64(len(x)), nil
}
