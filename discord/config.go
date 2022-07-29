package discord

import (
	"encoding/json"
	"fmt"
	"os"
)

type configType map[string]interface{}

var config = make(configType)

func LoadConfig() (configType, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config.json: %v", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding config.json: %v", err)
	}

	return config, nil
}
