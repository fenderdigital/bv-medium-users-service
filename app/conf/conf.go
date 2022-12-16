package conf

import "github.com/fenderdigital/bv-medium-users-service/internal/features/createuser"

func NewCreateUserFeature() (*createuser.CreateUser, *createuser.CreateUser) {
	return &createuser.CreateUser{}, nil
}
