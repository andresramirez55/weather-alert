package notifier

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmailAlert(to string, message string) error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Alerta Meteorológica ⚠️"
	body := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(body))
	if err != nil {
		return fmt.Errorf("falló el envío de email a %s: %w", to, err)
	}

	fmt.Printf("📧 Email enviado con éxito a %s\n", to)
	return nil
}
