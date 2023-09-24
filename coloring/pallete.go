package coloring

import "image/color"

type PaletteColor color.Color

type Palette struct {
    colors []PaletteColor 
}

func (palette *Palette) Len() int {
    return len(palette.colors)
}

