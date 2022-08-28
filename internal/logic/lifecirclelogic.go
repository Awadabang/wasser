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

func (l *LifeCircleLogic) LifeCircle(stream worker.Worker_LifeCircleServer) error {
	log.Println("LifeCircle Connect Established", runtime.NumGoroutine())
	l.svcCtx.WorkerHubMgr.LCRegister(stream)
	n := 0
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
		if res.Register {
			stream.Send(&worker.LifeCircleResp{
				Status: "OK",
				Alive:  true,
			})
		}
		if res.Alivecheck {
			if n == 2 {
				stream.Send(&worker.LifeCircleResp{
					Status: "OK",
					Alive:  false,
				})
			} else {
				stream.Send(&worker.LifeCircleResp{
					Status: "OK",
					Alive:  true,
				})
			}
			n++
		}
		continue
	}
	log.Println("LifeCircle Release")
	return nil
}
