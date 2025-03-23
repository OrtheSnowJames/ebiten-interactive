# Ebiten Interactive Library

The **Ebiten Interactive Library** is a collection of reusable interactive components for building graphical user interfaces (GUIs) in games or applications using the [Ebiten](https://ebiten.org/) game library. This library provides components such as buttons, text fields, and clipboard utilities to simplify the creation of interactive elements.

---

## Features

- **Buttons**: Fully customizable buttons with hover, press, and animation effects.
- **Text Fields**: Editable text fields with support for placeholders, clipboard operations, and cursor navigation.
- **Clipboard Utilities**: Easy-to-use clipboard integration for copying and pasting text.
- **Font Loading**: Load custom fonts for use in your interactive components.
- **Batch Operations**: Update and draw multiple interactive objects with a single function call.

---

## Installation

To use this library, ensure you have Go installed and run:

```bash
go get github.com/OrtheSnowJames/ebiten-interactive/interact
```

---

## Usage

### 1. Initialize Components

Import the library and initialize the components you need:

```go
package main

import (
    "github.com/OrtheSnowJames/ebiten-interactive/interact"
    "github.com/hajimehoshi/ebiten/v2"
)

func main() {
    button := interact.NewButton(50, 50, 200, 50, "Click Me")
    textField := interact.NewTextField(50, 120, 300, 40, 20)

    // Use these components in your game loop.
}
```

### 2. Drawing and Updating

Use `interact.DrawAll` and `interact.UpdateAll` to manage multiple components:

```go
func update(screen *ebiten.Image) error {
    interact.UpdateAll(button, textField)
    interact.DrawAll(screen, button, textField)
    return nil
}
```

### 3. Clipboard Integration

Copy and paste text using the clipboard utilities:

```go
err := interact.CopyClip("Hello, Clipboard!")
if err != nil {
    panic(err)
}

text, err := interact.PasteClip()
if err != nil {
    panic(err)
}
```

---

## Components

### Button

A customizable button with hover and press effects. Supports multiple styles including rounded corners and pointy (arrow-like) edges.

```go
button := interact.NewButton(50, 50, 200, 50, "Click Me")
button.SetColors(background, hover, pressed, border, text)
button.SetFontSize(24)

// Set button style
button.SetRoundedCorners(true)  // Enable rounded corners (default)
// OR
button.SetPointyStyle(true)     // Enable pointy/arrow style
button.SetPointyAmount(15)      // Adjust how pointy the arrows are
```

### Text Field

An editable text field with cursor navigation and clipboard support.

```go
textField := interact.NewTextField(50, 120, 300, 40, 20)
textField.SetPlaceholder("Enter your name...")
textField.SetFont(customFont)
```

### Clipboard Utilities

Easily copy and paste text to/from the system clipboard.

```go
interact.CopyClip("Sample Text")
text, _ := interact.PasteClip()
```

---

## Font Loading

Load custom fonts using the `LoadFontFace` function:

```go
fontFace, err := interact.LoadFontFace("path/to/font.ttf", 16)
if err != nil {
    panic(err)
}
interact.DefaultFont = fontFace
```

---

## Complete Example

Here's a comprehensive example showing how to use multiple components together with the simplified API:

```go
package main

import (
    "log"
    "image/color"
    
    "github.com/OrtheSnowJames/ebiten-interactive/interact"
    // for more colorschemes import colorscheme
    "github.com/OrtheSnowJames/ebiten-interactive/interact/colorscheme"
    "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
    normalButton   *interact.Button
    pointyButton   *interact.Button
    roundedButton  *interact.Button
    inputField     *interact.TextField
    submitButton   *interact.Button
    components     []interact.InteractiveObject
}

func NewGame() *Game {
    // Load font
    fontFace, err := interact.LoadFontFace("assets/fonts/RobotoMono-Regular.ttf", 20)
    if err != nil {
        log.Fatal(err)
    }
    interact.DefaultFont = fontFace

    g := &Game{}

    // Normal button with red scheme
    g.normalButton = interact.NewNormalButton(50, 50, 200, 40, "Normal Button")
    g.normalButton.SetColorScheme(colorscheme.RedScheme())

    // Pointy button with blue scheme
    g.pointyButton = interact.NewPointyButton(50, 100, 200, 40, "Pointy Button", 15)
    g.pointyButton.SetColorScheme(colorscheme.BlueScheme())

    // Rounded button with green scheme
    g.roundedButton = interact.NewRoundedButton(50, 150, 200, 40, "Rounded Button", 10)
    g.roundedButton.SetColorScheme(colorscheme.GreenScheme())

    // Text field with placeholder
    g.inputField = interact.NewTextFieldWithPlaceholder(50, 200, 300, 40, 50, "Type something here...")
    g.inputField.SetColors(
        color.RGBA{255, 255, 255, 255}, // background
        color.RGBA{100, 100, 100, 255}, // border
        color.RGBA{0, 0, 0, 255},       // text
    )

    // Submit button with violet scheme
    g.submitButton = interact.NewRoundedButton(360, 200, 100, 40, "Submit", 20)
    g.submitButton.SetColorScheme(colorscheme.VioletScheme())

    // Store all components for batch updates
    g.components = []interact.InteractiveObject{
        g.normalButton,
        g.pointyButton,
        g.roundedButton,
        g.inputField,
        g.submitButton,
    }

    return g
}

func (g *Game) Update() error {
    interact.UpdateAll(g.components...)

    if g.submitButton.IsClicked() {
        log.Printf("Submitted text: %s", g.inputField.GetText())
    }

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    interact.DrawAll(screen, g.components...)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 800, 600
}

func main() {
    ebiten.SetWindowSize(800, 600)
    ebiten.SetWindowTitle("Interactive Components Example")

    if err := ebiten.RunGame(NewGame()); err != nil {
        log.Fatal(err)
    }
}
```

This optimized example demonstrates:
- Simplified imports using just the main `interact` package
- Factory functions for different button styles
- Color scheme support for consistent styling
- Batch updates and drawing with `UpdateAll` and `DrawAll`
- Component storage in a slice for easy management
- Simplified text field creation with placeholder
- Event handling with button clicks

Run this example to see all the interactive components working together in a single window.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the library.

---

## Acknowledgments

- Built with [Ebiten](https://ebitengine.org/).
- Clipboard integration powered by [atotto/clipboard](https://github.com/atotto/clipboard).
- Font rendering provided by [golang.org/x/image/font](https://pkg.go.dev/golang.org/x/image/font).

---