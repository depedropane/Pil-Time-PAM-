package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
)

type EmailService struct {
	smtpHost       string
	smtpPort       string
	senderEmail    string
	senderPassword string
}

func NewEmailService() *EmailService {
	return &EmailService{
		smtpHost:       os.Getenv("SMTP_HOST"),
		smtpPort:       os.Getenv("SMTP_PORT"),
		senderEmail:    os.Getenv("SENDER_EMAIL"),
		senderPassword: os.Getenv("SENDER_PASSWORD"),
	}
}

func (e *EmailService) SendResetCode(recipientEmail, resetCode string) error {
	// Email content
	subject := "Kode Reset Password Pil Time"
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<style>
		body { font-family: Arial, sans-serif; }
		.container { max-width: 600px; margin: 0 auto; padding: 20px; }
		.header { background-color: #15BE77; color: white; padding: 20px; text-align: center; border-radius: 8px; }
		.content { padding: 20px; background-color: #f5f5f5; margin: 10px 0; border-radius: 8px; }
		.code { background-color: white; padding: 20px; text-align: center; font-size: 32px; font-weight: bold; color: #15BE77; border-radius: 8px; margin: 20px 0; }
		.footer { text-align: center; color: #757575; font-size: 12px; margin-top: 20px; }
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1>Reset Password</h1>
		</div>
		<div class="content">
			<p>Halo,</p>
			<p>Anda telah meminta untuk mereset password akun Pil Time Anda.</p>
			<p>Gunakan kode berikut untuk mereset password Anda:</p>
			<div class="code">%s</div>
			<p><strong>Kode ini berlaku selama 15 menit.</strong></p>
			<p>Jika Anda tidak meminta ini, abaikan email ini.</p>
		</div>
		<div class="footer">
			<p>Pil Time - Layanan Kesehatan Terpercaya</p>
			<p>&copy; 2026 All rights reserved.</p>
		</div>
	</div>
</body>
</html>
	`, resetCode)

	// SMTP configuration
	auth := smtp.PlainAuth("", e.senderEmail, e.senderPassword, e.smtpHost)
	addr := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s", recipientEmail, subject, body))

	var err error
	if e.smtpPort == "465" {
		// Port 465 requires Implicit TLS
		tlsconfig := &tls.Config{
			InsecureSkipVerify: false,
			ServerName:         e.smtpHost,
		}

		conn, errDial := tls.Dial("tcp", addr, tlsconfig)
		if errDial != nil {
			return errDial
		}
		defer conn.Close()

		client, errClient := smtp.NewClient(conn, e.smtpHost)
		if errClient != nil {
			return errClient
		}
		defer client.Quit()

		if err = client.Auth(auth); err != nil {
			return err
		}
		if err = client.Mail(e.senderEmail); err != nil {
			return err
		}
		if err = client.Rcpt(recipientEmail); err != nil {
			return err
		}

		w, errData := client.Data()
		if errData != nil {
			return errData
		}
		_, err = w.Write(msg)
		if err != nil {
			return err
		}
		err = w.Close()
	} else {
		// Default to STARTTLS for port 587
		err = smtp.SendMail(
			addr,
			auth,
			e.senderEmail,
			[]string{recipientEmail},
			msg,
		)
	}

	return err
}
