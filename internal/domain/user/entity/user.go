package entity

type User struct {
	ID       *ID    `gorm:"embedded;primaryKey"`
	Email    *Email `gorm:"embedded"`
	Password string `gorm:"column:password;type:varchar(255)"`
	Name     *Name  `gorm:"embedded"`
}

func New(id *ID, email *Email, name *Name, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
		Name:     name,
	}
}
