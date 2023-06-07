package archive

import "context"

type ProfileServices interface {
	GetProfile(context.Context, GetProfileRequest) (GetProfileReply, error)
}
