package configs

func GetWebsiteURL() string {
	return "https://oh-heck.dev"
}

func GetApiURL() string {
	cfg := ReadConfig()
	if !cfg.IsTesting {
		return "https://api.oh-heck.dev"
	} else {
		return "http://localhost:6000"
	}
}
