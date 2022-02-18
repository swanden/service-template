package entity

type Name struct {
	FirstName string `gorm:"column:first_name;type:varchar(255)"`
	LastName  string `gorm:"column:last_name;type:varchar(255)"`
}

func NewName(first, last string) *Name {
	return &Name{
		FirstName: first,
		LastName:  last,
	}
}
