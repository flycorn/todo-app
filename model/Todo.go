package model

import "todo-app/config"

type Todo struct{
	Id  int  `json:"id"`
	Uid int `json:"uid"`
	Content string `json:"content"`
	Status string `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (Todo) TableName() string{
	return config.Conf.Db.TablePrefix+"todo"
}


