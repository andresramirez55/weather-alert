package alerts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"weather-alert/models"
	"weather-alert/notifier"
)

func ShouldTriggerAlert(weather *models.WeatherResponse, emails []string) {
	condition := strings.ToLower(weather.Weather[0].Main)
	description := strings.ToLower(weather.Weather[0].Description)
	windSpeed := weather.Wind.Speed
	temp := weather.Main.Temp

	var alertMessage string
	var shouldSend bool

	if condition == "thunderstorm" || strings.Contains(description, "hail") {
		alertMessage = fmt.Sprintf("⚠️ Alerta de tormenta en %s: %s (%.1f°C)", weather.Name, description, temp)
		shouldSend = true
	} else if strings.Contains(condition, "rain") || strings.Contains(description, "rain") {
		alertMessage = fmt.Sprintf("🌧️ Lluvia pronosticada en %s: %s (%.1f°C)", weather.Name, description, temp)
		shouldSend = true
	} else if windSpeed > 15 {
		alertMessage = fmt.Sprintf("💨 Vientos fuertes en %s: %.1f m/s", weather.Name, windSpeed)
		shouldSend = true
	} else if temp > 38 {
		alertMessage = fmt.Sprintf("🔥 Ola de calor en %s: %.1f°C", weather.Name, temp)
		shouldSend = true
	} else if temp < -5 {
		alertMessage = fmt.Sprintf("❄️ Frío extremo en %s: %.1f°C", weather.Name, temp)
		shouldSend = true
	}

	if shouldSend {
		for _, email := range emails {
			err := notifier.SendEmailAlert(email, alertMessage)
			if err != nil {
				log.Printf("❌ Error enviando alerta por email a %s: %v\n", email, err)
			}
		}
	}
}

func LoadEmails(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var emails []string
	err = json.Unmarshal(bytes, &emails)
	return emails, err
}
