package main

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/utils/mailer"
	"time"
)

func main() {
	mails := mailer.GetMailsQueue()
	cfg := config.GetConfig()

	fmt.Println(cfg.Mail)

	mail1 := mailer.Mail{
		To:      "leperiton@yandex.ru",
		Subject: "Ежедневное оповещение",
		Body:    "Как дела?",
	}

	mailer, _ := mailer.NewMailer(cfg)

	go mailer.Start(context.TODO(), mails)

	mails <- mail1
	mails <- mail1
	mails <- mail1
	mails <- mail1

	time.Sleep(time.Minute)
}
