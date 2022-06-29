package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIconfig
	DB  DBconfig
}

type APIconfig struct {
	port string
}

type DBconfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("API.port", 3000)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = APIconfig{
		port: viper.GetString("API.port"),
	}

	cfg.DB = DBconfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DBconfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.port
}