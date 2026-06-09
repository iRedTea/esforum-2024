package service

import (
	"esforum"
)

type Authorization interface {
	Authorize(header string) (*esforum.AuthorizedUser, error)
	AuthorizeById(id int64) (*esforum.AuthorizedUser, error)
	AuthorizeAndUpdatePicture(header string, picUrl string) (*esforum.AuthorizedUser, error)
}

type Service struct {
	Authorization
}

func NewService() *Service {
	return &Service{
		Authorization: NewAuthService(),
	}
}
