package logic

import (
	"context"

	"github.com/Awadabang/wasser/internal/svc"
	"github.com/Awadabang/wasser/types/worker"

	"github.com/zeromicro/go-zero/core/logx"
)

type LifeCircleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLifeCircleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LifeCircleLogic {
	return &LifeCircleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LifeCircleLogic) LifeCircle(in *worker.LifeCircleReq) (*worker.LifeCircleResp, error) {
	// todo: add your logic here and delete this line

	return &worker.LifeCircleResp{}, nil
}
