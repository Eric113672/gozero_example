package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozero_example/bookstore/rpc/add/internal/config"
	"gozero_example/bookstore/rpc/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
