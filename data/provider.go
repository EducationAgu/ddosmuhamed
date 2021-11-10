package data

import (
	"backend/internal/interfaces"

	"github.com/go-pg/pg/v10"
)

var _ interfaces.Provider = (*Provider)(nil)

type Provider struct {
	user    interfaces.UserProvider
	session interfaces.SessionProvider
	account interfaces.AccountProvider
}

func New(db *pg.DB) *Provider {
	return &Provider{
		user:    NewUser(db),
		session: NewSession(db),
		account: NewAccount(db),
	}
}

func (p *Provider) User() interfaces.UserProvider {
	return p.user
}

func (p *Provider) Session() interfaces.SessionProvider {
	return p.session
}

func (p *Provider) Account() interfaces.AccountProvider {
	return p.account
}
