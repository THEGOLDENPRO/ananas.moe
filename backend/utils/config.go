package utils

import (
	"ananasmoe/types"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

func GetConfig[T any](key string) (T, error) {
	var none T
	configPath := os.Getenv("CONFIG_PATH")
	configFile := path.Join(configPath, "config.toml")

	var config types.TOML

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatalf("failed to read config TOML: %s", err)
		return none, fmt.Errorf("failed to read config TOML: %w", err)
	}

	switch key {
	case "cloud":
		var result T

		if val, ok := any(config.Clouds).(T); ok {
			result = val
		} else {
			return none, fmt.Errorf("could not cast Clouds to type %T", result)
		}
		return result, nil
	case "project":
		var result T

		if val, ok := any(config.Projects).(T); ok {
			result = val
		} else {
			return none, fmt.Errorf("could not cast Projects to type %T", result)
		}
		return result, nil
	case "redirect":
		var result T

		if val, ok := any(config.Redirects).(T); ok {
			result = val
		} else {
			return none, fmt.Errorf("could not cast Redirects to type %T", result)
		}
		return result, nil
	case "file":
		var result T

		if val, ok := any(config.Files).(T); ok {
			result = val
		} else {
			return none, fmt.Errorf("could not cast Redirects to type %T", result)
		}
		return result, nil
	default:
		return none, fmt.Errorf("unknown key: %s", key)
	}
}
