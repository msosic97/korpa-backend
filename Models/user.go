// tipovi za usera, struct, tipovi za frontend i backend (a ne baza)

package models

import (
	"errors"
	"regexp"
)

type User struct {
	ID        int64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firsname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	CreatedAt int64 `json:"created_at"`
}


type LoginRequest struct {
	ID       int64
	Username string
	Password string
}


// User validation
func (u *User) Validate() bool {
	// Validate ID
	if u.ID <= 0 {
		// return errors.New("ID must be a positive integer")
		return false
	}

	// Validate Username
	if len(u.Username) < 4 || len(u.Username) > 20 {
		// return errors.New("username must be at least 4 characters long")
		return false
	}

	// Validate Password
	if len(u.Password) < 4 || len(u.Password) > 20{
		// return errors.New("password must be at least 8 characters long")
		return false
	}

	// Validate Email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		// return errors.New("invalid email address")
		return false
	}

	// Validate FirstName
	if len(u.FirstName) < 2 || len(u.FirstName) > 15{
		// return errors.New("first name must be at least 2 characters long")
		return false
	}

	// Validate LastName
	if len(u.LastName) < 2 || len(u.LastName) > 25{
		// return errors.New("last name must be at least 2 characters long")
		return false
	}

	// Validate Phone
	phoneRegex := regexp.MustCompile(`^[0-9]+$`)
	if !phoneRegex.MatchString(u.Phone) {
		// return errors.New("phone number must be a 10-digit number")
		return false
	}

	return true
}

// Validation for login

func (l *LoginRequest) Validate() error {
	// Validate ID
	if l.ID <= 0 {
		return errors.New("ID must be a positive integer")
	}

	// Validate Username
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{4,}$`)
	if !usernameRegex.MatchString(l.Username) {
		return errors.New("invalid username. It must contain only letters, numbers, and underscores, with a minimum length of 4 characters")
	}

	// Validate Password
	if len(l.Password) < 4 || len(l.Password) < 25{
		return errors.New("password must be at least 4 characters long")
	}

	return nil
}