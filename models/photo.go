package models

type Photo struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(300)" json:"title"`
	Caption string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(300)" json:"photoUrl"`
	UserId int64 `gorm:"type:int" json:"userId"`
}