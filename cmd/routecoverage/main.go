package main

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/utils/mailer"
	"os/signal"
	"syscall"
	"time"

	"github.com/jbakhtin/driving-school-route-coverage/internal/application"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg := config.GetConfig()
	if err != nil {
		logger.Error(err.Error())
	}

	myServer, err := application.New(*cfg)
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

	mailsQueue := mailer.GetMailsQueue()
	mailSender, err := mailer.NewMailer(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	go func() {
		if err = mailSender.Start(ctxServer, mailsQueue); err != nil {
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
