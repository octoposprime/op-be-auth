package application

import (
	"context"

	me "github.com/octoposprime/op-be-auth/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-auth/internal/domain/model/object"
)

// UserServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the user service.
type UserServicePort interface {
	// CheckUserPassword checks the given user's password in the user service.
	CheckUserPassword(ctx context.Context, loginRequest *mo.LoginRequest) (*me.User, error)
}
