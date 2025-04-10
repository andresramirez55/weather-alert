package main

import (
	"fmt"
	"log"

	"weather-alert/alerts"
	"weather-alert/config"
	"weather-alert/services"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	locations, err := config.LoadLocations("config/locations.json")
	if err != nil {
		log.Fatal(err)
	}

	emails, err := alerts.LoadEmails("config/emails.json")
	if err != nil {
		log.Fatal("❌ error getting emails:", err)
	}

	for _, loc := range locations {
		query := fmt.Sprintf("%s,%s", loc.City, loc.Country)

		weather, err := services.GetWeather(query)
		if err != nil {
			log.Printf("❌ Error getting weather %s: %v\n", query, err)
			continue
		}

		fmt.Printf("📍 %s, %s: %.1f°C - %s\n",
			weather.Name,
			weather.Sys.Country,
			weather.Main.Temp,
			weather.Weather[0].Description)

		alerts.ShouldTriggerAlert(weather, emails)
	}
}
