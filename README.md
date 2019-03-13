# ipaddresser
This is a Go-based application for retrieving and monitoring the public IP address for a computer or network, using the API offered by [Ipify](https://www.ipify.org/).

This application can also run any number of webhooks (i.e., HTTP requests) whenever the IP address changes.

## Prerequisites 
To build and install ipaddresser from source, you will need the following prerequisites:

* [Git](https://git-scm.com/)
* [Golang](https://golang.org/)
* [Dep](https://github.com/golang/dep)

To run ipaddresser, you just need an active internet connection

## Installation
If you want to download a binary, head over to [releases](https://github.com/theoscion/ipaddresser/releases) and pull down the binary for your OS.

To build from source, you can do so using the standard Go commands:

```bash
git clone github.com/Theoscion/ipaddresser
cd ./ipaddresser
git checkout v1.0.1
dep ensure
go install
ipaddresser
```

_This assumes that your `$GOBIN` is in your `$PATH`._

## Configuration
The app looks for various flags and arguments to control how it runs.

### Flags

Here are all of the supported flags and a description of what they do

| Flag                     | Description                                                                                                   |
|--------------------------|---------------------------------------------------------------------------------------------------------------|
| `--verbose` or `-v`      | Enables verbose logging output so you can see what is happening                                               |
| `--daemon` or `-d`       | Tells the app to run as a daemon until the user breaks (e.g., Ctrl+C); the daemon runs a request every minute |
| `--always-hook` or `-a`  | Tells the app to run all webhooks, even if the IP address hasn't changed                                      |

_Order does not matter when specifying a flag._

### Arguments

All arguments that aren't flags are used to define webhook URLs. Each argument would be an endpoint that should be triggered. When the application is run, it will fire a POST request to each webhook URL specified

For example:

```
ipaddresser -v -d https://example.com/my-ip-monitor https://example.com/notification-service
```

The example command above will send a request to 2 webhook URLs every, but only when the IP address changes. The request body submitted to each webhook will look like this (as an example):

```json
{
    "source": "ipaddresser",
    "firstCheck": false,
    "oldIP": "1.2.3.4",
    "newIP": "4.3.2.1"
}
```

The `firstCheck` property specifies if it is the first time the IP is being reported.  When running as a single request, that property will always be `true`. When running as a daemon, it will be `true` the first time all webhooks are run; all subsequent runs, the value will be `false`. When `firstCheck` property is `false`, the `oldIP` will always be an empty string. The `oldIP` property will contain the last known IP address; this will only have a value when running as a daemon. The `newIP` property will contain the IP address returned from the IP service (i.e., Ipify).