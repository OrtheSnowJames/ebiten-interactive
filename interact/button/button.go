// SPDX-License-Identifier: MIT
package button

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func lerpColor(c1, c2 color.RGBA, t float32) color.RGBA {
	return color.RGBA{
		R: uint8(float32(c2.R-c1.R)*t + float32(c1.R)),
		G: uint8(float32(c2.G-c1.G)*t + float32(c1.G)),
		B: uint8(float32(c2.B-c1.B)*t + float32(c1.B)),
		A: uint8(float32(c2.A-c1.A)*t + float32(c1.A)),
	}
}

func fade(c color.RGBA, factor float32) color.RGBA {
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: uint8(float32(c.A) * factor),
	}
}

// DefaultFont is a package-level font face used for drawing text.
// Ensure that you set this variable (e.g., in your initialization code) to a valid font.Face.
var DefaultFont font.Face

type Button struct {
	Bounds            Rect
	Label             string
	BackgroundColor   color.RGBA
	HoverColor        color.RGBA
	PressedColor      color.RGBA
	BorderColor       color.RGBA
	TextColor         color.RGBA
	FontSize          int32
	FontFace          font.Face
	IsHovered         bool
	IsPressed         bool
	AnimationProgress float32
	Padding           float32
	CornerRadius      float32
	Enabled           bool
	Invisible         bool
	Uneditable        bool
	UseRoundedCorners bool

	prevMouseDown bool
	clicked       bool
}

func NewButton(x, y, width, height float32, label string) *Button {
	return &Button{
		Bounds:            NewRect(x, y, width, height),
		Label:             label,
		BackgroundColor:   color.RGBA{R: 211, G: 211, B: 211, A: 255}, // LightGray
		HoverColor:        color.RGBA{R: 200, G: 200, B: 200, A: 255},
		PressedColor:      color.RGBA{R: 169, G: 169, B: 169, A: 255}, // DarkGray
		BorderColor:       color.RGBA{R: 0, G: 0, B: 0, A: 255},         // Black
		TextColor:         color.RGBA{R: 0, G: 0, B: 0, A: 255},         // Black
		FontSize:          20,
		FontFace:          DefaultFont,
		IsHovered:         false,
		IsPressed:         false,
		AnimationProgress: 0.0,
		Padding:           5.0,
		CornerRadius:      5.0,
		Enabled:           true,
		Invisible:         false,
		Uneditable:        false,
		UseRoundedCorners: true,
		prevMouseDown:     false,
		clicked:           false,
	}
}

func (b *Button) SetColors(background, hover, pressed, border, text color.RGBA) {
	b.BackgroundColor = background
	b.HoverColor = hover
	b.PressedColor = pressed
	b.BorderColor = border
	b.TextColor = text
}

func (b *Button) SetFontSize(size int32) {
	b.FontSize = size
}

func (b *Button) SetCornerRadius(radius float32) {
	b.CornerRadius = radius
}

func (b *Button) SetPadding(padding float32) {
	b.Padding = padding
}

func (b *Button) SetEnabled(enabled bool) {
	b.Enabled = enabled
}

func (b *Button) SetInvisible(invisible bool) {
	b.Invisible = invisible
}

func (b *Button) IsInvisible() bool {
	return b.Invisible
}

func (b *Button) SetUneditable(uneditable bool) {
	b.Uneditable = uneditable
}

func (b *Button) IsUneditable() bool {
	return b.Uneditable
}

func (b *Button) SetRoundedCorners(rounded bool) {
	b.UseRoundedCorners = rounded
}

func (b *Button) IsRoundedCorners() bool {
	return b.UseRoundedCorners
}

// Update should be called every frame.
func (b *Button) Update() {
	if b.Uneditable {
		return
	}

	if !b.Enabled {
		b.IsHovered = false
		b.IsPressed = false
		return
	}

	mx, my := ebiten.CursorPosition()
	mouseX, mouseY := float32(mx), float32(my)
	b.IsHovered = pointInRect(mouseX, mouseY, b.Bounds)

	curMouseDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if b.IsHovered && curMouseDown {
		b.IsPressed = true
	} else {
		b.IsPressed = false
	}

	if b.IsHovered && !curMouseDown && b.prevMouseDown {
		b.clicked = true
	} else {
		b.clicked = false
	}
	b.prevMouseDown = curMouseDown

	// Update animation progress
	var targetProgress float32 = 0.0
	if b.IsPressed {
		targetProgress = 1.0
	} else if b.IsHovered {
		targetProgress = 0.5
	}

	animationSpeed := float32(8.0)
	// Assume a fixed frame time of 1/60 sec; in a real game, use actual delta time.
	frameTime := float32(1.0 / 60.0)
	if b.AnimationProgress < targetProgress {
		b.AnimationProgress += frameTime * animationSpeed
		if b.AnimationProgress > targetProgress {
			b.AnimationProgress = targetProgress
		}
	} else if b.AnimationProgress > targetProgress {
		b.AnimationProgress -= frameTime * animationSpeed
		if b.AnimationProgress < targetProgress {
			b.AnimationProgress = targetProgress
		}
	}
}

