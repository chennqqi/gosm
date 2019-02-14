# gosm
A program written in Golang for monitoring services using various protocols listed below. Alerts can be sent via SMTP and/or SMS (Twilio API/QCloud SMS).

![](http://i.imgur.com/Upsmhcy.png)


### Supported Protocols
* HTTP
* HTTPS
* ICMP
* TCP

### Installing from Source
~~~
# Download and install dependencies
$ go get -u github.com/chennqqi/gosm

# Build program
$ cd $GOPATH/src/github.com/chennqqi/gosm
$ ./build.sh

# Copy config.example.json to config.yml
$ cp config.example.json config.yml

# Modify config.yml with your preferred settings
$ vi config.yml

# Run program and supply path to config file
$ ./gosm /path/to/config.yml
~~~ 

### Current Build Version - 1.1
You can check the version of your build by running ``./gosm version``

### Web UI
By default, the web UI listens on localhost on port 8030, and can be accessed in your browser at http://127.0.0.1:8030. The web UI currently gives you access to view the realtime status of your services, and add/edit/remove services. 

### Config

There is an example config file in the repo named config.example.json to use as a reference. Below is a brief description of each configuration item. All items are required unless explicity stated.
* **verbose** - Whether or not to print information to the console

  verbose: true
  web_ui_host: 127.0.0.1
  web_ui_port: 8030
  check_interval: 30s
  pending_offline_check_interval: 5s
  max_concurrent_checks: 5
  connection_timeout: 3s
  successful_http_status_codes:
  - 200
  ignore_https_cert_errors: false
  failed_check_threshold: 5
  send_smtp: true
  smtp:
    host: ""
    port: 0
    email_address: ""
    username: ""
    password: ""
  email_recipients:
  - xx@xx.com
  send_sms: false
  qcloud_sms:
    app_id: ""
    app_key: ""
    tmpl_id: 0
  qcloud_enable: false
  twoii_sms:
    twilio_account_sid: ""
    twilio_auth_token: ""
    twilio_phone_number: ""
  twoii_enable: false
  sms_recipients: []
* **web_ui_host** - The host the web UI will listen on
* **web_ui_Port** - The port the web UI will listen on
* **check_interval** - How often to check each service that is in an online state (seconds)
* **pending_offline_interval** - How often to check each service is in a pending or offline state (seconds)
* **max_concurrent_checks** - The maximum concurrent checks
* **connection_timeout** - Timeout threshold for checks (milliseconds)
* **successful_http_status_codes** - Which HTTP/HTTPS status codes are considered successful. Any status code not listed will be considered a failure response
* **ignore_https_cert_errors** - Whether or not to ignore HTTPS cert errors
* **failed_check_threshold** - How many consecutive failed checks are needed to consider a service offline
* **send_email** - Whether or not to send alerts via email
* **email_recipients** - Recipients of email alerts
* **smtp_host** - The SMTP server host to send emails from
* **smtp_port** - The SMTP server port
* **smtp_email_address** - The email address to send from
* **smtp_username** - The username for the SMTP server
* **smtp_password** - The password for the SMTP server
* **send_sms** - Whether or not to send alerts via sms
* **sms_recipients** - Recipients of sms alerts
* **twilio_account_sid** - Your Twilio Account SID
* **twilio_auth_token** - Your Twilio Auth Token
* **twilio_from_number** - Your Twilio phone number to send the SMS alerts from


### TODO
* Implement SMTP and SMTP-TLS checks
* Add optional limits for email/sms alerts per second
* Add service status history to web UI
* Refactor to store config items in sqlite database and modify through web UI instead of JSON file
* Add optional authentication
