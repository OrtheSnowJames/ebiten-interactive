//go:build !js || !wasm
// +build !js !wasm

package clip

import (
	"fmt"

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
