package notifier

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendWhatsAppAlert(message string) error {
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_WHATSAPP_FROM")
	to := os.Getenv("TWILIO_WHATSAPP_TO")

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", message)

	twilioURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountSID)
	req, _ := http.NewRequest("POST", twilioURL, strings.NewReader(msgData.Encode()))
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("❌ Error al enviar alerta por WhatsApp: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Println("📲 Alerta de WhatsApp enviada con éxito")
		return nil
	}

	log.Printf("❌ Error al enviar alerta por WhatsApp: %d", resp.StatusCode)
	return nil
}
