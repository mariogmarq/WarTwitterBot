package utils

import (
	"errors"
	"strings"
)

// Parses the image name to the player name
func ParseImageName(name string) (string, error) {
	name = strings.Trim(name, " ")
	arr := strings.Split(name, ".")
	if len(arr) != 2 {
		return "", errors.New("no proper file name")
	}

	if arr[1] != "png" {
		return "", errors.New("no png image")
	}

	name = strings.ReplaceAll(arr[0], "_", " ")
	name = strings.Title(strings.ToLower(name))

	return name, nil
}
