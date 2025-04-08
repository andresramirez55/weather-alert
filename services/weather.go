package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather-alert/models"
)

func GetWeather(location string) (*models.WeatherResponse, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("no se encontró la clave OPENWEATHER_API_KEY en las variables de entorno")
	}

	encodedLocation := url.QueryEscape(location)

	apiURL := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		encodedLocation,
		apiKey,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error haciendo la petición HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo la respuesta: %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("❌ Ciudad '%s' no encontrada", location)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("respuesta inválida del servidor: %s", resp.Status)
	}

	var weatherResp models.WeatherResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		return nil, fmt.Errorf("error parseando JSON: %v", err)
	}

	// Validación mínima de campos esperados
	if weatherResp.Name == "" || len(weatherResp.Weather) == 0 {
		return nil, fmt.Errorf("respuesta inesperada del servidor: datos incompletos")
	}

	return &weatherResp, nil
}
