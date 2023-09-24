package mandelbrot

import "sync"

func Draw(width int, height int, iters int) []float64 {
    result := make([]float64, width * height)

    fHeight := float64(height)
    fWidth := float64(width)

    var wg sync.WaitGroup
    wg.Add(width)

    for x := 0; x < width; x++ {
        go func(x int) {
            for y := 0; y < height; y++ {
                z0 := complex(
                    -1.167 + 0.0005 / fWidth * float64(x), 
                    0.235 + 0.0005 / fHeight * float64(y),
                )

                zn := z0

                doneIters := 0

                for (real(zn) * real(zn) + imag(zn) * imag(zn) <= 4) && (doneIters < iters) {
                    zn = zn * zn + z0

                    doneIters += 1
                }

                result[y * height + x] = float64(doneIters)
            }
            wg.Done()
        }(x)
    }

    wg.Wait()

    return result
}

