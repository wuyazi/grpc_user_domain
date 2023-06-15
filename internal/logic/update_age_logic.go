package logic

import (
	"context"

	"github.com/wuyazi/grpc_user_domain/internal/svc"
	"github.com/wuyazi/grpc_user_domain/user_domain"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAgeLogic {
	return &UpdateAgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAgeLogic) UpdateAge(in *user_domain.UpdateAgeReq) (*user_domain.UserIdResp, error) {
	// todo: add your logic here and delete this line

	return &user_domain.UserIdResp{}, nil
}
