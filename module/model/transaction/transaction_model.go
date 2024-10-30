package transaction

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID        uint
	PaymentMethod string
	Status        string
	Amount        float64
}
