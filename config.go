package chatbot

type Config struct {
	QaPath    string `json:"qa_path"`
	RobotPath string `json:"robot_path"`
	LogPath   string `json:"log_path"`
}

var config *Config

func LoadConfig(filename string) *Config {
	conf := &Config{}
	if err := DecodeJsonFromFile(&conf, filename); err != nil {
		panic(err)
	}

	return conf
}

func getConfig() *Config {
	return config
}
