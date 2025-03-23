package clip
// A thin wrapper around the clipboard package that makes more sense for people when entering names.
import (
	"fmt"

    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
    "os"
	"strconv"
	"github.com/atotto/clipboard"
)

// CopyClip copies the text to the system clipboard
func CopyClip(text string) error {
	err := clipboard.WriteAll(text)
	if err != nil {
		return fmt.Errorf("failed to copy text to clipboard: %w", err)
	}

	return nil
}

// PasteClip pastes the text from the system clipboard
func PasteClip() (string, error) {
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", fmt.Errorf("failed to paste text from clipboard: %w", err)
	}

	return text, nil
}

// LoadFontFace loads a TTF/OTF font file and returns a font.Face
func LoadFontFace(fontPath string, size float64) (font.Face, error) {
    // Read the font file
    fontData, err := os.ReadFile(fontPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read font file: %v", err)
    }

    // Parse the font
    parsedFont, err := opentype.Parse(fontData)
    if err != nil {
        return nil, fmt.Errorf("failed to parse font: %v", err)
    }

	// Get the dpi env var
	dpistr := os.Getenv("DPI")
	var dpi float64
	if dpistr == "" {
		dpi = 72
	} else {
		dpi, err = strconv.ParseFloat(dpistr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse DPI: %v", err)
		}
	}

    // Create the font face
    face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
        Size:    size,
        DPI:     dpi,
        Hinting: font.HintingFull,
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create font face: %v", err)
    }

    return face, nil
}