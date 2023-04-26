package users

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	URLPrefix = "http://localhost:3000/"
}

func TestUser_Validate(t *testing.T) {
	var user User

	t.Run("first name is blank", func(t *testing.T) {
		user = User{
			LastName: "Shnq",
			Email:    "adam@shnq.com",
			Phone:    "0798099158",
		}
		assert.Equal(t, ErrUserFirstNameBlank, user.Validate())
	})

	t.Run("last name is blank", func(t *testing.T) {
		user = User{
			FirstName: "mohammad",
			Email:     "adam@shnq.com",
			Phone:     "0798099158",
		}
		assert.Equal(t, ErrUserLastNameBlank, user.Validate())
	})

	t.Run("email is blank", func(t *testing.T) {
		user = User{
			FirstName: "mohammad",
			LastName:  "Shnq",
			Phone:     "0798099158",
		}
		assert.Equal(t, ErrUserEmailBlank, user.Validate())
	})

	t.Run("phone is blank", func(t *testing.T) {
		user = User{
			FirstName: "mohammad",
			LastName:  "Shnq",
			Email:     "adam@shnq.com",
		}
		assert.Equal(t, ErrUserPhoneBlank, user.Validate())
	})

	t.Run("valid", func(t *testing.T) {
		user = User{
			FirstName: "mohammad",
			LastName:  "Shnq",
			Email:     "adam@shnq.com",
			Phone:     "0798099158",
		}
		assert.Nil(t, user.Validate())
	})
}

func TestUser_MarshalJSON(t *testing.T) {
	var (
		user = User{
			FirstName: "mohammad",
			LastName:  "Shnq",
			Email:     "adam@shnq.com",
			Phone:     "0798099158",
		}
		encoded, err = json.Marshal(user)
	)

	assert.Nil(t, err)
	assert.JSONEq(t, `{
				"created_at":"0001-01-01T00:00:00Z", 
				"email":"adam@shnq.com", "first_name":"mohammad", 
				"id":"00000000-0000-0000-0000-000000000000", 
				"last_name":"Shnq", 
				"password":"", 
				"phone":"0798099158", 
				"refresh_token":"", 
				"token":"", 
				"updated_at":"0001-01-01T00:00:00Z", 
				"url":"http://localhost:3000/00000000-0000-0000-0000-000000000000", 
				"user_type":""
			}`, string(encoded))
}
