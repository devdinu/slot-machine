package config

import (
	"fmt"
)

func AuthSecret() string { return appConfig.Authentication.Secret }

func AuthEnabled() bool { return appConfig.Authentication.Enabled }

func Address() string {
	return fmt.Sprintf("0.0.0.0:%d", appConfig.Server.Port)
}

func Gaming() Game { return appConfig.Game }

func Scoring() Score { return appConfig.Score }
