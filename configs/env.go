package configs

func GetWebsiteURL() string {
	cfg := ReadConfig()
	if !cfg.IsTesting {
		return "https://google.com" //TODO: Change to real URL
	} else {
		return "https://bing.com" // TODO: Change to real testing URL
	}
}

func GetApiURL() string {
	cfg := ReadConfig()
	if !cfg.IsTesting {
		return "https://api.oh-heck.dev"
	} else {
		return "http://localhost:6000"
	}
}
