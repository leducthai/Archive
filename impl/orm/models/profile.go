package models

type Profile struct {
	ID     string `gorm:"type:text;primaryKey"`
	Name   string `gorm:"type:text;not null"`
	Bio    string `gorm:"type:text"`
	IsMale bool   `gorm:"type:bool;not null"`
}
