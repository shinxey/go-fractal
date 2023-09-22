package imageout

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func WriteToFile(filename string, fractal []int, width int, height int, iters int) {
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}

    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            var pixelColor color.Color

            pixelIters := fractal[y * height + x]

            if pixelIters >= iters {
                pixelColor = color.White
            } else if pixelIters >= iters / 2 {
                pixelColor = color.RGBA{120, 120, 120, 0}
            } else {
                pixelColor = color.Black
            }

            img.Set(x, y, pixelColor)
        }
    }

    f, _ := os.Create("image.png")
    png.Encode(f, img)
}

