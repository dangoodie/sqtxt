package util

import (
	"fmt"
	"sync"
)

// Config holds configuration settings
type Config struct {
	FontSize  int
	RowHeight int
	ColWidth  int
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the configuration instance
func GetConfig() *Config {
	once.Do(func() {
		config, err := loadConfig("")
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			instance = defaultConfig()
		} else {
			instance = config
		}
	})
	return instance
}

// LoadConfig loads the configuration from a file
func loadConfig(filepath string) (*Config, error) {
	if filepath == "" {
		fmt.Println("No configuration file specified")
		fmt.Println("Using default configuration")
		return defaultConfig(), nil
	} else {
		// Load the configuration from the file
		fmt.Println("Loading configuration from", filepath)
		return parseConfig(filepath)
	}
}

func defaultConfig() *Config {
	return &Config{
		FontSize:  12,
		RowHeight: 12,
		ColWidth:  6,
	}
}

func parseConfig(filepath string) (*Config, error) {
	// Parse the configuration file
	// and return the configuration
	// TODO: Implement this function

	return defaultConfig(), nil
}
