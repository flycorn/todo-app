package model

import (
	"todo-app/config"
	"todo-app/helper"
)

type Todo struct{
	Id  int  `json:"id"`
	Uid int `json:"uid"`
	Title string `json:"title"`
	Content string `json:"content"`
	Status int `json:"status"`
	CreateTime helper.JSONTime `json:"create_time" gorm:"-"`
	UpdateTime helper.JSONTime `json:"update_time" gorm:"-"`
}

func (Todo) TableName() string{
	return config.Conf.Db.TablePrefix+"todo"
}

//创建时间格式化
//func (m *Todo) DateCreateTime() string{
	//createTime, _ := strconv.ParseInt(m.CreateTime, 10, 64)
	//cur := time.Unix(createTime, 0)
	//return cur.Format("2006-01-02 03:04:05")
//}

//更新时间格式化
//func (m *Todo) DateUpdateTime() string{
	//updateTime, _ := strconv.ParseInt(m.UpdateTime, 10, 64)
	//cur := time.Unix(updateTime, 0)
	//return cur.Format("2006-01-02 03:04:05")
//}


