package profile

import (
	"context"

	archive "github.com/leducthai/Archive"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) archive.ProfileServices {
	return service{db: db}
}

func (s service) GetProfile(ctx context.Context, req archive.GetProfileRequest) (archive.GetProfileReply, error) {
	return archive.GetProfileReply{}, nil
}
