package tweet

import "time"

type Tweet struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	Author    int       `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
