package models

type Player struct {
	Id    uint   `gorm:"primarykey;AUTO_INCREMENT;not null"`
	Name  string `gorm:"type:varchar(75);unique_index;not null"`
	Score int    `gorm:"not null"`
}
