package advert

import (
	"gorm.io/gorm"
)

type Advert struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text;not null" json:"description"`
	ImageURL    string `gorm:"type:varchar(512)" json:"image_url"`
	Price       uint64 `gorm:"not null" json:"price"`
	Author      string `gorm:"not null" json:"author"`
}
