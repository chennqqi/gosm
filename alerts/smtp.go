package alerts

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/chennqqi/gosm/models"
)

func sendSMTPAlert(service *models.Service) {
	cfg := &models.CurrentConfig.Smtp
	auth := smtp.PlainAuth("",
		cfg.Username,
		cfg.Password,
		cfg.Host)

	message := "Subject: [gosm] " + service.Name + " is " + service.Status + "\r\n"
	message += service.Name + " is now " + service.Status + "\r\n"
	message += "Protocol: " + service.Protocol + "\r\n"
	message += "Host: " + service.Host + "\r\n"
	if service.Port.Value != nil {
		message += "Port: " + strconv.FormatInt(service.Port.Int64, 10) + "\r\n"
	}
	err := smtp.SendMail(
		cfg.Host+":"+strconv.Itoa(cfg.Port),
		auth, cfg.EmailAddress,
		models.CurrentConfig.EmailRecipients,
		[]byte(message))
	if err != nil {
		fmt.Println(err)
	}
}
