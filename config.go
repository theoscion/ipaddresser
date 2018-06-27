package main

// jsonConfig contains the configuration for the app
type jsonConfig struct {
	Daemon        bool       `json:"daemon"`
	QueryInterval int        `json:"interval"`
	Hook          configHook `json:"hook"`
}

// configHook contains the configuration for the HTTP hook
type configHook struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
	Method  string `json:"method"`
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
