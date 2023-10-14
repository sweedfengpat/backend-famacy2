package service
import (
    "server/app/config"
    "log"

    "gopkg.in/gomail.v2"
)

type Mailer struct{}

func (m Mailer) Send(message *gomail.Message) {
    message.SetHeader("From", "etracking.th@gmail.com")

    if err := config.Mailer.DialAndSend(message); err != nil {
        log.Panicln("[Mailer] ", err)
    }
}