package mailer

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
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

func NewMail(to string, subject string, body string) *Mail {
	return &Mail{
		to,
		subject,
		body,
	}
}

var mails chan Mail
var once sync.Once

func GetMailsQueue() chan Mail {
	once.Do(func() {
		mails = make(chan Mail, 5)
	})
	return mails
}

type Mailer struct {
	config *config.Config
}

func NewMailer(cfg *config.Config) (*Mailer, error) {
	return &Mailer{
		cfg,
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
					fmt.Println(err)
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
