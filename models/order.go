package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Order struct{
	GormModel
	CustomerName	string  `gorm:"not null;" json:"customer_name" form:"customer_name" valid:"required"`
	Items	[]Items `gorm:"foreignKey:OrderID"`
}

type Items struct{
	GormModel
	ItemCode		string `gorm:"not null;" json:"itemcode"`
	Description		string `gorm:"not null;" json:"description"`
	Quantity		int `gorm:"not null;" json:"quantity"`
	OrderID			uint `gorm:"not null;" json:"orderid"`
}