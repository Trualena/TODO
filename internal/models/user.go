package models

import "time"

type Users struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"` //это поле не будет в джсон
	Created_at time.Time `json:"created_at"`
}
