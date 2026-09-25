package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(input RegisterInput) (*User, error)
}
type RegisterInput struct {
	UserName string
	Email    string
	Password string
}
type userService struct {
	repo UserRepository
}

var ErrEmailTaken = errors.New("email already in use")

func NewUserService(repo UserRepository) UserService {
	return userService{
		repo: repo,
	}
}
func (us userService) Register(input RegisterInput) (*User, error) {
	existingUser, err := us.repo.GetByEmail(input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrEmailTaken
	}
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := User{
		Username:     input.UserName,
		Email:        input.Email,
		PasswordHash: string(hashedBytes),
	}
	err = us.repo.Create(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
