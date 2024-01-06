package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"

	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/jordan-wright/email"
)

// SMTP
type SMTP struct {
	e *email.Email
}

// New SMTP
func NewSMTP() *SMTP {
	return &SMTP{
		e: email.NewEmail(),
	}
}

// send mail
func (s *SMTP) Send(email Email, config map[string]string) (err error) {
	s.e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	s.e.To = email.To
	s.e.Bcc = email.Bcc
	s.e.Cc = email.Cc
	s.e.Subject = email.Subject
	s.e.Text = email.Text
	s.e.HTML = email.HTML

	// if attach exist then send mail with attach
	if len(email.Attachments) > 0 {
		for _, v := range email.Attachments {
			// Base64 Standard Decoding
			sDec, err := base64.StdEncoding.DecodeString(string(v.Content))
			if err != nil {
				return err
			}
			s.e.Attach(bytes.NewReader(sDec), v.Cid, v.ContentType)
		}
	}

	// send
	err = s.e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),
		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)

	// save mail information with log
	log.DebugJSON("Send Email", "Details", s.e)

	// if err then log
	if err != nil {
		log.ErrorJSON("Send Email", "Error", err)
	}

	return
}
