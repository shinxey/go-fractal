package main

import (
	mandelbrot "shinxey.go-fractal/fractal"
	imageout "shinxey.go-fractal/imageout"
)

func main() {
    width := 2000
    height := 2000
    iters := 20

    result := mandelbrot.Draw(width, height, iters)

    imageout.WriteToFile("somefile.png", result, width, height, iters)
}

