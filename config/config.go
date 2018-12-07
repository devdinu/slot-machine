package config

import (
	"fmt"
)

func Address() string {
	return fmt.Sprintf("0.0.0.0:%d", appConfig.Server.Port)
}
