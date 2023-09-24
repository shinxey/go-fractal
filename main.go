package main

import (
	"os/exec"

	"shinxey.go-fractal/coloring"
	"shinxey.go-fractal/fractal/mandelbrot"
	"shinxey.go-fractal/imageout"
)

func main() {
    width := 2000
    height := 2000
    iters := 200

    result := mandelbrot.Draw(width, height, iters)

    imageName := "image.png"
    imageout.WriteToFile(imageName, result, width, height, iters, coloring.DefaultPalette)

    cmd := exec.Command("open", imageName)
    cmd.Run()
}

