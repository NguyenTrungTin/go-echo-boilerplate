package user

import "github.com/go-playground/validator"

func Validate(u *User) (bool, string) {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			switch err.StructField() {
			case "Email":
				return false, "Invalid email address!"
			case "Username":
				return false, "Username cannot be empty!"
			case "FirstName":
				return false, "FirstName cannot be empty!"
			case "LastName":
				return false, "LastName cannot be empty!"
			case "Password":
				return false, "Password must be at least 6 characters long!"
			default:
				return true, ""
			}
		}
	}

	return true, ""
}

func ValidateWithoutPassword(u *User) (bool, string) {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			switch err.StructField() {
			case "Email":
				return false, "Invalid email address!"
			case "Username":
				return false, "Username cannot be empty!"
			case "FirstName":
				return false, "FirstName cannot be empty!"
			case "LastName":
				return false, "LastName cannot be empty!"
			default:
				return true, ""
			}
		}
	}

	return true, ""
}

func ValidateField(field string, fieldType string) (bool, string) {
	validate := validator.New()

	switch fieldType {
	case "email":
		err := validate.Var(field, "required,email")
		if err != nil {
			return false, "Invalid email address!"
		}
		return true, ""
	case "username":
		err := validate.Var(field, "required,min=3")
		if err != nil {
			return false, "Username must be at least 3 character long!"
		}
		return true, ""
	case "password":
		err := validate.Var(field, "required,min=6")
		if err != nil {
			return false, "Password must be at least 6 characters long!"
		}
		return true, ""
	}

	return false, "Validation false"
}
