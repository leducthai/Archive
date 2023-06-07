package dm

import (
	archive "github.com/leducthai/Archive"
	"github.com/leducthai/Archive/impl/services/profile"
)

func (dm *DataManager) Profile() archive.ProfileServices {
	return profile.NewService(dm.db)
}
