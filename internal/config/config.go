package config

type Config struct {
	DB              string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"
