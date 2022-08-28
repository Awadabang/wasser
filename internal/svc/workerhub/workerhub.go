package workerhub

import "github.com/Awadabang/wasser/types/worker"

type Manager interface {
	LCRegister(worker.Worker_LifeCircleServer)
	SSRegister(worker.Worker_StatusSyncServer)
	Send()
}

type hub struct {
	lifeCircleServers []worker.Worker_LifeCircleServer
	statusSyncServers []worker.Worker_StatusSyncServer
}

func NewManager() Manager {
	return &hub{
		lifeCircleServers: make([]worker.Worker_LifeCircleServer, 0),
		statusSyncServers: make([]worker.Worker_StatusSyncServer, 0),
	}
}

func (h *hub) LCRegister(stream worker.Worker_LifeCircleServer) {
	h.lifeCircleServers = append(h.lifeCircleServers, stream)
}

func (h *hub) SSRegister(stream worker.Worker_StatusSyncServer) {
	h.statusSyncServers = append(h.statusSyncServers, stream)
}

func (h *hub) Send() {
	for _, server := range h.lifeCircleServers {
		server.Send(&worker.LifeCircleResp{
			Status: "广播消息",
			Alive:  true,
		})
	}
}
