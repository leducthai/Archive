package dm

import "gorm.io/gorm"

type DataManager struct {
	db *gorm.DB
}
