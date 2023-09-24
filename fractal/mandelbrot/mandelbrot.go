package mandelbrot

func Draw(width int, height int, iters int) []float64 {
    result := make([]float64, width * height)

    fHeight := float64(height)
    fWidth := float64(width)

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            x0 := -1.167 + 0.0005 / fWidth * float64(x)
            y0 := 0.235 + 0.0005 / fHeight * float64(y)

            var xn float64 = 0
            var yn float64 = 0

            doneIters := 0
 
            for k := 0; k < iters; k++ {
                if xn * xn + yn * yn > 4 {
                    break
                }

                xn1 := xn * xn - yn * yn + x0
                yn1 := 2 * xn * yn + y0

                xn = xn1
                yn = yn1

                doneIters += 1
            }

            result[y * height + x] = float64(doneIters)
        }
    }

    return result
}

