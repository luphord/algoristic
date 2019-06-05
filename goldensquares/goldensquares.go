package goldensquares

import (
	"io"
	"math"

	"github.com/ajstarks/svgo"
)

var GoldenRatio = (1 + math.Sqrt(5)) / 2

func GoldenSquares(writer io.Writer, width uint32) {
	canvas := svg.New(writer)
	height := int32(float64(width) / GoldenRatio)
	canvas.Start(int(width), int(height))
	canvas.Rect(0, 0, int(width), int(height), "fill:grey")
	canvas.End()
}
