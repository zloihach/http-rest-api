package model

import "testing"

// TestUser_Validate tests the User.Validate() method
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "test@test.com",
		Password: "password",
	}
}
