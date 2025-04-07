//go:build js && wasm
// +build js,wasm

package clip

// CopyClip does nothing on WASM.
func CopyClip(text string) error {
	return nil
}

// PasteClip returns an empty string on WASM.
func PasteClip() (string, error) {
	return "", nil
}
