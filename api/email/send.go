package email

import (
	"fmt"

	"github.com/go-mail/mail"
)

// Send is used to send Email to the provided email address
func Send(email, session, year, make, model string) error {
	d := Connect()
	m := mail.NewMessage()

	m.SetHeader("From", "admin@your_app.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "YOUR_APP")
	m.SetBody("text/html", "Hello from Go!")

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s - Email is sent to %s false!", session, email))
		fmt.Println(err)
		return err
	}

	fmt.Println(fmt.Sprintf("%s - Email is sent to %s successfully!", session, email))
	return nil
}
