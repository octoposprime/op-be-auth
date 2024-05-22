package infrastructure

import (
	"fmt"

	me "github.com/octoposprime/op-be-auth/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-auth/internal/domain/model/object"
	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
)

type User struct {
	proto *pb_user.User
}

// String returns a string representation of the User.
func (u *User) String() string {
	return fmt.Sprintf("User.Username: %v, "+
		"User.Email: %v, "+
		"User.Id: %v, "+
		"User.Role: %v, "+
		"User.UserType: %v, "+
		"User.UserStatus: %v, "+
		"User.Tags: %v, "+
		"User.FirstName: %v, "+
		"User.LastName: %v, "+
		"User.CreatedAt: %v, "+
		"User.UpdatedAt: %v, "+
		"User.CreatedBy: %v, "+
		"User.UpdatedBy: %v, "+
		"User.DeletedBy: %v",
		u.proto.Username,
		u.proto.Email,
		u.proto.Id,
		u.proto.Role,
		u.proto.UserType,
		u.proto.UserStatus,
		u.proto.Tags,
		u.proto.FirstName,
		u.proto.LastName,
		u.proto.CreatedAt,
		u.proto.UpdatedAt,
		u.proto.CreatedBy,
		u.proto.UpdatedBy,
		u.proto.DeletedBy,
	)
}

// NewUserFromPb creates a new User from a protobuf representation.
func NewUserFromPb(user *pb_user.User) *User {
	return &User{
		proto: user,
	}
}

// ToPb returns a protobuf representation of the User.
func (u *User) ToPb() *pb_user.User {
	return u.proto
}

func (u *User) ToObject() *mo.User {
	return &mo.User{
		UserName:   u.proto.Username,
		Email:      u.proto.Email,
		Role:       u.proto.Role,
		UserType:   mo.UserType(u.proto.UserType),
		UserStatus: mo.UserStatus(u.proto.UserStatus),
		Tags:       u.proto.Tags,
		FirstName:  u.proto.FirstName,
		LastName:   u.proto.LastName,
	}
}

func (u *User) ToEntity() *me.User {
	return &me.User{
		Id: tuuid.FromString(u.proto.Id),
		User: mo.User{
			UserName:   u.proto.Username,
			Email:      u.proto.Email,
			Role:       u.proto.Role,
			UserType:   mo.UserType(u.proto.UserType),
			UserStatus: mo.UserStatus(u.proto.UserStatus),
			Tags:       u.proto.Tags,
			FirstName:  u.proto.FirstName,
			LastName:   u.proto.LastName,
		},
	}
}

// UserWithPassword is a struct that represents the mapper of a user with password.
type UserWithPassword struct {
	proto *pb_user.UserWithPassword
}

// NewUserWithPassword creates a new UserWithPassword object.
func NewUserWithPassword(userWithPassword *pb_user.UserWithPassword) *UserWithPassword {
	return &UserWithPassword{
		proto: userWithPassword,
	}
}

// String returns a string representation of the UserWithPassword.
func (u *UserWithPassword) String() string {
	return fmt.Sprintf("User.Username: %v, "+
		"User.Email: %v, "+
		"User.Id: %v, "+
		"User.Role: %v, "+
		"User.UserType: %v, "+
		"User.UserStatus: %v, "+
		"User.Tags: %v, "+
		"User.FirstName: %v, "+
		"User.LastName: %v, "+
		"User.CreatedAt: %v, "+
		"User.UpdatedAt: %v, "+
		"User.CreatedBy: %v, "+
		"User.UpdatedBy: %v, "+
		"User.DeletedBy: %v, "+
		"UserPassword.Id: %v, "+
		"UserPassword.UserId: %v, "+
		"UserPassword.Password: ***, "+
		"UserPassword.PasswordStatus: %v",
		u.proto.User.Username,
		u.proto.User.Email,
		u.proto.User.Id,
		u.proto.User.Role,
		u.proto.User.UserType,
		u.proto.User.UserStatus,
		u.proto.User.Tags,
		u.proto.User.FirstName,
		u.proto.User.LastName,
		u.proto.User.CreatedAt,
		u.proto.User.UpdatedAt,
		u.proto.User.CreatedBy,
		u.proto.User.UpdatedBy,
		u.proto.User.DeletedBy,
		u.proto.UserPassword.Id,
		u.proto.UserPassword.UserId,
		u.proto.UserPassword.PasswordStatus,
	)
}

// NewUserWithPasswordFromLoginRequest creates a new UserWithPassword from a LoginRequest.
func NewUserWithPasswordFromLoginRequest(loginRequest mo.LoginRequest) *UserWithPassword {
	return &UserWithPassword{
		proto: &pb_user.UserWithPassword{
			User: &pb_user.User{
				Username: loginRequest.UserName,
				Email:    loginRequest.Email,
			},
			UserPassword: &pb_user.UserPassword{
				Password: loginRequest.Password,
			},
		},
	}
}

// ToPb returns a protobuf representation of the UserWithPassword.
func (u *UserWithPassword) ToPb() *pb_user.UserWithPassword {
	return u.proto
}

// UserWithPasswordToLoginRequest converts the UserWithPassword to a LoginRequest.
func (u *UserWithPassword) UserWithPasswordToLoginRequest() mo.LoginRequest {
	return mo.LoginRequest{
		UserName: u.proto.User.Username,
		Email:    u.proto.User.Email,
		Password: u.proto.UserPassword.Password,
	}
}

// LoginRequestToUserWithPassword converts the LoginRequest to a UserWithPassword.
func LoginRequestToUserWithPassword(loginRequest mo.LoginRequest) *pb_user.UserWithPassword {
	return &pb_user.UserWithPassword{
		User: &pb_user.User{
			Username: loginRequest.UserName,
			Email:    loginRequest.Email,
		},
		UserPassword: &pb_user.UserPassword{
			Password: loginRequest.Password,
		},
	}
}

// type UserWithPasswords []*UserWithPassword

// // NewUserWithPasswordFromLoginRequests creates a new []*UserWithPassword from login requests.
// func NewUserWithPasswordFromLoginRequests(loginRequests []mo.LoginRequest) UserWithPasswords {
// 	userWithPasswords := make([]*UserWithPassword, len(loginRequests))
// 	for i, loginRequest := range loginRequests {
// 		userWithPasswords[i] = NewUserWithPasswordFromLoginRequest(loginRequest)
// 	}
// 	return userWithPasswords
// }
