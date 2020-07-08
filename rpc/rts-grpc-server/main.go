package main

import (
	fmt "fmt"
	"log"
	"net"

	iot_realtime "./iot_realtime"
	grpc "google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5000))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	iot_realtime.RegisterRealTimeServiceServer(grpcServer, &RealTimeServiceServer{})

	log.Println("GRPC Server is running on :5000")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
