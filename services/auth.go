package services

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
	"go-tutorial/internal/model"
	"go-tutorial/internal/utils"
)

type AuthService struct{
	db *gorm.DB
}


func InitAuthService(db *gorm.DB) *AuthService{
	db.AutoMigrate(&model.User{})
		return &AuthService{
			db:db,
		}
}

func(a *AuthService) CheckUserExistsOrNot(email *string) bool {
	var user model.User
	if err:=a.db.Where("email=?",email).Find(&user).Error; err!=nil{
		return false;
	}
	if user.Email!=""{
		return true;
	}
	return false;
}

func (a *AuthService) Login(email *string, password *string)(*model.User,error){
	if email==nil{
       return nil,errors.New("Email can't be null")
	}

	if password==nil{
     return nil,errors.New("Password can't be empty")
	}

	var user model.User

	if err:=a.db.Where("email=?",email).Find(&user).Error; err!=nil{
		return nil,err;
	}

	fmt.Println(user)
	if user.Email==""{
		return nil,errors.New("No user found.")
	}

	if utils.CheckPasswordHash(*password,user.Password) == false {
		return nil,errors.New("Password is incorrect")
	}

	return &user,nil;

}


func (a *AuthService) Register(email *string, password *string)(*model.User,error){
	if email==nil{
       return nil,errors.New("Email can't be null")
	}

	if password==nil{
     return nil,errors.New("Password can't be empty")
	}

	if a.CheckUserExistsOrNot(email) {
		return nil,errors.New("User already exists.")
	}

	hashedPwd,err:=utils.HashPassword(*password)

	if err!=nil{
		return nil,err
	}

	var user model.User

	user.Email=*email
	user.Password=hashedPwd

	if err:=a.db.Create(&user).Error; err!=nil{
		return nil,err;
	}

	return &user,nil;

}