package email

import (
	"fmt"

	"github.com/nguyentrungtin/go-echo-boilerplate/config"

	"github.com/go-mail/mail"
)

func SendResetPassword(email string, name string, token string) error {
	d := Connect()
	m := mail.NewMessage()

	reset_password_url := fmt.Sprintf("%s/reset_password?token=%s", config.FRONTEND_URL, token)
	message := fmt.Sprintf("Hi %s, <br/><br/>A password reset for your account was requested.<br/>Please click the link below to change your password.<br/>Note that this link is valid for 24 hours. After the time limit has expired, you will have to resubmit the request for a password reset.<br/>Reset password link: %s <br/>", name, reset_password_url)

	m.SetHeader("From", "admin@your_app.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "YOUR_APP - Reset Password!")
	m.SetBody("text/html", message)

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(fmt.Sprintf("RESET PASSWORD - Email is sent to %s false!", email))
		fmt.Println(err)
		return err
	}

	fmt.Println(fmt.Sprintf("RESET PASSWORD - Email is sent to %s successfully!", email))
	return nil
}
