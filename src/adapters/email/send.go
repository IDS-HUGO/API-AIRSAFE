package email

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendEmail envía un correo electrónico con credenciales.
func SendEmail(to, subject, body string) error {
	// Configurar credenciales SMTP
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Crear mensaje
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)

	// Enviar correo
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("error al enviar correo: %v", err)
	}

	fmt.Println("Correo enviado a", to)
	return nil
}
