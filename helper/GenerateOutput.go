package helper

import (
	"fmt"
	"os"
	"strings"
)

func OutputFlag(s string) string {
	var outputFileName string
	if strings.HasPrefix(s, "--output=") {
		outputFileName = strings.TrimPrefix(s, "--output=")
		if !CheckOutput(outputFileName) {
			os.Exit(0)
		}

	}
	return outputFileName
}

func CheckOutput(s string) bool {
	sl := strings.ToLower(s)

	if !strings.HasSuffix(sl, ".txt") || sl == "shadow.txt" || sl == "standard.txt" || sl == "thinkertoy.txt" || strings.Contains(sl, "/") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return false
	}
	return true
}
