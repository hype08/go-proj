package domains

import (
	"github.com/hype08/go-proj/internal/architecture/domains/users"
	"github.com/hype08/go-proj/internal/architecture/repositories"
	"github.com/hype08/go-proj/internal/database"
)

type Domains struct {
	user users.Domain
}

func (d *Domains) User() users.Domain {
	return d.user
}

func NewDomains(
	db *database.DB,
	txm *database.TxManager,
) *Domains {
	userRepository := repositories.NewUserRepository(db.Pool)

	return &Domains{
		user: users.New(userRepository),
	}
}
