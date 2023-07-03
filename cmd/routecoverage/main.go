package main

import (
	"context"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/logger"
	"github.com/jbakhtin/driving-school-route-coverage/internal/utils/mailer"
	"github.com/jbakhtin/driving-school-route-coverage/pkg/closer"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	osCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	closer, _ := closer.New(logger)

	var myServer *application.Server
	if myServer, err = application.New(*cfg); err != nil {
		logger.Fatal(err.Error())
	}

	if err = myServer.Start(); err != nil {
		logger.Fatal(err.Error())
	}

	mailsQueue, err := mailer.GetMailsQueue()
	if err != nil {
		logger.Fatal(err.Error())
	}
	mailSender, err := mailer.NewMailer(cfg, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	go func() {
		if err = mailSender.Start(osCtx, mailsQueue); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	closer.Add(myServer.Shutdown)
	closer.Add(mailer.Shutdown)

	// Gracefully shut down
	<-osCtx.Done()
	withTimeout, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	err = closer.Close(withTimeout)
	if err != nil {
		logger.Error(err.Error())
	}
}
