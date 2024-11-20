package symbol

import (
	"errors"
	"unicode"

	"github.com/manifoldco/promptui"
)

func GetSymbol() (string, error) {
	symbolPrompt := promptui.Prompt{
		Label:   "Input a Symbol or a Character (default: ▧)",
		Default: "▧",
		Validate: func(input string) error {
			if len([]rune(input)) != 1 {
				return errors.New("symbol must be a single symbol or character only")
			}

			r := []rune(input)[0]

			if unicode.IsSymbol(r) || unicode.Is(unicode.S, r) {
				return nil
			}

			return nil
		},
	}

	symbol, err := symbolPrompt.Run()

	return symbol, err
}
