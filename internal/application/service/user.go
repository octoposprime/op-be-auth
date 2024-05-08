package application

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-auth/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-auth/internal/domain/model/object"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// GetUsersByFilter returns the users that match the given filter.
func (a *Service) GetUsersByFilter(ctx context.Context, userFilter me.UserFilter) (me.Users, error) {
	return a.DbPort.GetUsersByFilter(ctx, userFilter)
}

// GetUserPasswordByUserId returns active password of the given user.
func (a *Service) GetUserPasswordByUserId(ctx context.Context, userId uuid.UUID) (me.UserPassword, error) {
	if userId.String() == "" || userId == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUserPasswordByUserId", userId, err.Error()))
		return me.UserPassword{}, err
	}
	// if the userPassword is cached in the redis repository return it
	userPassword, err := a.RedisPort.GetUserPasswordByUserId(ctx, userId)
	if err == nil && userPassword.UserId == userId {
		return userPassword, err
	}
	// else the userPassword is not cached in the redis repository get and return the userPassword from db
	userPassword, err = a.DbPort.GetUserPasswordByUserId(ctx, userId)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUserPasswordByUserId", userId, err.Error()))
		return userPassword, err
	}
	// and also write it to redis.
	err = a.RedisPort.ChangePassword(ctx, userPassword)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetUserPasswordByUserId", userId, err.Error()))
		return userPassword, err
	}
	return userPassword, err
}
