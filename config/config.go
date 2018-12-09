package config

import (
	"fmt"
)

func Address() string {
	return fmt.Sprintf("0.0.0.0:%d", appConfig.Server.Port)
}

func Gaming() Game { return appConfig.Game }

func Scoring() Score { return appConfig.Score }

func StopperLimit() int {
	//All reels should be of same len
	return len(appConfig.Game.ReelsOfSymbols[0])
}

func PositionStopper() Stopper {
	return appConfig.Stopper
}
