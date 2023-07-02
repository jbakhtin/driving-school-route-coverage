package mailer

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/logger"
	"github.com/jbakhtin/driving-school-route-coverage/internal/utils/ratelimiter"
	"net/smtp"
	"sync"
	"time"
)

type Mail struct {
	To string
	Subject string
	Body string
}

func NewMail(to, subject, body string) *Mail {
	return &Mail{
		to,
		subject,
		body,
	}
}

var mails chan Mail
var once sync.Once

func GetMailsQueue() (chan Mail, error) {
	var err error
	once.Do(func() {
		cfg, err := config.GetConfig()
		if err != nil {
			return
		}
		mails = make(chan Mail, cfg.Mail.QueueSize)
	})
	return mails, err
}

type Mailer struct {
	config *config.Config
	logger *logger.Logger
}

func NewMailer(cfg *config.Config, logger *logger.Logger) (*Mailer, error) {
	return &Mailer{
		cfg,
		logger,
	}, nil
}

func (m *Mailer) Start(ctx context.Context, mails chan Mail) error {
	limiter := ratelimiter.New(time.Second, m.config.Mail.SendPerSecond)
	err := limiter.Run(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case mail := <-mails:
			limiter.Wait()

			go func() {
				err := m.Send(&mail)
				if err != nil {
					m.logger.Error(err.Error())
				}
			}()
		}
	}
}

func (m *Mailer) Send(mail *Mail) error {
	auth := smtp.PlainAuth("", m.config.Mail.UserName, m.config.Mail.UserPassword, m.config.Mail.Host)
	smtpAddr := fmt.Sprintf("%v:%v", m.config.Mail.Host, m.config.Mail.Port)

	msg := []byte("To: " + mail.To + "\r\n" +
		"Subject: " + mail.Subject + "\r\n" +
		"\r\n" + mail.Body + "\r\n")

	err := smtp.SendMail(smtpAddr, auth, m.config.Mail.FromAddress, []string{mail.To}, msg)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func Shutdown(ctx context.Context) error {
	fmt.Println("Mailer shut down")
	timer := time.NewTimer(time.Second * 12)
	<-timer.C
	return nil
}
