// // tipovi za usera, struct, tipovi za frontend i backend (a ne baza)

// package models

// import (
// 	"errors"
// 	"regexp"
// )

// type User struct {
// 	ID        int64
// 	Username  string
// 	Password  string
// 	Email     string
// 	FirstName string
// 	LastName  string
// 	Phone     string
// 	CreatedAt int64
// }

// func (u *User) Validate() error {
// 	// Validate ID
// 	if u.ID <= 0 {
// 		return errors.New("ID must be a positive integer")
// 	}

// 	// Validate Username
// 	if len(u.Username) < 4 {
// 		return errors.New("Username must be at least 4 characters long")
// 	}

// 	// Validate Password
// 	if len(u.Password) < 8 {
// 		return errors.New("Password must be at least 8 characters long")
// 	}

// 	// Validate Email
// 	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
// 	if !emailRegex.MatchString(u.Email) {
// 		return errors.New("Invalid email address")
// 	}

// 	// Validate FirstName
// 	if len(u.FirstName) < 2 {
// 		return errors.New("First name must be at least 2 characters long")
// 	}

// 	// Validate LastName
// 	if len(u.LastName) < 2 {
// 		return errors.New("Last name must be at least 2 characters long")
// 	}

// 	// Validate Phone
// 	phoneRegex := regexp.MustCompile(`^[0-9]{10}$`)
// 	if !phoneRegex.MatchString(u.Phone) {
// 		return errors.New("Phone number must be a 10-digit number")
// 	}

// 	return nil
// }