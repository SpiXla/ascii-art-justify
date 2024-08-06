package main

import (
	"fmt"
	"helper/helper"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}

	var outputFileName, inputText, bannerName, alignOption string

	switch len(os.Args) {
	case 2:
		inputText = os.Args[1]
		bannerName = "standard"
	case 3:
		if strings.HasPrefix(os.Args[1], "--output=") {
			outputFileName = helper.OutputFlag(os.Args[1])
			inputText = os.Args[2]
			bannerName = "standard"
		} else if strings.HasPrefix(os.Args[1], "--align=") {
			alignOption = helper.AlignFlag(os.Args[1])
			inputText = os.Args[2]
			bannerName = "standard"
		} else {
			inputText = os.Args[1]
			bannerName = os.Args[2]
		}
	case 4:
		flag := os.Args[1]
		if strings.HasPrefix(flag, "--output=") {
			outputFileName = helper.OutputFlag(flag)
		}
		if strings.HasPrefix(flag, "--align=") {
			alignOption = helper.AlignFlag(flag)
		}
		inputText = os.Args[2]
		bannerName = os.Args[3]
	}

	input := strings.Split(inputText, "\\n")
	if !helper.CheckNline(input) && len(input) > 0 {
		input = input[1:]
	}
	inputText = strings.Join(input, "\\n")

	if !helper.IsValidString(inputText) {
		fmt.Println("Error: Input string contains invalid characters")
		return
	}

	bannerContent, err := os.ReadFile(bannerName + ".txt")
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	banner := helper.ParseBanner(string(bannerContent))
	result := helper.GenerateArt(inputText, banner)
	if result == "" {
		fmt.Println("Error generating art")
		return
	}

	if alignOption != "" {
		result = helper.GenerateAlign(result, alignOption)
	}

	if outputFileName != "" {
		err := os.WriteFile(outputFileName, []byte(result), 0644)
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		fmt.Println("Result has been written successfully")
	} else {
		fmt.Print(result)
	}
}
