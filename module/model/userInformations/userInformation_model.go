package userinformations

import "gorm.io/gorm"

type UserInformation struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string
	Telephone string
	Email     string
}
