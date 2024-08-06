package helper
import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
	//Secure Shell (SSH)
)

// GenerateAlign generates aligned text based on the specified alignment.
func GenerateAlign(text string, align string) string {
	var result string
	_, cols, err := GetTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		os.Exit(0)
	}

	lines := strings.Split(text, "\n")
	switch align {
	case "center":
		for _, line := range lines {
			result += CenterText(line, cols) + "\n"
		}
	case "right":
		for _, line := range lines {
			result += RightAlignText(line, cols) + "\n"
		}
	case "justify":
		for _, line := range lines {
			result += JustifyText(line, cols) + "\n"
		}
	case "left":
	// 	fallthrough
	// default:
		result = text
	}

	return result
}

// Align extracts the alignment option from the flag.
func AlignFlag(flag string) string {
	if strings.HasPrefix(flag, "--align=") {
		align := strings.TrimPrefix(flag, "--align=")
		if !CheckAlign(align) {
			os.Exit(0)
		}
		return align
	}
	return ""
}

// CheckAlign validates the alignment option.
func CheckAlign(align string) bool {
	validAlignments := []string{"center", "left", "right", "justify"}
	for _, valid := range validAlignments {
		if align == valid {
			return true
		}
	}
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	return false
}

// GetTerminalSize retrieves the terminal's width and height.
func GetTerminalSize() (int, int, error) {
	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, err
	}
	return height, width, nil
}

// CenterText centers the given text based on the terminal width.
func CenterText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	padding := (width - len(text)) / 2
	return fmt.Sprintf("%*s%s%*s", padding, "", text, padding, "")
}

// RightAlignText right-aligns the given text based on the terminal width.
func RightAlignText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	padding := width - len(text)
	return fmt.Sprintf("%*s%s", padding, "", text)
}

func JustifyText(text string, width int) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	if len(words) == 1 {
		// If there's only one word, center it
		return CenterText(text, width)
	}

	totalWordLen := 0
	for _, word := range words {
		totalWordLen += len(word)
	}
	totalSpaces := width - totalWordLen
	spacesPerGap := totalSpaces / (len(words) - 1)
	extraSpaces := totalSpaces % (len(words) - 1)

	var result strings.Builder
	for i, word := range words {
		result.WriteString(word)
		if i < len(words)-1 {
			spaces := spacesPerGap
			if i < extraSpaces {
				spaces++
			}
			result.WriteString(strings.Repeat(" ", spaces))
		}
	}

	return result.String()
}