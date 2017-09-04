package main

import "github.com/jinzhu/gorm"

/*
The User GORM model, which accepts the following fields: FirstName(string),
LastName(string), EmailAddress(string)
*/
type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	EmailAddress string
}
