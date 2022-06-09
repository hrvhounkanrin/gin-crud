package config

import "github.com/spf13/viper"

type Config struct {
	Port     string `mapstructure:"PORT"`
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"DBUSER"`
	Password string `mapstructure:"PASSWORD"`
	Dbport   string `mapstructure:"DBPORT"`
	Dbname   string `mapstructure:"DBNAME"`
}

func LoadConfig() (c Config, err error) {
	/*viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")*/
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
