package models

type PlayerModel struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	Score int
}
