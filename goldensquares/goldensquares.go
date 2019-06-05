package goldensquares

import (
	"math"
	"os"

	"github.com/ajstarks/svgo"
)

var GoldenRatio = (1 + math.Sqrt(5)) / 2

func TestSvg() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}
