package models

import "time"

type Tasks struct {
	ID          int64      `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"` //мб нулевое
	Done        bool       `json:"done"`
	Created_at  time.Time  `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at"`
}
