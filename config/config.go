package config

import (
	"fmt"
)

func AuthSecret() []byte {
	//TODO: could encode this
	return []byte(appConfig.Authentication.Secret)
}

func AuthEnabled() bool { return appConfig.Authentication.Enabled }

func Address() string {
	return fmt.Sprintf("0.0.0.0:%d", appConfig.Server.Port)
}

func Gaming() Game { return appConfig.Game }

func Scoring() Score { return appConfig.Score }

func StopperLimit() int {
	//All reels should be of same len
	return len(appConfig.Game.ReelsOfSymbols[0])
}
