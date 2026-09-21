package users

import "time"

type User struct {
	ID                int       `json:"id" gorm:"primaryKey"`
	Username          string    `json:"username" gorm:"unique;not null"`
	Email             string    `json:"email" gorm:"unique;not null"`
	Password_hash     string    `json:"-"`
	Display_name      string    `json:"display_name"`
	Profile_image_url *string   `json:"profile_image_url,omitempty" gorm:"column:profile_image_url"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}
