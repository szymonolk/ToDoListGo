package models

type Task struct {
	Id        uint   `gorm:"primary_key;auto_increment" json:"id"`
	TaskName  string `gorm:"size:255;not null" json:"album_name"`
	Completed bool   `gorm:"default:false" json:"completed"`
}
