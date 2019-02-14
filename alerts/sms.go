package alerts

import (
	"fmt"
	"time"

	"github.com/chennqqi/gosm/models"
	"github.com/sfreiberg/gotwilio"

	"github.com/qichengzx/qcloudsms_go"
)

func sendSMSAlert(service *models.Service) error {
	cfg := &models.CurrentConfig
	if cfg.QcloudEnable {
		return sendQcloudSMS(service)
	} else if cfg.TwoiiEnable {
		return sendTwoiiSMSAlert(service)
	}
	return nil
}

func sendTwoiiSMSAlert(service *models.Service) error {
	cfg := &models.CurrentConfig.Twoii

	twilio := gotwilio.NewTwilioClient(cfg.TwilioAccountSID, cfg.TwilioAuthToken)
	for _, number := range models.CurrentConfig.SMSRecipients {
		_, exception, err := twilio.SendSMS(cfg.TwilioPhoneNumber, number, "[gosm] "+service.Name+" ("+service.Protocol+") is now "+service.Status, "", "")
		if exception != nil {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func sendQcloudSMS(service *models.Service) error {
	cfg := &models.CurrentConfig.Qcloud
	now := time.Now()

	//var msg = "[gosm] " + service.Name + " (" + service.Protocol + ") is now " + service.Status
	var day = now.Format("2006.01.02")
	var t = now.Format("15:04:05")

	var params = []string{service.Name, day, t, service.Protocol, service.Status }

	opt := qcloudsms.NewOptions(cfg.AppID, cfg.AppKey, "")

	var client = qcloudsms.NewClient(opt)
	client.SetDebug(true)

	var tels = make([]qcloudsms.SMSTel, len(models.CurrentConfig.SMSRecipients))

	for i := 0; i < len(models.CurrentConfig.SMSRecipients); i++ {
		tel := &tels[i]
		tel.Mobile = models.CurrentConfig.SMSRecipients[i]
		tel.Nationcode = "86"
	}

	_, err := client.SendSMSMulti(qcloudsms.SMSMultiReq{
		Tel:    tels,
		Type:   0,
		Sign:   "", //短信签名，如果使用默认签名，该字段可缺省
		TplID:  cfg.TmplId,
		Params: params,
	//	Msg:    msg,//没有这个字段
		Sig:    "", //App 凭证
		Time:   time.Now().Unix(),
		Extend: "",
		Ext:    "",
	})

	return err
}
