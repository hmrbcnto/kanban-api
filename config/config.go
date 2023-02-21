package config

import "github.com/spf13/viper"

type Config struct {
	DbConfig DbConfig
}

// Config loader
func LoadConfig() (*Config, error) {
	dbConfig := loadDbConfig()

	err := dbConfig.validate()

	if err != nil {
		return nil, err
	}

	return &Config{
		DbConfig: dbConfig,
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
