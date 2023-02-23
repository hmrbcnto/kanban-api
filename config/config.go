package config

import "github.com/spf13/viper"

type Config struct {
	DbConfig  DbConfig
	EnvConfig EnvConfig
}

// Config loader
func LoadConfig() (*Config, error) {
	dbConfig := loadDbConfig()
	envConfig := loadEnvConfig()

	err := dbConfig.validate()

	if err != nil {
		return nil, err
	}

	err = envConfig.validate()

	return &Config{
		DbConfig:  dbConfig,
		EnvConfig: envConfig,
	}, nil
}

// Db Config LOader
func loadDbConfig() DbConfig {
	viper.SetEnvPrefix("MONGO")
	viper.BindEnv("URI")

	return DbConfig{
		MongoURI: viper.GetString("URI"),
	}
}

// Env Config Loader
func loadEnvConfig() EnvConfig {
	viper.BindEnv("SECRET_KEY")

	return EnvConfig{
		SecretKey: viper.GetString("SECRET_KEY"),
	}
}
