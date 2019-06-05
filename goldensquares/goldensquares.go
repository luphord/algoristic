package goldensquares

import (
	"io"
	"math"
	"math/rand"

	svg "github.com/ajstarks/svgo"
)

var GoldenRatio = (1 + math.Sqrt(5)) / 2

func GoldenSquares(writer io.Writer, width uint32) {
	canvas := svg.New(writer)
	height := int32(float64(width) / GoldenRatio)
	canvas.Start(int(width), int(height))
	canvas.Rect(0, 0, int(width), int(height), "fill:grey")
	gen := &goldenSquareGenerator{&goldenSquare{0, 0, 0, 0}}
	for iter := gen; iter.HasNext(); {
		iter.Next()
	}
	canvas.End()
}

type goldenSquare struct {
	X      int
	Y      int
	Width  int
	Height int
}

type goldenSquareGenerator struct {
	next *goldenSquare
}

func (gen *goldenSquareGenerator) HasNext() bool {
	return gen.next != nil
}

func (gen *goldenSquareGenerator) Next() *goldenSquare {
	r := rand.Float64()
	if r < 0.3 {
		gen.next = nil
	} else {
		gen.next = &goldenSquare{1, 2, 3, 4}
	}
	return gen.next
}
