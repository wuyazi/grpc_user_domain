package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"github.com/wuyazi/gddd"
	"github.com/wuyazi/grpc_user_domain/internal/aggregate"

	"github.com/wuyazi/grpc_user_domain/internal/svc"
	"github.com/wuyazi/grpc_user_domain/user_domain"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user_domain.CreateReq) (*user_domain.UserIdResp, error) {
	meta := dtmgimp.GetMetaFromContext(l.ctx, "test_header")
	fmt.Println("=== === ===", meta)
	ctx0 := gddd.NewContext(l.ctx, l.svcCtx.Repository)
	// create user
	userAgg := &aggregate.UserAggregate{}
	userAgg.InitId()
	err := userAgg.Create(ctx0, in.Nickname)
	if err != nil {
		return nil, err
	}
	ok, saveErr := ctx0.Save(userAgg)
	if saveErr != nil {
		err = saveErr
		return nil, saveErr
	}
	if !ok {
		err = fmt.Errorf("create user failed")
		return nil, err
	}
	return &user_domain.UserIdResp{UserId: userAgg.Id}, nil
}
