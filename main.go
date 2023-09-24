package main

import (
	"fmt"
	"os/exec"

	"shinxey.go-fractal/coloring"
	"shinxey.go-fractal/fractal/mandelbrot"
	"shinxey.go-fractal/imageout"
)

func main() {
    width := 2000
    height := 2000
    iters := byte(200)

    result := mandelbrot.Draw(width, height, iters)
    fmt.Print(result)

    imageName := "image.png"
    imageout.WriteToFile(imageName, result, width, height, iters, coloring.WhiteBlackPalette)

    cmd := exec.Command("open", imageName)
    cmd.Run()
}

