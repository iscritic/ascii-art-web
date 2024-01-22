package main

import "ascii/internal"

func GenerateAscii(text, banner string) (string, error) {
	data, err := internal.FontPicker(banner)
	if err != nil {
		return "", err
	}

	output := internal.GetAscii(text, data)

	return output, nil
}
