package infrastructure

import (
	"context"
	"fmt"

	me "github.com/octoposprime/op-be-auth/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-auth/internal/domain/model/object"
	mapper "github.com/octoposprime/op-be-auth/pkg/infrastructure/mapper/service"
	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CheckUserPassword sends the given user to the user service and returns the result of the password check.
func (a ServiceAdapter) CheckUserPassword(ctx context.Context, loginRequest *mo.LoginRequest) (*me.User, error) {
	conn, err := grpc.Dial(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println(tconfig.GetInternalConfigInstance().Grpc.UserHost+":"+tconfig.GetInternalConfigInstance().Grpc.UserPort, "tconfig.GetInternalConfigInstance().Grpc.UserHost+\":\"+tconfig.GetInternalConfigInstance().Grpc.UserPort")
	if err != nil {
		fmt.Println("loginRequest.String()", loginRequest.String())
		fmt.Println(err)
		return &me.User{}, err
	}
	user, err := pb_user.NewUserSvcClient(conn).CheckUserPassword(ctx, mapper.NewUserWithPasswordFromLoginRequest(*loginRequest).ToPb())
	if err != nil {
		fmt.Println("NewUserSvcClient(conn).CheckUserPassword(ctx, userWithPassword)", user, err)
		fmt.Println(loginRequest.String())
		fmt.Println(err)
		return &me.User{}, err
	}
	return mapper.NewUserFromPb(user).ToEntity(), nil
}
