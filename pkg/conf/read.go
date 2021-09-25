package conf

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gleich/lumber/v2"
	"github.com/pelletier/go-toml/v2"
)

// Read from the configuration file and parse it
func Read() (Conf, error) {
	lumber.Info("Loading configuration")
	// Configuration location
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Conf{}, err
	}
	location := filepath.Join(homeDir, ".config", "verpi", "conf.toml")

	// Reading the binary from the file
	b, err := os.ReadFile(location)
	if err != nil {
		return Conf{}, err
	}

	// Parsing toml
	var data Conf
	err = toml.Unmarshal(b, &data)
	if err != nil {
		return Conf{}, err
	}

	// Validate config
	if data.Token == "" {
		return Conf{}, errors.New("token value in configuration file is required")
	}
	if data.Brightness == nil {
		*data.Brightness = 0.1
	}

	lumber.Success("Loaded configuration")

	return data, nil
}
