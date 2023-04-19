package main

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/server"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err := config.NewConfigBuilder().WithAllFromEnv().Build()
	if err != nil {
		logger.Error(err.Error())
	}

	myServer, err := server.New(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	ctxServer, cancel := context.WithCancel(context.Background())
	go func() {
		if err = myServer.Start(ctxServer); err != nil {
			logger.Info(err.Error())
			return
		}
	}()

	ctxOS, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Gracefully shut down
	<-ctxOS.Done()
	err = myServer.Shutdown(ctxServer)
	if err != nil {
		logger.Info(err.Error())
	}

	cancel()
	time.Sleep(2 * time.Second)
}
