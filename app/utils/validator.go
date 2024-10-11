package utils

import (
	"os"
)

// Cleanstring removes all unprintable characters from a string
func CleanString(s string) (string, bool) {
	var cleanString []rune
	var hasInvalidCharacters bool
	for _, r := range s {
		if (r >= 32 && r <= 126) || r == 10 { // Check if the character is printable
			cleanString = append(cleanString, r)
			continue
		} else {
			hasInvalidCharacters = true
		}

	}

	return string(cleanString), hasInvalidCharacters
}

func IsEmpty(text []string) bool {
	for _, l := range text {
		if l != "" {
			return false
		}
	}
	return true
}

func IsValidBanner(BannerName string) bool {
	_, err := os.Stat("./assets/banners/" + BannerName + ".txt")
	return !os.IsNotExist(err)
}

func GetBanners() (string, error) {
	// Specify the directory you want to list files from
	dir := "./assets/banners"
	banners := ""
	// Read the directory entries
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", nil
	}

	// Loop through the directory entries and print their names
	for _, entry := range entries {
		banners += entry.Name()[:len(entry.Name())-4] + ","
	}
	return banners, nil
}
