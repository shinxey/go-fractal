package coloring

import (
	"image/color"
	"math"
)

func (palette *Palette) Color(pixelValue byte, iterations byte) PaletteColor {
    paletteLen := palette.Len()

    if palette == nil {
        return color.Black
    }

    if paletteLen < 1 {
        return color.Black
    }

    iterationStep := float32(paletteLen) / float32(iterations)
    colorOffsetFloat := float32(pixelValue) * iterationStep
    colorOffset := int(math.Floor(float64(colorOffsetFloat))) - 1
    colorOffset = max(colorOffset, 0)

    return palette.colors[colorOffset]
}

func (palette *Palette) Color2(pixelValue byte, iterations byte) PaletteColor {
    if palette == nil {
        return Black
    }

    rawOffset := math.Log2(float64(pixelValue))

    upperOffset := math.Floor(rawOffset + 1)
    lowerOffset := math.Floor(rawOffset)

    offsetPortion := rawOffset - lowerOffset
    lastColorIdx := palette.Len() - 1
    lr, lg, lb, _ := palette.colors[min(lastColorIdx, int(lowerOffset))].RGBA()
    ur, ug, ub, _ := palette.colors[min(lastColorIdx, int(upperOffset))].RGBA()
    resultColor := color.RGBA{
        uint8(min(lr, ur)) + uint8(math.Abs(float64(ur - lr) * offsetPortion)),
        uint8(min(lg, ug)) + uint8(math.Abs(float64(ug - lg) * offsetPortion)),
        uint8(min(lb, ub)) + uint8(math.Abs(float64(ub - lb) * offsetPortion)),
        255,
    }

    return resultColor

    //colorOffset := min(palette.Len() - 1, int(math.Log2(float64(pixelValue))))

    //return palette.colors[colorOffset]
}

func (palette *Palette) Color3(pixelValue byte, iterations byte) PaletteColor {
    if palette == nil {
        return Black
    }

    paletteSize := palette.Len()
    segmentSize := int(iterations) / (paletteSize - 1)

    pixelValue = max(pixelValue - 1, 0)
    segmentNumber := pixelValue / byte(segmentSize)
    segmentOffset := float64(pixelValue % byte(segmentSize)) / float64(segmentSize)
    
    firstColor := palette.colors[segmentNumber]
    secondColor := palette.colors[segmentNumber + 1]

    fr, fg, fb, _ := firstColor.RGBA()
    sr, sg, sb, _ := secondColor.RGBA()

    
    resultColor := color.RGBA{
        uint8(min(fr, sr)) + uint8(math.Abs(float64(fr - sr) * segmentOffset)),
        uint8(min(fg, sg)) + uint8(math.Abs(float64(fg - sg) * segmentOffset)),
        uint8(min(fb, sb)) + uint8(math.Abs(float64(fb - sb) * segmentOffset)),
        255,
    }

    return resultColor
}

