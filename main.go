package main

import (
	"os/exec"
	"runtime"

	"shinxey.go-fractal/coloring"
	"shinxey.go-fractal/fractal/mandelbrot"
	"shinxey.go-fractal/imageout"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    width := 4000
    height := 4000
    iters := 200

    result := mandelbrot.Draw(width, height, iters)

    imageName := "image.png"
    imageout.WriteToFile(imageName, result, width, height, iters, coloring.DefaultPalette)

    cmd := exec.Command("open", imageName)
    cmd.Run()
}

