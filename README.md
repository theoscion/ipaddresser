# ipaddresser
This is a Go-based, cross-platform app for retrieving, storing and/or monitoring the public IP address for a computer or network

## Prerequisites 
The only requirements is that you have [Golang](https://golang.org/) installed and have an active internet connection.

## Installation
To install ipaddresser, simply run the following command:

```bash
go get github.com/Theoscion/ipaddresser
go install github.com/Theoscion/ipaddresser
```

Once installed, you can run the following command:

```bash
ipaddresser
```

## Configuration
The app looks for a JSON string from STDIN. This string can controls various app settings and should match the following structure:

```json
{
	// Controls if the app should run as a daemon
	"daemon": false,

	// Controls the interval in minutes between each query to ipify to get the public IP address; only applicable when running as a daemon
	"interval": 1,

	// Specifies an HTTP endpoint to query when an IP address is loaded or changes
	"hook": {
		// Specifies if the HTTP hook is enabled
		"enabled": false,

		// Specifies the URL for the HTTP endpoint
		"url": "",

		// Specifies the HTTP method to use when pinging the HTTP endpoint
		"method": ""
	}
}
```

__Passing no configuration will result the app running as a single query (i.e., not a daemon) with no hook.__