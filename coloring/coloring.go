package coloring

import (
	"image/color"
	"math"
)

func (palette *Palette) Color(pixelValue float64, iterations int) PaletteColor {
    if palette == nil {
        return Black
    }

    paletteSize := palette.Len()
    segmentSize := (iterations + 1) / (paletteSize - 1)

    pixelValue = max(pixelValue - 1, 0)
    segmentNumber, segmentOffset := math.Modf(pixelValue / float64(segmentSize))

    segmentIdx := int(segmentNumber)
    firstColor := palette.colors[segmentIdx]
    secondColor := palette.colors[segmentIdx + 1]

    fr, fg, fb, _ := firstColor.RGBA()
    sr, sg, sb, _ := secondColor.RGBA()

    resultColor := color.RGBA{
        blend(fr, sr, segmentOffset),
        blend(fg, sg, segmentOffset),
        blend(fb, sb, segmentOffset),
        255,
    }

    return resultColor
}

func blend(first uint32, second uint32, offset float64) uint8 {
    fFirst := float64(first >> 8)
    fSecond := float64(second >> 8)

    return uint8(fFirst + (fSecond - fFirst) * offset)
}


