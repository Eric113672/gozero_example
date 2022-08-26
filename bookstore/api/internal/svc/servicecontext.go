package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gozero_example/bookstore/api/internal/config"
	"gozero_example/bookstore/rpc/add/adder"
	"gozero_example/bookstore/rpc/check/checker"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
