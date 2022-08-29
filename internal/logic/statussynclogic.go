package logic

import (
	"context"
	"io"
	"log"
	"runtime"

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

func (l *StatusSyncLogic) StatusSync(stream worker.Worker_StatusSyncServer) error {
	log.Println("StatusSync Connect Established", runtime.NumGoroutine())
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(res)
		continue
	}
	log.Println("StatusSync Release")
	return nil
}
