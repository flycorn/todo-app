package model

import (
	"todo-app/helper"
	"todo-app/config"
	"todo-app/service"
	"github.com/pkg/errors"
)

type User struct{
	Id  int  `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Salt string `json:"salt"`
	CreateTime string `json:"create_time" gorm:"-"`
	UpdateTime string `json:"update_time" gorm:"-"`
}

func (User) TableName() string{
	return config.Conf.Db.TablePrefix+"user"
}

//检测密码
func (u User) CheckPassword(password string) bool{
	if helper.Md5String(password+u.Salt) != u.Password{
		return false
	}
	return true
}

//生成密码
func (u *User) CreatePassword(pwd string){
	salt := helper.RandStringBytes(6)
	password := helper.Md5String(pwd+salt)
	u.Password = password
	u.Salt = salt
}

//创建用户
func (user * User)CreateUser()(int, error){
	tx := service.Db.Begin()
	defer func(){
		if r := recover(); r != nil{
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil{
		return 0, err
	}

	//检测该用户是否已存在
	userModel := User{}
	rs := tx.Where("nickname = ?", user.Nickname).First(&userModel).RecordNotFound()
	if !rs {
		//已存在该用户
		return 0, errors.New("该昵称已被注册")
	}
	//写入数据
	if err := tx.Create(&user).Error; err != nil{
		//回滚
		tx.Rollback()
		return 0, err
	}

	//提交
	return user.Id, tx.Commit().Error
}

//查找用户
func (user *User) FindUser(nickname string) bool{
	rs := service.Db.Where("nickname = ?", nickname).First(user).RecordNotFound()
	if rs {
		return false
	}
	return true
}