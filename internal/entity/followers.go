package entity

type Followers struct {
	id         int    `json:"id"`
	email      string `json:"email"`
	name       string `json:"name"`
	secondName string `json:"second_name"`
	//birthDay data   тип данных дата
}
