package application

import (
	"context"

	me "github.com/octoposprime/op-be-auth/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-auth/internal/domain/model/object"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tjwt "github.com/octoposprime/op-be-shared/tool/jwt"
)

// Login generates an authentication token if the given login request values are valid.
func (a *Service) Login(ctx context.Context, loginRequest mo.LoginRequest) (mo.Token, error) {
	if err := loginRequest.Validate(); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	user, err := a.CheckUserPassword(ctx, &loginRequest)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	if err := a.CheckIsAuthenticable(user); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	claims := tjwt.NewClaims(user.Id.String(), user)
	accessToken, refreshToken, err := claims.GenerateJWT()
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeINFO, "Login", userId, "login succeeded"))
	return *mo.NewToken(accessToken, refreshToken), nil

}

// Refresh regenerate an authentication token.
func (a *Service) Refresh(ctx context.Context, token mo.Token) (mo.Token, error) {
	return *mo.NewEmptyToken(), nil
}

// Logout clears some footprints for the user.
func (a *Service) Logout(ctx context.Context, token mo.Token) error {
	if err := token.Validate(); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Logout", userId, err.Error()))
		return err
	}
	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Logout", userId, "logout succeeded"))
	return nil
}
