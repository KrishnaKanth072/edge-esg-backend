package main

import (
	"fmt"
	"net"

	"github.com/edgeesg/edge-esg-backend/internal/loggers"
	"google.golang.org/grpc"
)

func main() {
	loggers.Init()
	lis, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}
	s := grpc.NewServer()
	loggers.Info("Compliance Agent starting", map[string]interface{}{"port": 50054})
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to serve: %v", err))
	}
}
