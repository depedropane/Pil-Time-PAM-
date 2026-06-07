package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EmailService struct {
	// Variabel SMTP tidak akan dipakai lagi, kita langsung pakai Resend API
}

func NewEmailService() *EmailService {
	return &EmailService{}
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

	resendAPIKey := "re_4LPFvZP2_9r1FpcT8DDAybXrDK2CNeSj2"
	
	payload := map[string]interface{}{
		// Wajib menggunakan onboarding@resend.dev jika domain belum diverifikasi di Resend
		"from":    "Pil Time <onboarding@resend.dev>",
		"to":      []string{recipientEmail},
		"subject": subject,
		"html":    body,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("gagal memproses payload email: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("gagal membuat request email: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+resendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("koneksi ke Resend gagal: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("resend error (status %d): %s", resp.StatusCode, string(respBody))
	}

	return nil
}
