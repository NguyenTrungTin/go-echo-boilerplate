package email

import "github.com/go-playground/validator"

// Validate is used to validate the email address is valid or not
func Validate(email string) error {
	validate := validator.New()
	err := validate.Var(email, "omitempty,email")
	return err
}
