package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username        string    `json:"username" gorm:"unique;not null"`
	Email           string    `json:"email" gorm:"unique;not null"`
	PasswordHash    string    `json:"-"`
	DisplayName     string    `json:"display_name"`
	ProfileImageURL *string   `json:"profile_image_url,omitempty" gorm:"column:profile_image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
