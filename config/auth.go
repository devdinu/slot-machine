package config

import "time"

type Authentication struct {
	Enabled       bool
	Secret        string
	ExpiryMinutes int `yaml:"expiry_minutes"`
}

func AuthEnabled() bool {
	return appConfig.Authentication.Enabled
}

func AuthTokenExpiryMinutes() time.Duration {
	return time.Minute * time.Duration(appConfig.Authentication.ExpiryMinutes)
}

func AuthSecret() []byte {
	//TODO: could encode this
	return []byte(appConfig.Authentication.Secret)
}