// Draw draws the button onto the given screen.
func (b *Button) Draw(screen *ebiten.Image) {
	if b.Invisible {
		return
	}

	var currentColor color.RGBA
	if !b.Enabled {
		currentColor = fade(b.BackgroundColor, 0.5)
	} else {
		if b.AnimationProgress <= 0.5 {
			t := b.AnimationProgress * 2.0
			currentColor = lerpColor(b.BackgroundColor, b.HoverColor, t)
		} else {
			t := (b.AnimationProgress - 0.5) * 2.0
			currentColor = lerpColor(b.HoverColor, b.PressedColor, t)
		}
	}

	borderThickness := float32(2.0)
	if b.IsPressed {
		borderThickness = 3.0
	}

	// Draw the button rectangle
	if b.UseRoundedCorners {
		drawRoundedRect(screen, b.Bounds.X, b.Bounds.Y, b.Bounds.W, b.Bounds.H, b.CornerRadius, currentColor)
		drawRoundedRectOutline(screen, b.Bounds.X, b.Bounds.Y, b.Bounds.W, b.Bounds.H, b.CornerRadius, borderThickness, b.BorderColor)
	} else {
		ebitenutil.DrawRect(screen, float64(b.Bounds.X), float64(b.Bounds.Y), float64(b.Bounds.W), float64(b.Bounds.H), currentColor)
		drawRectOutline(screen, b.Bounds, borderThickness, b.BorderColor)
	}

	// Draw label text centered
	if b.FontFace == nil {
		return // Cannot draw text without a valid font face.
	}
	bounds := text.BoundString(b.FontFace, b.Label)
	textWidth := float32(bounds.Dx())
	textX := b.Bounds.X + (b.Bounds.W-textWidth)/2.0
	textY := b.Bounds.Y + (b.Bounds.H-float32(b.FontSize))/2.0

	offsetX, offsetY := float32(0), float32(0)
	if b.IsPressed {
		offsetX, offsetY = 1.0, 1.0
	}

	text.Draw(screen, b.Label, b.FontFace, int(textX+offsetX), int(textY+offsetY)+int(b.FontSize), b.TextColor)
}

// IsClicked returns true if the button was clicked this frame.
func (b *Button) IsClicked() bool {
	return b.Enabled && b.IsHovered && b.clicked
}

// Helper function to draw a rounded rectangle filled with color.
func drawRoundedRect(screen *ebiten.Image, x, y, w, h, r float32, col color.RGBA) {
    path := &vector.Path{}
    // Clamp radius to half of width/height.
    r = float32(math.Min(float64(r), float64(w/2)))
    r = float32(math.Min(float64(r), float64(h/2)))

    // Start at top-left corner.
    path.MoveTo(x+r, y)
    // Top edge
    path.LineTo(x+w-r, y)
    // Top-right corner arc
    path.Arc(x+w-r, y+r, r, -90, 0, vector.Clockwise)
    // Right edge
    path.LineTo(x+w, y+h-r)
    // Bottom-right corner arc
    path.Arc(x+w-r, y+h-r, r, 0, 90, vector.Clockwise)
    // Bottom edge
    path.LineTo(x+r, y+h)
    // Bottom-left corner arc
    path.Arc(x+r, y+h-r, r, 90, 180, vector.Clockwise)
    // Left edge
    path.LineTo(x, y+r)
    // Top-left corner arc
    path.Arc(x+r, y+r, r, 180, 270, vector.Clockwise)
    path.Close()

	op := &ebiten.DrawTrianglesOptions{}
    // Draw filled path
    vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)
	screen.DrawTriangles(vertices, indices, nil, op)
}

// Helper function to draw rounded rectangle outline.
func drawRoundedRectOutline(screen *ebiten.Image, x, y, w, h, r, thickness float32, col color.RGBA) {
	// For simplicity, we re-use the filled rounded rect as an outline.
	// A full implementation would stroke the path.
	drawRoundedRect(screen, x, y, w, h, r, col)
}

// Helper function to draw rectangle outline for non-rounded rectangles.
func drawRectOutline(screen *ebiten.Image, r Rect, thickness float32, col color.RGBA) {
	// Top line
	ebitenutil.DrawLine(screen, float64(r.X), float64(r.Y), float64(r.X+r.W), float64(r.Y), col)
	// Bottom line
	ebitenutil.DrawLine(screen, float64(r.X), float64(r.Y+r.H), float64(r.X+r.W), float64(r.Y+r.H), col)
	// Left line
	ebitenutil.DrawLine(screen, float64(r.X), float64(r.Y), float64(r.X), float64(r.Y+r.H), col)
	// Right line
	ebitenutil.DrawLine(screen, float64(r.X+r.W), float64(r.Y), float64(r.X+r.W), float64(r.Y+r.H), col)
}

