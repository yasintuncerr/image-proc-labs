package model

import (
	"image"
	"math"
)

func normalize(src []uint8, dst []float64, level int) {

	for i, v := range src {
		dst[i] = float64(v) / float64(level)
	}
}

func quantize(src []float64, dst []uint8, level int) {

	for i, v := range src {
		dst[i] = uint8(math.Floor(v * float64(level)))
	}

}

func ResampleImage(src, dst *image.Gray, srcBitLevel, dstBitLevel int) {

	// Calculate the levels
	srcLevel := int(math.Pow(2, float64(srcBitLevel)))
	dstLevel := int(math.Pow(2, float64(dstBitLevel)))

	// allocate the normalized slices
	src_norm := make([]float64, len(src.Pix))
	dst_norm := make([]float64, len(dst.Pix))

	// allocate the quantized slices
	quantizedData := make([]uint8, len(src.Pix))

	//Quantization

	normalize(src.Pix, src_norm, srcLevel)

	quantize(src_norm, quantizedData, dstLevel)

	// Dequantization

	normalize(quantizedData, dst_norm, dstLevel)

	quantize(dst_norm, dst.Pix, srcLevel)

}
