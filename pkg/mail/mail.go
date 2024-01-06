package mail

import (
	"fmt"
	"sync"

	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/util"
)

// Email info
type Email struct {
	From        From         // sender
	To          []string     // recipient
	Bcc         []string     // blind carbon copy
	Cc          []string     // carbon copy
	Subject     string       // subject
	Text        []byte       // Plaintext message (optional)
	HTML        []byte       // Html message (optional)
	Attachments []Attachment // attachments
}

// sender
type From struct {
	Address string // mail address
	Name    string // sender name
}

type Attachment struct {
	Filename    string // file name+".(png/jpeg/jpg)"
	Content     string // base64 encoded content
	ContentType string // ex: image/png
	Cid         string // id
}

type Mailer struct {
	Driver
}

// only new one time (singleton pattern)
var once sync.Once

// internalMailer use Mailer
var internalMailer *Mailer

func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: NewSMTP(),
		}
	})

	return internalMailer
}

// send mail
func (m *Mailer) Send(email Email, conf map[string]string) (err error) {
	if conf == nil {
		conf = make(map[string]string)
		conf["host"] = config.Config.Mail.HOST
		conf["port"] = config.Config.Mail.PORT
		conf["username"] = config.Config.Mail.USERNAME
		conf["password"] = config.Config.Mail.PASSWORD
	}

	setKey := []string{"host", "port", "username", "password"}
	if !util.CheckMapStrStr(setKey, conf) {
		return fmt.Errorf("send email: error: %v", "config is required")
	}

	return m.Driver.Send(email, conf)
}
