package application

import (
	"context"

	me "github.com/octoposprime/op-be-auth/internal/domain/model/entity"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type UserQueryPort interface {
	// GetUsersByFilter returns the users that match the given filter.
	GetUsersByFilter(ctx context.Context, userFilter me.UserFilter) (me.Users, error)
}
