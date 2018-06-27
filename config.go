package main

// jsonConfig contains the configuration for the app
type jsonConfig struct {
	Daemon        bool       `json:"daemon"`   // Specifies if the app will run as a daemon
	QueryInterval int        `json:"interval"` // Specifies how many minutes to wait between each query to ipify
	Hook          configHook `json:"hook"`     // Specifies an HTTP hook configuration
}

// configHook contains the configuration for the HTTP hook
type configHook struct {
	Enabled bool   `json:"enabled"` // Specifies if the HTTP hook is enabled
	URL     string `json:"url"`     // Specifies the URL to use for the HTTP hook
	Method  string `json:"method"`  // Specifies the HTTP method (e.g., GET, POST) to use for the HTTP hook
}

// defaultConfig returns a default jsonConfig{}, which is used when no JSON configuration is provided
func defaultConfig() jsonConfig {
	// Return a JSON configuration that is not a daemon, has a query interval of 1 minute, and no hook
	return jsonConfig{
		Daemon:        false,
		QueryInterval: 1,
		Hook: configHook{
			Enabled: false,
		},
	}
}
