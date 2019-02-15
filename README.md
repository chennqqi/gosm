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
$ ./gosm -c /path/to/config.yml

# show useage
$ ./gosm -h
~~~ 

### build and run docker
~~~

# Build
$ cd $GOPATH/src/github.com/chennqqi/gosm
$ ./builddocker.sh

# Copy config.example.json to /path_to_gosm/config.yml
$ cp config.example.json config.yml

# Modify config.yml with your preferred settings
$ vi config.yml

# Run docker
$  docker run -d -p8030:8030 -v /path_to_gosm/:/data sort/gosm:1.2.0 -d /data/gosm.db -c /data/config.yml

~~~ 



### Current Build Version - 1.2
You can check the version of your build by running ``./gosm -h``

### Web UI

By default, the web UI listens on localhost on port 8030, and can be accessed in your browser at http://127.0.0.1:8030. The web UI currently gives you access to view the realtime status of your services, and add/edit/remove services. 

### Config

There is an example config file in the repo named config.example.json to use as a reference. Below is a brief description of each configuration item. All items are required unless explicity stated.
* **verbose** - Whether or not to print information to the console

	  verbose: true
	  web_ui_host: 127.0.0.1 # The host the web UI will listen on
	  web_ui_port: 8030 # The port the web UI will listen on
	  check_interval: 30s # How often to check each service that is in an online state (time.duration)
	  pending_offline_check_interval: 5s # How often to check each service is in a pending or offline state (time.duration)
	  max_concurrent_checks: 5 # The maximum concurrent checks
	  connection_timeout: 3s # Timeout threshold for checks (time.duration)
	  successful_http_status_codes: # Which HTTP/HTTPS status codes are considered successful. Any status code not listed will be considered a failure response
	  - 200
	  ignore_https_cert_errors: false # Whether or not to ignore HTTPS cert errors
	  failed_check_threshold: 5  # How many consecutive failed checks are needed to consider a service offline
	  send_smtp: true # - Whether or not to send alerts via email
	  smtp:
	    host: "" # The SMTP server host to send emails from
	    port: 0 #  The SMTP server port
	    email_address: "" # The email address to send from
	    username: "" # The username for the SMTP server
	    password: "" # The password for the SMTP server
	  email_recipients: # Recipients of email alerts
	  - xx@xx.com
	  send_sms: false #  Whether or not to send alerts via sms
	  qcloud_sms:
	    app_id: "" # qcloud sms appId
	    app_key: "" # qcloud sms appKey
	    tmpl_id: 0 # qcloud sms tmpl_id, should contain 4 parameters.
	  qcloud_enable: false # enable qcloud sms
	  twoii_sms:
	    twilio_account_sid: "" # Your Twilio Account SID
	    twilio_auth_token: "" # Your Twilio Auth Token
	    twilio_phone_number: "" # Your Twilio phone number to send the SMS alerts from
	  twoii_enable: false # enable send twoii sms
	  sms_recipients: [] # Recipients of sms alerts

### TODO
* Implement SMTP and SMTP-TLS checks
* Add optional limits for email/sms alerts per second
* Add service status history to web UI
* Refactor to store config items in sqlite database and modify through web UI instead of JSON file
* Add optional authentication
