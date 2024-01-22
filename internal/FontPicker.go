package internal

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
)

const (
	StandardHash   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	ShadowHash     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	ThinkertoyHash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

func FontPicker(font string) (string, error) {
	errf := errors.New("the font does not exist, or has been corrupted")

	file, err := os.Open("./fonts/" + font + ".txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := os.ReadFile("./fonts/" + font + ".txt")
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(data)
	generatedHash := fmt.Sprintf("%x", hasher.Sum(nil))

	if generatedHash != StandardHash && generatedHash != ShadowHash && generatedHash != ThinkertoyHash {
		return "", errf
	}

	return string(data), nil
}
