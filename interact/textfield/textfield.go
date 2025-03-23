// SPDX-License-Identifier: MIT
package textfield

import (
	"fmt"
	"image/color"

	"github.com/OrtheSnowJames/ebiten-interactive/interact/clip"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Rect defines a rectangle with float32 coordinates.
type Rect struct {
	X, Y, W, H float32
}

func NewRect(x, y, w, h float32) Rect {
	return Rect{x, y, w, h}
}

func pointInRect(x, y float32, r Rect) bool {
	return x >= r.X && x <= r.X+r.W && y >= r.Y && y <= r.Y+r.H
}

// DefaultFont is a package-level font face used for drawing text.
// Set this to a valid font.Face during initialization.
var DefaultFont font.Face

type TextField struct {
	Bounds             Rect
	Text               string
	MaxLength          int
	BackgroundColor    color.RGBA
	BorderColor        color.RGBA
	TextColor          color.RGBA
	FontSize           int32
	FontFace           font.Face
	IsActive           bool
	CursorPosition     int
	CursorBlinkTimer   float32
	BackspaceHoldTimer float32
	Invisible          bool
	Uneditable         bool
	Placeholder        string
}

func NewTextField(x, y, width, height float32, maxLength int) *TextField {
	return &TextField{
		Bounds:             NewRect(x, y, width, height),
		Text:               "",
		MaxLength:          maxLength,
		BackgroundColor:    color.RGBA{R: 255, G: 255, B: 255, A: 255}, // White
		BorderColor:        color.RGBA{R: 0, G: 0, B: 0, A: 255},       // Black
		TextColor:          color.RGBA{R: 0, G: 0, B: 0, A: 255},       // Black
		FontSize:           20,
		FontFace:           DefaultFont,
		IsActive:           false,
		CursorPosition:     0,
		CursorBlinkTimer:   0.0,
		BackspaceHoldTimer: 0.0,
		Invisible:          false,
		Uneditable:         false,
		Placeholder:        "",
	}
}

func (tf *TextField) SetColors(background, border, text color.RGBA) {
	tf.BackgroundColor = background
	tf.BorderColor = border
	tf.TextColor = text
}

func (tf *TextField) SetFontSize(fontSize int32) {
	tf.FontSize = fontSize
}

func (tf *TextField) SetInvisible(invisible bool) {
	tf.Invisible = invisible
}

func (tf *TextField) IsInvisible() bool {
	return tf.Invisible
}

func (tf *TextField) SetUneditable(uneditable bool) {
	tf.Uneditable = uneditable
}

func (tf *TextField) IsUneditable() bool {
	return tf.Uneditable
}

// Update should be called every frame.
func (tf *TextField) Update() {
	if tf.Uneditable {
		return
	}

	// Update cursor blink timer (assuming 60 FPS)
	tf.CursorBlinkTimer += 1.0 / 60.0
	if tf.CursorBlinkTimer >= 1.0 {
		tf.CursorBlinkTimer = 0.0
	}

	// Handle mouse click to activate/deactivate the text field.
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if pointInRect(float32(mx), float32(my), tf.Bounds) {
			tf.IsActive = true
		} else {
			tf.IsActive = false
		}
	}

	if tf.IsActive {
		// Append typed characters.
		for _, ch := range ebiten.InputChars() {
			if len(tf.Text) < tf.MaxLength {
				tf.Text = tf.Text[:tf.CursorPosition] + string(ch) + tf.Text[tf.CursorPosition:]
				tf.CursorPosition++
			}
		}

		// Handle backspace (single press).
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if tf.CursorPosition > 0 {
				tf.Text = tf.Text[:tf.CursorPosition-1] + tf.Text[tf.CursorPosition:]
				tf.CursorPosition--
			}
		}

		// Handle left arrow.
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && tf.CursorPosition > 0 {
			tf.CursorPosition--
		}

		// Handle right arrow.
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && tf.CursorPosition < len(tf.Text) {
			tf.CursorPosition++
		}

		// Handle home key.
		if inpututil.IsKeyJustPressed(ebiten.KeyHome) {
			tf.CursorPosition = 0
		}

		// Handle end key.
		if inpututil.IsKeyJustPressed(ebiten.KeyEnd) {
			tf.CursorPosition = len(tf.Text)
		}

		// Handle continuous backspace hold.
		if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
			tf.BackspaceHoldTimer += 1.0 / 60.0
			if tf.BackspaceHoldTimer > 0.5 {
				tf.BackspaceHoldTimer = 1.0
				if tf.CursorPosition > 0 {
					tf.Text = tf.Text[:tf.CursorPosition-1] + tf.Text[tf.CursorPosition:]
					tf.CursorPosition--
				}
			}
		} else {
			tf.BackspaceHoldTimer = 0.0
		}

		// Handle Control+A (select all).
		if inpututil.IsKeyJustPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyControl) {
			tf.CursorPosition = len(tf.Text)
		}

		// Handle Control+V (paste).
		if inpututil.IsKeyJustPressed(ebiten.KeyV) && ebiten.IsKeyPressed(ebiten.KeyControl) {
			clipboardText, err := clip.PasteClip()

			if err != nil {
				panic(fmt.Errorf("failed to paste text from clipboard (from ebiten-interactive textfield): %w", err))
			}

			if clipboardText != "" {
				remainingSpace := tf.MaxLength - len(tf.Text)
				if remainingSpace > 0 {
					toPaste := clipboardText
					if len(toPaste) > remainingSpace {
						toPaste = toPaste[:remainingSpace]
					}
					tf.Text = tf.Text[:tf.CursorPosition] + toPaste + tf.Text[tf.CursorPosition:]
					tf.CursorPosition += len(toPaste)
				}
			}
		}
	}
}

