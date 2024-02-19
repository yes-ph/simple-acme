package simple_acme

import (
	"crypto"

	"github.com/go-acme/lego/v4/registration"
)

type user struct {
	email        string
	registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *user) GetEmail() string {
	return u.email
}
func (u user) GetRegistration() *registration.Resource {
	return u.registration
}
func (u *user) GetPrivateKey() crypto.PrivateKey {
	return u.key
}
