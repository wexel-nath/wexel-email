package config

import (
	"github.com/spf13/viper"
)

func Configure() {

	viper.Set("PUBLIC_KEY_PATH", "keys/test.public.pem")

	// Heroku Port
	viper.SetDefault("PORT", "4000")
	viper.BindEnv("PORT")

	// Postgres URL
	viper.SetDefault("DATABASE_URL", "")
	viper.BindEnv("DATABASE_URL")

	// Mailgun
	viper.BindEnv("MAILGUN_DOMAIN")
	viper.BindEnv("MAILGUN_API_KEY")
	viper.BindEnv("MAILGUN_PUBLIC_KEY")
}

func GetPublicKeyPath() string {
	return viper.GetString("PUBLIC_KEY_PATH")
}

func GetPort() string {
	return viper.GetString("PORT")
}

func GetDatabaseURL() string {
	return viper.GetString("DATABASE_URL")
}

func GetMailgunDomain() string {
	return viper.GetString("MAILGUN_DOMAIN")
}

func GetMailgunApiKey() string {
	return viper.GetString("MAILGUN_API_KEY")
}

func GetMailgunPublicKey() string {
	return viper.GetString("MAILGUN_PUBLIC_KEY")
}
