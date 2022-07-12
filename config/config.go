package config

import "os"

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type MongoConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

type Config struct {
	ApiConfig
	MongoConfig
}

func (c Config) readConfig() Config {
	c.MongoConfig = MongoConfig{
		Host: os.Getenv("MONGO_HOST"), //localhost
		Port : os.Getenv("MONGO_PORT"), //27017
		DbName : os.Getenv("MONGO_DB"), //enigma
		User: os.Getenv("MONGO_USER"), //stevejo
		Password: os.Getenv("MONGO_PASSWORD"), //password
	}

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"), //localhost
		ApiPort: os.Getenv("API_PORT"), //8888
	}

	return c

}

func NewConfig() Config{
	cfg := Config{}
	return cfg.readConfig()
}