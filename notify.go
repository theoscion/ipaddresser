package main

type snsAWSConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	TopicARN  string `json:"topicARN"`
}

type snsConfig struct {
	SNS     snsAWSConfig `json:"sns"`
	Subject string       `json:"subject"`
}

type emailSMTPConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type emailRecipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type emailConfig struct {
	SMTP emailSMTPConfig  `json:"smtp"`
	From emailRecipient   `json:"from"`
	To   []emailRecipient `json:"to"`
}

var snsPrepared = false

var emailPrepared = false

var sns snsConfig

var email emailConfig

func prepareSNSConfig() {
	if snsPrepared == false {
		snsPrepared = true
	}
}

func prepareEmailConfig() {
	if emailPrepared == false {
		emailPrepared = true
	}
}

func sendSNSNotification() {
	prepareSNSConfig()
}

func sendEmailNotification() {
	prepareEmailConfig()
}
