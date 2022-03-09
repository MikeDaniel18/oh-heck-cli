package configs

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path"
)

type Config struct {
	ApiKey    string
	IsTesting bool
}

func configPath() string {
	cfgFile := ".oh-heck"
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, cfgFile)
}

func SaveConfig(c Config) {
	jsonC, _ := json.Marshal(c)
	ioutil.WriteFile(configPath(), jsonC, 0666)
}

func ReadConfig() *Config {
	data, err := ioutil.ReadFile(configPath())
	if err != nil {
		// fmt.Println(err.Error())
		return new(Config)
	} else {
		var cfg Config
		json.Unmarshal(data, &cfg)
		return &cfg
	}
}

func (c Config) Save() {
	SaveConfig(c)
}
