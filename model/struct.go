package model

import (
	"fmt"
	"webserver/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type M map[string]interface{}

type Item struct {
	Name string `json:"name" form:"name"`
	EmployeeId int
	Employee *Employee
}

type Employee struct {
	ID int `json:"id" form:"id" swagger:"description(id)"`
	Name string `json:"name" form:"name" swagger:"description(name)" valid:"required"`
	Email string `json:"email" form:"email" swagger:"description(email)" valid:"required"`
	Password string `json:"password" form:"password" swagger:"description(password)" valid:"required"`
	Age int `json:"age" form:"age" swagger:"description(age)" valid:"required"`
	Division string `json:"division" form:"division" swagger:"description(division)" valid:"required"`
	Item []Item
}

type DeleteResponse struct {
	Status int
	Message string
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(e)
	
	if errCreate != nil {
		fmt.Println(errCreate)
		err = errCreate
		return err
	}

	e.Password = helpers.HassPass(e.Password)
	err = nil
	return
}
