package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/jbakhtin/driving-school-route-coverage/internal/application"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/logger"
	"github.com/jbakhtin/driving-school-route-coverage/internal/utils/mailer"
	"github.com/jbakhtin/driving-school-route-coverage/pkg/closer"
)

func main() {
	osCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	lg, err := logger.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	cl, _ := closer.New(lg)

	var myServer *application.Server
	if myServer, err = application.New(osCtx, *cfg); err != nil {
		lg.Fatal(err.Error())
	}

	if err = myServer.Start(); err != nil {
		lg.Fatal(err.Error())
	}

	mailsQueue, err := mailer.GetMailsQueue()
	if err != nil {
		lg.Fatal(err.Error())
	}
	mailSender, err := mailer.NewMailer(cfg, lg)
	if err != nil {
		lg.Fatal(err.Error())
	}

	go func() {
		if err = mailSender.Start(osCtx, mailsQueue); err != nil {
			lg.Fatal(err.Error())
		}
	}()

	cl.Add(myServer.Shutdown)
	cl.Add(mailer.Shutdown)

	// Gracefully shut down
	<-osCtx.Done()
	withTimeout, cancel := context.WithTimeout(context.Background(), time.Second*cfg.ShutdownTimeout)
	defer cancel()

	err = cl.Close(withTimeout)
	if err != nil {
		lg.Error(err.Error())
	}
}
