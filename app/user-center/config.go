package main

import (
	"fmt"
	"github.com/eolinker/apinto-dashboard/config"
)

func GetConsoleUrl() string {
	
	return fmt.Sprintf("127.0.0.1:%d", config.GetConfigData().Port)
}
