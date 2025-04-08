package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func LoadLocations(filePath string) ([]Location, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("no se pudo leer el archivo: %v", err)
	}

	var locations []Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return nil, fmt.Errorf("error parseando locations.json: %v", err)
	}

	return locations, nil
}
