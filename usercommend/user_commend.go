// Code generated by goctl. DO NOT EDIT.
// Source: grpc_user_domain.proto

package usercommend

import (
	"context"

	"github.com/wuyazi/grpc_user_domain/user_domain"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateReq         = user_domain.CreateReq
	UpdateAgeReq      = user_domain.UpdateAgeReq
	UpdateGenderReq   = user_domain.UpdateGenderReq
	UpdateNicknameReq = user_domain.UpdateNicknameReq
	UserIdResp        = user_domain.UserIdResp

	UserCommend interface {
		Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*UserIdResp, error)
		UpdateNickname(ctx context.Context, in *UpdateNicknameReq, opts ...grpc.CallOption) (*UserIdResp, error)
		UpdateAge(ctx context.Context, in *UpdateAgeReq, opts ...grpc.CallOption) (*UserIdResp, error)
		UpdateGender(ctx context.Context, in *UpdateGenderReq, opts ...grpc.CallOption) (*UserIdResp, error)
	}

	defaultUserCommend struct {
		cli zrpc.Client
	}
)

func NewUserCommend(cli zrpc.Client) UserCommend {
	return &defaultUserCommend{
		cli: cli,
	}
}

func (m *defaultUserCommend) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*UserIdResp, error) {
	client := user_domain.NewUserCommendClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultUserCommend) UpdateNickname(ctx context.Context, in *UpdateNicknameReq, opts ...grpc.CallOption) (*UserIdResp, error) {
	client := user_domain.NewUserCommendClient(m.cli.Conn())
	return client.UpdateNickname(ctx, in, opts...)
}

func (m *defaultUserCommend) UpdateAge(ctx context.Context, in *UpdateAgeReq, opts ...grpc.CallOption) (*UserIdResp, error) {
	client := user_domain.NewUserCommendClient(m.cli.Conn())
	return client.UpdateAge(ctx, in, opts...)
}

func (m *defaultUserCommend) UpdateGender(ctx context.Context, in *UpdateGenderReq, opts ...grpc.CallOption) (*UserIdResp, error) {
	client := user_domain.NewUserCommendClient(m.cli.Conn())
	return client.UpdateGender(ctx, in, opts...)
}
