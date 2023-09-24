package coloring

import "image/color"

var (
    Black = color.Black
    White = color.RGBA{255, 255, 255, 255}
    Coral = color.RGBA{255, 127, 80, 255}
    Gold = color.RGBA{255, 215, 0, 255}
    Orangered = color.RGBA{255, 69, 0, 255}
    Blue = color.RGBA{0, 0, 255, 255}
)

var (
    BlackWhitePalette = Palette{
        colors: []PaletteColor{Black, White},
    }

    WhiteBlackPalette = Palette{
        colors: []PaletteColor{White, Black},
    }

    DefaultPalette = Palette{
        colors: []PaletteColor{White, Coral, Black},
    }
)

