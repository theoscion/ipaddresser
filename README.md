# ipaddresser
This is a Go-based, cross-platform app for retrieving, storing and/or monitoring the public IP address for a computer or network, using the API offered by [Ipify](https://www.ipify.org/).

## Prerequisites 
You must have [Golang](https://golang.org/) installed and there must be an active internet connection.

## Installation
To install ipaddresser, simply run the following command:

```bash
go get github.com/Theoscion/ipaddresser
go install github.com/Theoscion/ipaddresser
```

To install at a custom location (e.g., `/usr/local/bin`), you can run:

```bash
go build -o /usr/local/bin/ipaddresser github.com/Theoscion/ipaddresser
```

Once installed, you can run the following command:

```bash
ipaddresser
```

To update the package, you can follow standard updating procedures for Go packages:

```bash
go get -u github.com/Theoscion/ipaddresser
```

Once that is complete, you can repeat the `go install` or `go build` commands listed above.

## Configuration
The app looks for a JSON string from STDIN. This string controls various app settings and should match the following structure:

```json
{
	"daemon": false,
	"interval": 1,
	"hook": {
		"enabled": false,
		"submitForSameIP": false,
		"url": "",
		"method": "",
	}
}
```

_Passing no configuration will result in the app running as a single query (i.e., not a daemon) with no hook._