package main

import (
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/logger"
)

func main() {
	config := config.GetConfig()
	logger := logger.New(*config)

	for i := 0; i < 20000; i++ {
		logger.Info("logger construction succeeded 2")
		logger.Error("logger construction succeeded")
	}

}
