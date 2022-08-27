package logic

import (
	"context"

	"github.com/Awadabang/wasser/internal/svc"
	"github.com/Awadabang/wasser/types/worker"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatusSyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStatusSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatusSyncLogic {
	return &StatusSyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StatusSyncLogic) StatusSync(in *worker.StatusSyncReq) (*worker.StatusSyncResp, error) {
	// todo: add your logic here and delete this line

	return &worker.StatusSyncResp{}, nil
}
