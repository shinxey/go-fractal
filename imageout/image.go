package imageout

import (
	"image"
	"image/png"
	"os"
	"sync"

	"shinxey.go-fractal/coloring"
)

func WriteToFile(filename string, fractal []float64, width int, height int, iters int, palette coloring.Palette) {
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}

    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    var wg sync.WaitGroup
    wg.Add(width)

    for x := 0; x < width; x++ {
        go func(x int) {
            for y := 0; y < height; y++ {
                pixelIters := fractal[y * height + x]
                pixelColor := palette.Color(pixelIters, iters)

                img.Set(x, y, pixelColor)
            }
            wg.Done()
        }(x)
    }

    wg.Wait()

    f, _ := os.Create(filename)
    png.Encode(f, img)
}

