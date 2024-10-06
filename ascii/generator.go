package ascii

import (
	"errors"
	"os"
	"strings"

	"ascii-art-web/ascii/utils"
)

type Args struct {
	Text       string
	BannerName string
}

func Generate(args Args) (ascii string, errs error) {
	plainTxt, err := os.ReadFile("./assets/banners/" + args.BannerName + ".txt")
	if err != nil {
		return "", err
	}
	txt := strings.ReplaceAll(string(plainTxt), "\r\n", "\n")
	fileLines := strings.Split(txt, "\n")
	if len(fileLines) != 856 {
		messsage := "banner file " + args.BannerName + " has been modified and is invalid"
		return "", errors.New(messsage)
	}
	// Fetch the input from command-line arguments and clean it
	userInput := utils.CleanString(args.Text)

	if len(userInput) == 0 {
		return "", nil
	} else if userInput == "\\n" {
		return "\n", nil
	}
	// Split the input based on the newline delimiter
	inputWords := strings.Split(userInput, "\n")

	if utils.IsEmpty(inputWords) {
		inputWords = inputWords[:len(inputWords)-1]
	}
	// Iterate through each word and process it
	for _, word := range inputWords {
		if word == "" || word == "\n" {
			ascii += "\n"
			continue
		}

		// Render each word line by line
		for lineOffset := 0; lineOffset < 8; lineOffset++ {
			var renderedLine string

			for _, runeChar := range word {
				startIndex := int((runeChar-32)*9) + 1
				renderedLine += fileLines[startIndex+lineOffset]
			}

			// Output the constructed line
			ascii += renderedLine + "\n"
		}
	}
	return ascii, nil
}
