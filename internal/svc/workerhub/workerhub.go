package workerhub

import "github.com/Awadabang/wasser/types/worker"

type Manager interface {
	Register(id string, lcStream worker.Worker_LifeCircleServer, ssStream worker.Worker_StatusSyncServer)
	Send()
}

type workerItem struct {
	lifeCircleServer worker.Worker_LifeCircleServer
	statusSyncServer worker.Worker_StatusSyncServer
}

type hub struct {
	workers map[string]*workerItem
}

func NewManager() Manager {
	return &hub{
		workers: make(map[string]*workerItem),
	}
}

func (h *hub) Register(id string, lcStream worker.Worker_LifeCircleServer, ssStream worker.Worker_StatusSyncServer) {
	h.workers[id] = &workerItem{
		lifeCircleServer: lcStream,
		statusSyncServer: ssStream,
	}
}

func (h *hub) Send() {
	for _, server := range h.workers {
		server.lifeCircleServer.Send(&worker.LifeCircleResp{
			Status: "广播消息",
			Alive:  true,
		})
	}
}
