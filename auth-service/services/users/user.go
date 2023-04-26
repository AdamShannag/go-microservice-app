package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	// URLPrefix to be returned when encoding user.
	URLPrefix = os.Getenv("URL") + "user/"
	// ErrUserFirstNameBlank validation error.
	ErrUserFirstNameBlank = errors.New("First Name can't be blank")
	// ErrUserLastNameBlank validation error.
	ErrUserLastNameBlank = errors.New("Last Name can't be blank")
	// ErrUserEmailBlank validation error.
	ErrUserEmailBlank = errors.New("Email can't be blank")
	// ErrUserPhoneBlank validation error.
	ErrUserPhoneBlank = errors.New("Phone Name can't be blank")
)

// User represent a record stored in todos table.
type User struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name" validate:"required, min=2, max=100"`
	LastName     string    `json:"last_name" validate:"required, min=2, max=100"`
	Password     string    `json:"password" validate:"required, min=6"`
	Email        string    `json:"email" validate:"email, required"`
	Phone        string    `json:"phone" validate:"required"`
	Token        string    `json:"token"`
	UserType     string    `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Validate user.
func (u User) Validate() error {
	var err error
	switch {
	case len(u.FirstName) == 0:
		err = ErrUserFirstNameBlank
	case len(u.LastName) == 0:
		err = ErrUserLastNameBlank
	case len(u.Email) == 0:
		err = ErrUserEmailBlank
	case len(u.Phone) == 0:
		err = ErrUserPhoneBlank
	}
	return err
}

// MarshalJSON implement custom marshaller to marshal url.
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User

	return json.Marshal(struct {
		Alias
		URL string `json:"url"`
	}{
		Alias: Alias(u),
		URL:   fmt.Sprint(URLPrefix, u.ID),
	})
}