// Draw draws the text field onto the given screen.
func (tf *TextField) Draw(screen *ebiten.Image) {
	if tf.Invisible {
		return
	}

	// Draw background.
	ebitenutil.DrawRect(screen, float64(tf.Bounds.X), float64(tf.Bounds.Y), float64(tf.Bounds.W), float64(tf.Bounds.H), tf.BackgroundColor)

	// Border color: red if active.
	drawBorderColor := tf.BorderColor
	if tf.IsActive {
		drawBorderColor = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	}
	drawRectOutline(screen, tf.Bounds, 2.0, drawBorderColor)

	// Draw text with a small padding.
	padding := float32(5)
	if tf.FontFace == nil {
		return
	}
	textY := tf.Bounds.Y + (tf.Bounds.H-float32(tf.FontSize))/2.0

	// Draw either the text or placeholder
	displayText := tf.Text
	textColor := tf.TextColor
	if displayText == "" && tf.Placeholder != "" {
		displayText = tf.Placeholder
		textColor = color.RGBA{R: 128, G: 128, B: 128, A: 255} // Gray color for placeholder
	}
	text.Draw(screen, displayText, tf.FontFace, int(tf.Bounds.X)+int(padding), int(textY)+int(tf.FontSize), textColor)

	// Draw the cursor if active and during the blink phase.
	if tf.IsActive && tf.CursorBlinkTimer < 0.5 {
		subText := tf.Text[:tf.CursorPosition]
		bounds := text.BoundString(tf.FontFace, subText)
		textWidth := float32(bounds.Dx())
		cursorX := tf.Bounds.X + padding + textWidth
		ebitenutil.DrawLine(screen, float64(cursorX), float64(textY), float64(cursorX), float64(textY+float32(tf.FontSize)), tf.TextColor)
	}
}

func drawRectOutline(screen *ebiten.Image, r Rect, thickness float32, col color.RGBA) {
	// Top line.
	ebitenutil.DrawLine(screen, float64(r.X), float64(r.Y), float64(r.X+r.W), float64(r.Y), col)
	// Bottom line.
	ebitenutil.DrawLine(screen, float64(r.X), float64(r.Y+r.H), float64(r.X+r.W), float64(r.Y+r.H), col)
	// Left line.
	ebitenutil.DrawLine(screen, float64(r.X), float64(r.Y), float64(r.X), float64(r.Y+r.H), col)
	// Right line.
	ebitenutil.DrawLine(screen, float64(r.X+r.W), float64(r.Y), float64(r.X+r.W), float64(r.Y+r.H), col)
}

func (tf *TextField) GetText() string {
	return tf.Text
}

func (tf *TextField) SetValue(value string) {
	if len(value) > tf.MaxLength {
		tf.Text = value[:tf.MaxLength]
	} else {
		tf.Text = value
	}
	tf.CursorPosition = len(tf.Text)
}

func (tf *TextField) Activate() {
	tf.IsActive = true
	tf.CursorPosition = len(tf.Text)
}

func (tf *TextField) Deactivate() {
	tf.IsActive = false
}

// Add a method to set the placeholder text
func (tf *TextField) SetPlaceholder(placeholder string) {
	tf.Placeholder = placeholder
}

func (tf *TextField) SetFont(face font.Face) {
	tf.FontFace = face
}
