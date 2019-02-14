package alerts

import (
	"strconv"

	"github.com/chennqqi/gosm/models"
	"gopkg.in/gomail.v2"
)

func sendSMTPAlert(service *models.Service) error {
	cfg := &models.CurrentConfig.Smtp
	var message string
	message += service.Name + " is now " + service.Status + "\r\n"
	message += "Protocol: " + service.Protocol + "\r\n"
	message += "Host: " + service.Host + "\r\n"
	if service.Port.Value != nil {
		message += "Port: " + strconv.FormatInt(service.Port.Int64, 10) + "\r\n"
	}

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EmailAddress)
	m.SetHeader("To", models.CurrentConfig.EmailRecipients...)
	m.SetHeader("Subject", "Subject: [gosm] "+service.Name+" is "+service.Status+"\r\n")
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)

	// Send the email to Bob, Cora and Dan.
	return d.DialAndSend(m)
}
