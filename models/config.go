package models

import (
	"os"
	"time"

	"github.com/chennqqi/goutils/utime"
	"github.com/chennqqi/goutils/yamlconfig"
)

type QCloudSMS struct {
	AppID  string `json:"app_id" yaml:"app_id"`
	AppKey string `json:"app_key" yaml:"app_key"`
	TmplId uint   `json:"tmpl_id" yaml:"tmpl_id"`
}

type TwoiiSMS struct {
	TwilioAccountSID  string `json:"twilio_account_sid" yaml:"twilio_account_sid"`
	TwilioAuthToken   string `json:"twilio_auth_token" yaml:"twilio_auth_token"`
	TwilioPhoneNumber string `json:"twilio_phone_number" yaml:"twilio_phone_number"`
}

type SMTP struct {
	Host         string `json:"host" yaml:"host"`
	Port         int    `json:"port" yaml:"port"`
	EmailAddress string `json:"email_address" yaml:"email_address"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
}

// Config The application configuration and settings
type Config struct {
	Verbose                     bool           `json:"verbose" yaml:"verbose"`
	WebUIHost                   string         `json:"web_ui_host" yaml:"web_ui_host"`
	WebUIPort                   int            `json:"web_ui_port" yaml:"web_ui_port"`
	CheckInterval               utime.Duration `json:"check_interval" yaml:"check_interval"`
	PendingOfflineCheckInterval utime.Duration `json:"pending_offline_check_interval" yaml:"pending_offline_check_interval"`
	MaxConcurrentChecks         int            `json:"max_concurrent_checks" yaml:"max_concurrent_checks"`
	ConnectionTimeout           utime.Duration `json:"connection_timeout" yaml:"connection_timeout"`
	SuccessfulHTTPStatusCodes   []int          `json:"successful_http_status_codes" yaml:"successful_http_status_codes"`
	IgnoreHTTPSCertErrors       bool           `json:"ignore_https_cert_errors" yaml:"ignore_https_cert_errors"`
	FailedCheckThreshold        int            `json:"failed_check_threshold" yaml:"failed_check_threshold"`

	SendSMTP        bool     `json:"send_smtp" yaml:"send_smtp"`
	Smtp            SMTP     `json:"smtp" yaml:"smtp"`
	EmailRecipients []string `json:"email_recipients" yaml:"email_recipients"`

	SendSMS       bool      `json:"send_sms" yaml:"send_sms"`
	Qcloud        QCloudSMS `json:"qcloud_sms" yaml:"qcloud_sms"`
	QcloudEnable  bool      `json:"qcloud_enable" yaml:"qcloud_enable"`
	Twoii         TwoiiSMS  `json:"twoii_sms" yaml:"twoii_sms"`
	TwoiiEnable   bool      `json:"twoii_enable" yaml:"twoii_enable"`
	SMSRecipients []string  `json:"sms_recipients" yaml:"sms_recipients"`
}

var (
	// CurrentConfig The current configuration
	CurrentConfig Config
)

// ParseConfigFile Parses the config.json file
func ParseConfigFile() Config {
	if len(os.Args) < 2 {
		panic("Expected run syntax: './gosm /path/to/config.yaml'")
	}
	var cfg Config
	err := yamlconfig.Load(&cfg, os.Args[1])
	if os.IsNotExist(err) {
		yamlconfig.Save(Config{
			Verbose:                     true,
			WebUIHost:                   "127.0.0.1",
			WebUIPort:                   8030,
			CheckInterval:               utime.Duration(30 * time.Second),
			PendingOfflineCheckInterval: utime.Duration(5 * time.Second),
			MaxConcurrentChecks:         5,
			ConnectionTimeout:           utime.Duration(time.Millisecond * 3000),
			SuccessfulHTTPStatusCodes:   []int{200},
			IgnoreHTTPSCertErrors:       false,
			FailedCheckThreshold:        5,
			SendSMTP:                    true,
			EmailRecipients:             []string{"xx@xx.com"},
		}, "")
	}
	return cfg
}
