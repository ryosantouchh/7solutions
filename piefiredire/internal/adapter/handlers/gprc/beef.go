package gprc

import (
	"context"

	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/ports"
)

type BeefgRPCServer struct {
	store ports.BeefService
}

func NewBeefgRPCServer(service ports.BeefService) *BeefgRPCServer {
	return &BeefgRPCServer{store: service}
}

func (server *BeefgRPCServer) GetSummary(ctx context.Context, empty *message.Empty) {
	beefString, err := server.store.Get()
	if err != nil {
		return nil
	}
}
