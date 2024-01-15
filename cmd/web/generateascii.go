package main

import (
	. "ascii/internal"
)

func GenerateAscii(text, banner string) (string, error) {

	data, err := FontPicker(banner)
	if err != nil {
		return "", err
	}

	output := GetAscii(text, data)

	return output, nil

}
