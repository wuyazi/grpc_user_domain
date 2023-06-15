package logic

import (
	"context"

	"github.com/wuyazi/grpc_user_domain/internal/svc"
	"github.com/wuyazi/grpc_user_domain/user_domain"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGenderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGenderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGenderLogic {
	return &UpdateGenderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGenderLogic) UpdateGender(in *user_domain.UpdateGenderReq) (*user_domain.UserIdResp, error) {
	// todo: add your logic here and delete this line

	return &user_domain.UserIdResp{}, nil
}
