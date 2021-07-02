package utils

import (
	"errors"
	"os"
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

func ReadNamesFromImagesFolder() []string {
	//read all images from directory
	entries, err := os.ReadDir(os.Getenv("IMAGES_DIR"))
	Must(err)

	var filenames []string
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			filenames = append(filenames, entry.Name())
		}
	}

	return filenames
}

func Must(e error) {
	if e != nil {
		panic(e.Error())
	}
}
