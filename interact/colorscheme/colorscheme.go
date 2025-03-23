package colorscheme

import (
	"image/color"
)

type ColorScheme struct {
	Background color.RGBA
	Hover      color.RGBA
	Pressed    color.RGBA
	Border     color.RGBA
	Text       color.RGBA
}

func DefaultColorScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{211, 211, 211, 255}, // LightGray
		Hover:      color.RGBA{200, 200, 200, 255},
		Pressed:    color.RGBA{169, 169, 169, 255}, // DarkGray
		Border:     color.RGBA{0, 0, 0, 255},       // Black
		Text:       color.RGBA{0, 0, 0, 255},       // Black
	}
}

// Rainbow color schemes
func RedScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{255, 200, 200, 255},
		Hover:      color.RGBA{255, 150, 150, 255},
		Pressed:    color.RGBA{255, 100, 100, 255},
		Border:     color.RGBA{200, 0, 0, 255},
		Text:       color.RGBA{100, 0, 0, 255},
	}
}

func OrangeScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{255, 220, 180, 255},
		Hover:      color.RGBA{255, 200, 140, 255},
		Pressed:    color.RGBA{255, 180, 100, 255},
		Border:     color.RGBA{230, 140, 0, 255},
		Text:       color.RGBA{150, 80, 0, 255},
	}
}

func YellowScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{255, 255, 200, 255},
		Hover:      color.RGBA{255, 255, 150, 255},
		Pressed:    color.RGBA{255, 255, 100, 255},
		Border:     color.RGBA{200, 200, 0, 255},
		Text:       color.RGBA{100, 100, 0, 255},
	}
}

func GreenScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{200, 255, 200, 255},
		Hover:      color.RGBA{150, 255, 150, 255},
		Pressed:    color.RGBA{100, 255, 100, 255},
		Border:     color.RGBA{0, 200, 0, 255},
		Text:       color.RGBA{0, 100, 0, 255},
	}
}

func BlueScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{200, 200, 255, 255},
		Hover:      color.RGBA{150, 150, 255, 255},
		Pressed:    color.RGBA{100, 100, 255, 255},
		Border:     color.RGBA{0, 0, 200, 255},
		Text:       color.RGBA{0, 0, 100, 255},
	}
}

func IndigoScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{200, 180, 255, 255},
		Hover:      color.RGBA{180, 150, 255, 255},
		Pressed:    color.RGBA{160, 120, 255, 255},
		Border:     color.RGBA{75, 0, 130, 255},
		Text:       color.RGBA{40, 0, 80, 255},
	}
}

func VioletScheme() ColorScheme {
	return ColorScheme{
		Background: color.RGBA{230, 190, 255, 255},
		Hover:      color.RGBA{220, 150, 255, 255},
		Pressed:    color.RGBA{200, 100, 255, 255},
		Border:     color.RGBA{148, 0, 211, 255},
		Text:       color.RGBA{75, 0, 130, 255},
	}
}
