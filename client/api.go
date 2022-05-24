package client

import (
	"github.com/go-resty/resty/v2"

	"github.com/ryanc414/magic-admin-go"
	"github.com/ryanc414/magic-admin-go/user"
)

type API struct {
	User magic.User
}

func New(secret string, client *resty.Client) *API {
	return &API{
		User: user.NewUserClient(secret, client),
	}
}
