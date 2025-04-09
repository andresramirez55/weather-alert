package main

import (
	"fmt"
	"log"

	"weather-alert/alerts"
	"weather-alert/config"
	"weather-alert/notifier"
	"weather-alert/services"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	locations, err := config.LoadLocations("config/locations.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, loc := range locations {
		query := fmt.Sprintf("%s,%s", loc.City, loc.Country)

		weather, err := services.GetWeather(query)
		if err != nil {
			log.Printf("❌ Error al obtener clima para %s: %v\n", query, err)
			continue
		}

		//print for test
		fmt.Printf("📍 %s, %s: %.1f°C - %s\n",
			weather.Name,
			weather.Sys.Country,
			weather.Main.Temp,
			weather.Weather[0].Description)

		shouldAlert, message := alerts.ShouldTriggerAlert(weather)
		if shouldAlert {
			err := notifier.SendWhatsAppAlert(message)
			if err != nil {
				log.Printf("❌ Error al enviar alerta: %v\n", err)
			}
		} else {
			fmt.Println("✅ No se requiere alerta.")
		}
	}
}
