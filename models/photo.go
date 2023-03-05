package models

type Photo struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(300)" json:"title"`
	Caption string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"primaryKey" json:"photoUrl"`
	UserId string `gorm:"primaryKey" json:"userId"`
}