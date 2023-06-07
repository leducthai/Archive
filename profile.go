package archive

type GetProfileReply struct {
	Name string
	Bio  string
	Sex  string
}

type GetProfileRequest struct {
	ID string
}
