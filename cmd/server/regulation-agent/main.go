package main

import (
	"fmt"
	"net"

	"github.com/edgeesg/edge-esg-backend/internal/loggers"
	"google.golang.org/grpc"
)

func main() {
	loggers.Init()
	lis, err := net.Listen("tcp", "0.0.0.0:50059")
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}
	s := grpc.NewServer()
	loggers.Info("Regulation Agent starting", map[string]interface{}{"port": 50059})
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to serve: %v", err))
	}
}
