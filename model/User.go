package model

type User struct{
	Id  int  `json:"id"`
	Nickname string `json:"nickname"`
}

func (User) TableName() string{
	return "user"
}
