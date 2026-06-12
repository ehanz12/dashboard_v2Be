package responses

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	NomorHP   string    `json:"nomor_hp"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}

type UserMini struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
