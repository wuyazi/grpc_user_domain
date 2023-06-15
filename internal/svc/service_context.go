package svc

import (
	"context"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/wuyazi/gddd"
	"github.com/wuyazi/grpc_user_domain/internal/aggregate"
	"github.com/wuyazi/grpc_user_domain/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config     config.Config
	Repository *gddd.Repository
}

func NewServiceContext(c config.Config) *ServiceContext {
	repository, err := gddd.NewRepository(context.TODO(), &gddd.RepositoryConfig{
		DomainName:          "apple",
		SubDomainName:       "user",
		EventBusNameServers: c.DtmServers,
		DtmDBConf: dtmimp.DBConf{
			Driver:   c.DataBase.Driver,
			Host:     c.DataBase.Host,
			Port:     c.DataBase.Port,
			User:     c.DataBase.User,
			Password: c.DataBase.Password,
			Db:       c.DataBase.Db,
			Schema:   c.DataBase.Schema,
		},
	})
	logx.Must(err)
	err = repository.SetSaveListener(context.TODO(), &aggregate.UserRepositorySaveListener{})
	logx.Must(err)
	err = repository.RegisterAggregates(context.TODO(), &aggregate.UserAggregate{})
	logx.Must(err)
	logx.DisableStat()

	return &ServiceContext{
		Config:     c,
		Repository: repository,
	}
}
