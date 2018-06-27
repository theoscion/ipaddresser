package main

type jsonConfig struct {
	Daemon        bool       `json:"daemon"`
	QueryInterval int        `json:"interval"`
	Hook          configHook `json:"hook"`
}

type configHook struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
	Method  string `json:"method"`
}

func defaultConfig() jsonConfig {
	return jsonConfig{
		Daemon:        false,
		QueryInterval: 1,
		Hook: configHook{
			Enabled: false,
		},
	}
}
