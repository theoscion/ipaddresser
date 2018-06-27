package main

type snsAWSJSONConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	TopicARN  string `json:"topicARN"`
}

type snsJSONConfig struct {
	SNS     snsAWSJSONConfig `json:"sns"`
	Subject string           `json:"subject"`
}

type emailSMTPJSONConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type emailRecipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type emailJSONConfig struct {
	SMTP emailSMTPJSONConfig `json:"smtp"`
	From emailRecipient      `json:"from"`
	To   []emailRecipient    `json:"to"`
}

func sendSNSNotification() {

}

func sendEmailNotification() {

}
