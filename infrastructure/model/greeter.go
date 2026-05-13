package model

type Greeter struct {
	ID   int64  `gorm:"primaryKey;autoIncrement"`
	Hello string `gorm:"column:hello"`
}

func (Greeter) TableName() string { return "greeter" }
