# ipaddresser
This is a Go-based, cross-platform app for retrieving, storing and/or monitoring the public IP address for a computer or network

## Prerequisites 
The only requirements is that you have [Golang](https://golang.org/) installed and have an active internet connection.

## Installation
To install ipaddresser, simply run the following command:

```bash
go get github.com/Theoscion/ipaddresser
```

Once installed, you can run the following command:

```bash
ipaddresser
```

## Options
`-v` Specifies if logs should be verbose
`-d` Specifies to run as a daemon to watch for IP address changes
`-i=` **integer** Specifies how often the daemon should check for IP address changes; defaults to 1 minute
`-log=` **string** Specifies an optional file to output all logs to instead of STDOUT
`-out=` **string** Specifies an optional file to output the current IP address to, as plain text
`-sns=` **string** Specifies an optional JSON configuration to use AWS SNS for IP address change notifications
`-email=` **string** Specifies an optional JSON configuration to use an email for IP address change notifications

## SNS Configuration
The app can be configured to publish a message to an AWS SNS topic when it detects a change to the IP address. Specify the path to a JSON configuration file using the `-sns=` option. The configuration should match the following structure:

```json
{
	// Contains information for publishing to AWS SNS
	"sns": {
		"accessKey": "",     // Access key for the IAM role with SNS publishing permissions
		"secretKey": "",     // Secret key for the IAM role with SNS publishing permissions
		"topicARN": "",     // The ARN for the topic to publish to
	},

	// Contains the subject to use when publishing to the SNS topic
	"subject": ""
}
```

## Email Configuration
The app can be configured to send emails when it detects a change to the IP address. Specify the path to a JSON configuration file using the `-email=` option. The configuration should match the following structure:

```json
{
	// Contains information for the SMTP server
	"smtp": {
		"host": "",     // Hostname of the SMTP server
		"port": "",     // Port for the SMTP server
		"username": "", // Username for the SMTP server, if applicable
		"password": ""  // Password for the SMTP server, if applicable
	},
	
	// Contains information for who the email should be sent from
	"from": {
		"name": "",  // Name for who the email should be sent from
		"email": "" // Email address for who the email should be sent from
	},

	// Contains information for all recipients that should receive email
	"to": [
		{
			"name": "",  // Name for who the email should be sent to
			"email": "" // Email address for who the email should be sent to
		},
		{
			"name": "",  // Name for who the email should be sent to
			"email": "" // Email address for who the email should be sent to
		}
	]
}
```