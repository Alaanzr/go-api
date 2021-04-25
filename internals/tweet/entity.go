package tweet

import "time"

type Tweet struct {
	Id        int       `json:"id"`
	Content   string    `json:"content" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
