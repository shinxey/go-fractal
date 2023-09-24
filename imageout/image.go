package imageout

import (
	"image"
	"image/png"
	"os"

	"shinxey.go-fractal/coloring"
)

func WriteToFile(filename string, fractal []byte, width int, height int, iters byte, palette coloring.Palette) {
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}

    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            pixelIters := fractal[y * height + x]
            pixelColor := palette.Color3(pixelIters, iters)

            img.Set(x, y, pixelColor)
        }
    }

    f, _ := os.Create(filename)
    png.Encode(f, img)
}

