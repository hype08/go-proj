package domains

import (
	"github.com/hype08/go-proj/internal/architecture/domains/users"
	"github.com/hype08/go-proj/internal/architecture/repositories"
	"github.com/hype08/go-proj/internal/database"
)

type Domains struct {
	Users users.Domain
}

func (d *Domains) User() users.Domain {
	return d.Users
}

func NewDomains(
	db *database.DB,
	txm *database.TxManager,
) *Domains {
	userRepository := repositories.NewUserRepository(db.Pool)

	return &Domains{
		Users: users.New(userRepository),
	}
}
