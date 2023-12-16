package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// Arrange
	name := "Foo Bar"
	email := "teste@email"

	// Act
	user, err := NewUser(name, email, "123456")
	if err != nil {
		t.Errorf("error should be nil, but was %v", err)
	}

	// Assert
	if user.Name != name {
		t.Errorf("expected name %q, but got %q", name, user.Name)
	}
	if user.Email != email {
		t.Errorf("expected email %q, but got %q", email, user.Email)
	}
	if user.Password == "" {
		t.Errorf("expected password to be set, but it was empty")
	}
	if user.ID == uuid.Nil {
		t.Errorf("expected id to be set, but it was empty")
	}
}

func TestUserTestify(t *testing.T) {
	user, err := NewUser("Foo Bar", "teste@email.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Name, "Foo Bar")
	assert.Equal(t, user.Email, "teste@email.com")
	assert.NotEqual(t, user.Password, "123456")
}

func TestValidadePassword(t *testing.T) {
	passwd := "123456"
	user, err := NewUser("Foo Bar", "teste@email.com", passwd)
	validatedPassword := user.ValidatePassword(passwd)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, validatedPassword)
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}