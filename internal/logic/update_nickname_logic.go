package logic

import (
	"context"
	"fmt"
	"github.com/wuyazi/gddd"
	"github.com/wuyazi/grpc_user_domain/internal/aggregate"

	"github.com/wuyazi/grpc_user_domain/internal/svc"
	"github.com/wuyazi/grpc_user_domain/user_domain"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNicknameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNicknameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNicknameLogic {
	return &UpdateNicknameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNicknameLogic) UpdateNickname(in *user_domain.UpdateNicknameReq) (*user_domain.UserIdResp, error) {
	ctx0 := gddd.NewContext(l.ctx, l.svcCtx.Repository)
	// load user
	userAgg := &aggregate.UserAggregate{}
	userAgg.Id = in.UserId
	var has bool
	has, err := ctx0.Load(userAgg)
	if err != nil {
		return nil, err
	}
	if !has {
		err = fmt.Errorf("user not exist")
		return nil, err
	}
	// update user nickname
	err = userAgg.UpdateNickname(ctx0, in.Nickname)
	if err != nil {
		return nil, err
	}
	ok, saveErr := ctx0.Save(userAgg)
	if saveErr != nil {
		return nil, saveErr
	}
	if !ok {
		err = fmt.Errorf("update user nickname failed")
		return nil, err
	}
	return &user_domain.UserIdResp{UserId: userAgg.Id}, nil
}
