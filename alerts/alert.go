package alerts

import (
	"fmt"

	"github.com/chennqqi/gosm/models"
)

// SendAlerts Sends the alerts for a services current status
func SendAlerts(service *models.Service) {
	var err error
	if models.CurrentConfig.Verbose {
		fmt.Println(service.Name + " is now in the " + service.Status + " state")
	}
	if models.CurrentConfig.SendSMTP {
		err = sendSMTPAlert(service)
		if err!=nil {
			fmt.Println("SEND SMTP ERROR:", err)
		}
	}
	if models.CurrentConfig.SendSMS {
		err = sendSMSAlert(service)
		fmt.Println("SEND SMS ERROR:", err)
	}
}
