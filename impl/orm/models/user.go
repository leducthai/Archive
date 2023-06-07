package models

type User struct {
	ID       string        `gorm:"type:text;primaryKey"`
	Email    Email         `gorm:"type:text;unique"`
	Password EncryptedData `gorm:"type:bytea;not null"`
	Profile  string        `gorm:"type:text;not null"`
}
