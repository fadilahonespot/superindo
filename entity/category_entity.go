package entity

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product 
}
