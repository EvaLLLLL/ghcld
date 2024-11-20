package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/EvaLLLLL/ghcld/symbol"
	"github.com/EvaLLLLL/ghcld/types"
	"github.com/manifoldco/promptui"
)

var CONFIG_PATH = filepath.Join(os.Getenv("HOME"), ".config", "ghcld")

func CheckConfigValue() (*types.Config, error) {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		return &types.Config{}, fmt.Errorf("configuration file %s does not exist", CONFIG_PATH)
	}

	data, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		return &types.Config{}, fmt.Errorf("error reading configuration file: %s", err)
	}

	config := types.Config{}
	if err := json.Unmarshal(data, &config); err != nil {
		return &types.Config{}, fmt.Errorf("configuration file format error: %s", err)
	}

	if value := config; value.USER_NAME != "" && value.TOKEN != "" {
		return &value, nil
	} else {
		fmt.Printf("Configuration item '%s' does not exist.\n", CONFIG_PATH)
		return &types.Config{}, fmt.Errorf("configuration item '%s' does not exist", CONFIG_PATH)
	}
}

func InitConfig() (*types.Config, error) {
	if _, err := os.Stat(CONFIG_PATH); err == nil {
		prompt := promptui.Prompt{
			Label:     "Configuration file already exists, do you want to delete the old file?",
			IsConfirm: true,
		}

		deleteOld, _ := prompt.Run()

		if deleteOld == "y" {
			err := os.Remove(CONFIG_PATH)
			if err != nil {
				return &types.Config{}, fmt.Errorf("error deleting old configuration file: %s", err)
			}
			fmt.Println("Old configuration file has been deleted.")
		} else {
			config, err := CheckConfigValue()
			return config, err
		}
	}

	usernamePrompt := promptui.Prompt{
		Label: "Please enter your github username",
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("username cannot be empty")
			}
			return nil
		},
	}

	tokenPrompt := promptui.Prompt{
		Label: "Please enter your github Token ( Get token here: https://github.com/settings/tokens/new )",
		Mask:  '*',
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("token cannot be empty")
			}
			return nil
		},
	}

	username, usernamePromptErr := usernamePrompt.Run()
	token, tokenPromptErr := tokenPrompt.Run()
	symbol, symbolPromptErr := symbol.GetSymbol()

	if usernamePromptErr != nil || tokenPromptErr != nil || symbolPromptErr != nil {
		return &types.Config{}, fmt.Errorf("error creating config")
	}

	config := types.Config{
		USER_NAME: username,
		TOKEN:     token,
		SYMBOL:    symbol,
	}

	configFile, err := os.Create(CONFIG_PATH)
	if err != nil {
		return &types.Config{}, fmt.Errorf("error creating configuration file: %s", err)
	}
	defer configFile.Close()

	encoder := json.NewEncoder(configFile)
	err = encoder.Encode(config)
	if err != nil {
		return &types.Config{}, fmt.Errorf("error writing to configuration file: %s", err)
	}

	return &config, err
}
