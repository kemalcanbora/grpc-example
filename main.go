package main

import (
	"google.golang.org/grpc"
	pb "grpc-example/proto"
	"grpc-example/rpc"
	"log"
	"net"
)

func main()  {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUsersServer(s, &rpc.Handler{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}