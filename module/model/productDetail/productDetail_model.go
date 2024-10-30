package productdetail

import "gorm.io/gorm"

type ProductDetail struct {
	gorm.Model
	ProductID uint
	ImageID   uint   `json:"iamge"`
	Size      string `json:"size"`
	Warna     string `json:"warna"`
	Stock     int    `json:"stock"`
}
