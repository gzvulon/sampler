package console

import (
	"fmt"
	ui "github.com/sqshq/termui"
)

type Theme string

const (
	ThemeDark  Theme = "dark"
	ThemeLight Theme = "light"
)

const (
	ColorOlive       ui.Color = 178
	ColorDeepSkyBlue ui.Color = 39
	ColorDeepPink    ui.Color = 198
	ColorCian        ui.Color = 43
	ColorOrange      ui.Color = 166
	ColorPurple      ui.Color = 129
	ColorGreen       ui.Color = 64
	ColorDarkGrey    ui.Color = 235
	ColorGrey        ui.Color = 242
	ColorWhite       ui.Color = 15
	ColorBlack       ui.Color = 0
	ColorClear       ui.Color = -1
)

const (
	MenuColorBackground ui.Color = 235
	MenuColorText       ui.Color = 255
)

type Palette struct {
	Colors []ui.Color
	// TODO Menu colors, like Dark, Medium, Light etc
}

func GetPalette(theme Theme) Palette {
	switch theme {
	case ThemeDark:
		return Palette{Colors: []ui.Color{ColorOlive, ColorDeepSkyBlue, ColorDeepPink, ColorWhite, ColorGrey, ColorGreen, ColorOrange, ColorCian, ColorPurple}}
	case ThemeLight:
		return Palette{Colors: []ui.Color{ColorOlive, ColorDeepSkyBlue, ColorDeepPink, ColorWhite, ColorGrey, ColorGreen, ColorOrange, ColorCian, ColorPurple}}
	default:
		panic(fmt.Sprintf("Following theme is not supported: %v", theme))
	}
}
