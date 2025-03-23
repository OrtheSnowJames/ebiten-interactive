package interact

import (
	"golang.org/x/image/font"

	"github.com/OrtheSnowJames/ebiten-interactive/interact/button"
	"github.com/OrtheSnowJames/ebiten-interactive/interact/clip"
	"github.com/OrtheSnowJames/ebiten-interactive/interact/textfield"
	"github.com/OrtheSnowJames/ebiten-interactive/interact/colorscheme"
	"github.com/hajimehoshi/ebiten/v2"
)

// An interface that all interactive objects must conform to
type InteractiveObject interface {
	Update()
	Draw(screen *ebiten.Image)
}

func NewButton(x, y, width, height float32, text string) *button.Button {
	return button.NewButton(x, y, width, height, text)
}

func NewRoundedButton(x, y, width, height float32, text string, cornerRadius float32) *button.Button {
	btn := button.NewButton(x, y, width, height, text)
	btn.SetRoundedCorners(true)
	btn.SetCornerRadius(cornerRadius)
	return btn
}

func NewPointyButton(x, y, width, height float32, text string, pointyAmount float32) *button.Button {
	btn := button.NewButton(x, y, width, height, text)
	btn.SetPointyStyle(true)
	btn.SetPointyAmount(pointyAmount)
	return btn
}

func NewNormalButton(x, y, width, height float32, text string) *button.Button {
	btn := button.NewButton(x, y, width, height, text)
	btn.SetRoundedCorners(false)
	return btn
}

func NewTextField(x, y, width, height float32, maxLength int) *textfield.TextField {
	return textfield.NewTextField(x, y, width, height, maxLength)
}

func NewTextFieldWithPlaceholder(x, y, width, height float32, maxLength int, placeholder string) *textfield.TextField {
	tf := textfield.NewTextField(x, y, width, height, maxLength)
	tf.SetPlaceholder(placeholder)
	return tf
}

func CopyClip(text string) error {
	return clip.CopyClip(text)
}

func PasteClip() (string, error) {
	return clip.PasteClip()
}

func LoadFontFace(fontPath string, size float64) (font.Face, error) {
	return clip.LoadFontFace(fontPath, size)
}

func DefaultColorScheme() colorscheme.ColorScheme {
	return colorscheme.DefaultColorScheme()
}

// Can draw multiple objects at once instead of doing .draw .draw ..., call like this: interact.DrawAll(screen, obj1, obj2, obj3)
func DrawAll(screen *ebiten.Image, objects ...InteractiveObject) {
	for _, obj := range objects {
		obj.Draw(screen)
	}
}

// Can update multiple objects at once instead of doing .update .update ..., call like this: interact.UpdateAll(obj1, obj2, obj3)
func UpdateAll(objects ...InteractiveObject) {
	for _, obj := range objects {
		obj.Update()
	}
}
