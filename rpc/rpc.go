package rpc

import (
	"context"
	pb "grpc-example/proto"
)

type Handler struct {
	pb.UnimplementedUsersServer
}


func (h *Handler) GetUsers(ctx context.Context, r *pb.EmptyReq) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{
		Users: []*pb.User{
			{
				Name: "test",
				Age:  10,
			},
		},
	}, nil
}