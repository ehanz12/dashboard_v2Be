package services

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendVerificationEmail sends a verification code to the user's email
func SendVerificationEmail(email, verificationCode string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	appName := os.Getenv("APP_NAME")

	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPassword == "" {
		return fmt.Errorf("SMTP configuration not set")
	}

	// Email body
	subject := "Email Verification - " + appName
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background-color: #4CAF50; color: white; padding: 20px; text-align: center; }
        .content { padding: 20px; background-color: #f9f9f9; }
        .code { font-size: 24px; font-weight: bold; color: #4CAF50; text-align: center; padding: 20px; }
        .footer { text-align: center; padding: 10px; color: #999; font-size: 12px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Email Verification</h1>
        </div>
        <div class="content">
            <p>Hello,</p>
            <p>Thank you for registering with %s! To complete your registration, please verify your email address using the code below:</p>
            <div class="code">%s</div>
            <p>This code will expire in 24 hours.</p>
            <p>If you didn't create this account, please ignore this email.</p>
        </div>
        <div class="footer">
            <p>&copy; 2024 %s. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
`, appName, verificationCode, appName)

	// Set email headers
	message := fmt.Sprintf("From: %s\r\n", smtpUser)
	message += fmt.Sprintf("To: %s\r\n", email)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "MIME-Version: 1.0\r\n"
	message += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	message += "\r\n"
	message += body

	// Send email via SMTP
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		smtpUser,
		[]string{email},
		[]byte(message),
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
