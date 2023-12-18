package database

// Integration test based on in memory database

import (
	"testing"

	"github.com/luizhenrique-dev/go-products-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepository_Create(t *testing.T) {
	// Arrange
	name := "Foo Bar"
	email := "test@email.com"
	db := newInMemoryDatabase()
	userRepository := NewUserRepository(db)
	user, _ := entity.NewUser(name, email, "123456")

	// Act
	err := userRepository.Create(user)
	
	assert.Nil(t, err)

	var userFromDB entity.User
	err = db.First(&userFromDB, user.ID).Error

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, userFromDB)
	assert.Equal(t, user.ID, userFromDB.ID)
	assert.Equal(t, user.Name, userFromDB.Name)
	assert.Equal(t, user.Email, userFromDB.Email)
	assert.NotEmpty(t, userFromDB.Password)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	// Arrange
	name := "Foo Bar"
	email := "test@email.com"
	db := newInMemoryDatabase()
	userRepository := NewUserRepository(db)
	user, _ := entity.NewUser(name, email, "123456")
	err := userRepository.Create(user)
	assert.Nil(t, err)
	
	// Act
	userFromDB, err := userRepository.FindByEmail(email)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, userFromDB)
	assert.Equal(t, user.ID, userFromDB.ID)
	assert.Equal(t, user.Name, userFromDB.Name)
	assert.Equal(t, user.Email, userFromDB.Email)
	assert.NotEmpty(t, userFromDB.Password)
}

func newInMemoryDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}