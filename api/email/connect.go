package email

import (
	"crypto/tls"
	"strconv"

	"github.com/go-mail/mail"
	"github.com/nguyentrungtin/go-echo-boilerplate/config"
)

// Connect is used to open Email SMTP connection to server
func Connect() *mail.Dialer {
	port, _ := strconv.Atoi(config.EMAIL_PORT)
	d := mail.NewDialer(config.EMAIL_SERVER, port, config.EMAIL_USER, config.EMAIL_PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d
}
