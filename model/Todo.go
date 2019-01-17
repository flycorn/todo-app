package model

import "todo-app/config"

type Todo struct{
	Id  int  `json:"id"`
	Uid int `json:"uid"`
	Title string `json:"title"`
	Content string `json:"content"`
	Status int `json:"status"`
	CreateTime string `json:"create_time" gorm:"-"`
	UpdateTime string `json:"update_time" gorm:"-"`
}

func (Todo) TableName() string{
	return config.Conf.Db.TablePrefix+"todo"
}


