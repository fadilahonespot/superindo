package entity

type Category struct {
	ID      int `gorm:"primarykey"`
	Name    string
	Product []Product
}
