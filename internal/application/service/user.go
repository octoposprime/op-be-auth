package application

import (
	"context"
	"fmt"

	me "github.com/octoposprime/op-be-auth/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-auth/internal/domain/model/object"
)

// CheckUserPassword checks the given user's password in the user service.
func (a *Service) CheckUserPassword(ctx context.Context, loginRequest mo.LoginRequest) (*me.User, error) {
	fmt.Println("CheckUserPassword request received from client internal/application/service/user.go")
	user, err := a.ServicePort.CheckUserPassword(ctx, &loginRequest)
	if err != nil {
		fmt.Println("CheckUserPassword request failed", err)
		return &me.User{}, err
	}
	return user, nil
}
