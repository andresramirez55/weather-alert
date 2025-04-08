package alerts

import (
	"fmt"
	"strings"

	"weather-alert/models"
)

func ShouldTriggerAlert(weather *models.WeatherResponse) (bool, string) {
	condition := strings.ToLower(weather.Weather[0].Main)
	description := strings.ToLower(weather.Weather[0].Description)
	windSpeed := weather.Wind.Speed
	temp := weather.Main.Temp

	// Lógica simple: alerta si hay tormenta, granizo, viento fuerte o calor extremo
	if condition == "thunderstorm" || strings.Contains(description, "hail") {
		return true, fmt.Sprintf("⚠️ Alerta de tormenta en %s: %s (%.1f°C)", weather.Name, description, temp)
	}

	if windSpeed > 15 {
		return true, fmt.Sprintf("💨 Vientos fuertes en %s: %.1f m/s", weather.Name, windSpeed)
	}

	if temp > 38 {
		return true, fmt.Sprintf("🔥 Ola de calor en %s: %.1f°C", weather.Name, temp)
	}

	if temp < -5 {
		return true, fmt.Sprintf("❄️ Frío extremo en %s: %.1f°C", weather.Name, temp)
	}

	return false, ""
}
