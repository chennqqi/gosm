package checker

import (
	"fmt"
	"time"

	"github.com/chennqqi/gosm/alerts"
	"github.com/chennqqi/gosm/models"
)

var (
	checkChannel chan (*models.Service)
)

// Start Starts the service checker process
func Start() {
	checkChannel = make(chan *models.Service, models.CurrentConfig.MaxConcurrentChecks)
	go processChecks()
	go checkOnlineServices()
	checkPendingOfflineServices()
}

func checkOnlineServices() {
	for {
		for i := range models.CurrentServices {
			if len(models.CurrentServices) <= i {
				break
			}
			if models.CurrentServices[i].Status == models.Online {
				checkChannel <- &models.CurrentServices[i]
			}
		}
		time.Sleep(time.Duration(models.CurrentConfig.CheckInterval))
	}
}

func checkPendingOfflineServices() {
	for {
		for i := range models.CurrentServices {
			if len(models.CurrentServices) <= i {
				break
			}
			if models.CurrentServices[i].Status != models.Online {
				checkChannel <- &models.CurrentServices[i]
			}
		}
		time.Sleep(time.Duration(models.CurrentConfig.PendingOfflineCheckInterval))
	}
}

func processChecks() {
	for {
		service := <-checkChannel
		online := service.CheckService()
		if online {
			if service.Status == models.Offline {
				service.Status = models.Online
				service.UptimeStart = time.Now().Unix()
				go alerts.SendAlerts(service)
			} else if service.Status == models.Pending {
				service.Status = models.Online
				if models.CurrentConfig.Verbose {
					fmt.Println(service.Name + " is now in the " + service.Status + " state")
				}
			}
			service.FailureCount = 0
		} else {
			if service.Status == models.Online {
				service.Status = models.Pending
				service.FailureCount = 1
				if models.CurrentConfig.Verbose {
					fmt.Println(service.Name + " is now in the " + service.Status + " state")
				}
			} else if service.Status == models.Pending {
				service.FailureCount++
				if service.FailureCount >= models.CurrentConfig.FailedCheckThreshold {
					service.Status = models.Offline
					go alerts.SendAlerts(service)
				}
			}
		}
	}
}
