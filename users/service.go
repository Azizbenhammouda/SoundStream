package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(input RegisterInput) (*User, error)
	Login(email, password string) (string, error)
}
type RegisterInput struct {
	UserName string
	Email    string
	Password string
}
type userService struct {
	repo      UserRepository
	jwtSecret string
}

var ErrInvalidCredentials = errors.New("Email or Password are invalid")
var ErrEmailTaken = errors.New("email already in use")

func NewUserService(repo UserRepository, secret string) UserService {
	return userService{
		repo:      repo,
		jwtSecret: secret,
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

func (us userService) Login(email, password string) (string, error) {
	user, err := us.repo.GetByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}
	token, err := GenerateToken(user.ID, us.jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
