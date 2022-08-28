package svc

import (
	"github.com/Awadabang/wasser/internal/config"
	"github.com/Awadabang/wasser/internal/svc/workerhub"
)

type ServiceContext struct {
	Config       config.Config
	WorkerHubMgr workerhub.Manager
}

func NewServiceContext(c config.Config) *ServiceContext {
	workerHubMgr := workerhub.NewManager()

	return &ServiceContext{
		Config:       c,
		WorkerHubMgr: workerHubMgr,
	}
}
