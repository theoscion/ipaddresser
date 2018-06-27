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

func sendSNSNotification() {

}

func sendEmailNotification() {

}
