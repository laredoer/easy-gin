package models

type Person struct{
	Id int `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName string `json:"last_name" form:"last_name"`
}

