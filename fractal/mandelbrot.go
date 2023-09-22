package mandelbrot

func Draw(width int, height int, iters int) []int {
    result := make([]int, width * height)

    fHeight := float32(height)
    fWidth := float32(width)

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            x0 := -2 + 4 / fWidth * float32(x)
            y0 := -2 + 4 / fHeight * float32(y)

            var xn float32 = 0
            var yn float32 = 0

            doneIters := 0
 
            for k := 0; k < iters; k++ {
                xn1 := xn * xn - yn * yn + x0
                yn1 := 2 * xn * yn + y0

                xn = xn1
                yn = yn1

                doneIters += 1

                if xn * xn + yn + yn > 4 {
                    break
                }
            }

            result[y * height + x] = doneIters
        }
    }

    return result
}

