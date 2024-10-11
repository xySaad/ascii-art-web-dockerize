package ascii

import (
	"ascii-art-web/utils"
	"errors"
	"os"
	"strings"
)

type Args struct {
	Text       string
	BannerName string
}

type Ascii struct {
	Value   string
	Message string
}

func Generate(args Args) (ascii Ascii, errs error) {
	plainTxt, err := os.ReadFile("./assets/banners/" + args.BannerName + ".txt")
	if err != nil {
		return Ascii{Value: "", Message: "Error Opening Font File"}, err
	}
	txt := strings.ReplaceAll(string(plainTxt), "\r\n", "\n")
	fileLines := strings.Split(txt, "\n")
	if len(fileLines) != 856 {
		messsage := "banner file " + args.BannerName + " has been modified and is invalid"
		return Ascii{Value: "", Message: "Corrupted Font File"}, errors.New(messsage)
	}
	// Fetch the input from command-line arguments and clean it
	userInput, hasInvalidCharacters := utils.CleanString(args.Text)

	if hasInvalidCharacters {
		ascii.Message = "Removed Invalid Characters"
	}

	if len(userInput) == 0 {
		ascii.Value = ""
		return ascii, nil
	} else if userInput == "\\n" {
		ascii.Value = "\\n"
		return ascii, nil
	}
	// Split the input based on the newline delimiter
	inputWords := strings.Split(userInput, "\n")

	if utils.IsEmpty(inputWords) {
		inputWords = inputWords[:len(inputWords)-1]
	}
	// Iterate through each word and process it
	for _, word := range inputWords {
		if word == "" || word == "\n" {
			ascii.Value += "\n"
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
			ascii.Value += renderedLine + "\n"
		}
	}
	if !hasInvalidCharacters {
		ascii.Message = "Success"
	}

	return ascii, nil
}
