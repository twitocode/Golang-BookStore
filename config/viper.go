package config

import "github.com/spf13/viper"

type EnvConfig struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	Port       string `mapstructure:"PORT"`
}

func LoadViperConfig(path string) (*EnvConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("server")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	var config EnvConfig
	err := viper.ReadInConfig()

	if err != nil {
    return nil, err
  }

	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}

  return &config, nil
}
