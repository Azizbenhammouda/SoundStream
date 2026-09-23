package users

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) error
	GetByID(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db: db,
	}
}
func (u userRepository) Create(user *User) error {
	result := u.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u userRepository) GetByID(id uuid.UUID) (*User, error) {
	var user User
	result := u.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (u userRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (u userRepository) Update(user *User) error {
	result := u.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u userRepository) Delete(id uuid.UUID) error {
	result := u.db.Where("id = ?", id).Delete(&User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
