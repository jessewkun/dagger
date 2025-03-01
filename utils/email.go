package utils

import (
	"net/smtp"
)

var (
	smtpHost = "smtp.example.com"
	smtpPort = "25"
	smtpUser = "user@example.com"
	smtpPass = "password"
)

// SendMail 发送邮件
func SendMail(to, subject, body string) error {
	auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPass,
		smtpHost,
	)

	return smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		smtpUser,
		[]string{to},
		[]byte("Subject: "+subject+"\n"+body),
	)
}
