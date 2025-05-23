package config

type AppConfig struct {
	SEGA_ID string `mapstructure:"token" env:"SEGA_ID" default:""`
	SEGA_PW string `mapstructure:"prefix" env:"SEGA_PW" default:"!"`
}
