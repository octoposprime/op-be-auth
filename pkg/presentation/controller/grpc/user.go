package presentation

import (
	"context"

	dto "github.com/octoposprime/op-be-auth/pkg/presentation/dto"
	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
)

// GetUsersByFilter returns the users that match the given filter.
func (a *Grpc) GetUsersByFilter(ctx context.Context, filter *pb_user.UserFilter) (*pb_user.Users, error) {
	users, err := a.queryHandler.GetUsersByFilter(ctx, *dto.NewUserFilter(filter).ToEntity())
	return dto.NewUserFromEntities(users).ToPbs(), err
}
