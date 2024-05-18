package beef_grpc

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BeefgRPCServer struct {
	store ports.BeefService

	// embed for implement the BeefServiceServer
	UnimplementedBeefServiceServer
}

func NewBeefgRPCServer(service ports.BeefService) *BeefgRPCServer {
	return &BeefgRPCServer{store: service}
}

// func (s *BeefgRPCServer) mustEmbedUnimplementedBeefServiceServer() {}

func (s *BeefgRPCServer) GetSummary(ctx context.Context, empty *Empty) (*BeefSummary, error) {
	beefString, err := s.store.Get()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error fetch bacon ipsum: %v \n", err))
	}

	regX := regexp.MustCompile(`[a-zA-Z0-9_]+`)
	matchWords := regX.FindAllString(beefString, -1)

	var beefCount map[string]int32 = make(map[string]int32)
	for i := range matchWords {
		word := strings.ToLower(matchWords[i])
		if _, ok := beefCount[word]; ok {
			beefCount[word] += 1
		} else {
			beefCount[word] = 1
		}
	}

	response := &BeefSummary{Beef: beefCount}
	return response, nil
}
