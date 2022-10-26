package entity

import "time"

type Followers struct {
	id         int       `json:"id"`
	email      string    `json:"email"`
	name       string    `json:"name"`
	secondName string    `json:"second_name"`
	birthDay   time.Time `json:"birth_day"`
}
