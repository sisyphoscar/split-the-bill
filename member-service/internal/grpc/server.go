package grpc

import (
	"context"
	"log"
	"member-service/internal/configs"
	"member-service/internal/grpc/proto"
	"member-service/internal/services"
	"net"

	"google.golang.org/grpc"
)

type MemberServer struct {
	proto.UnimplementedMembersServiceServer
	service *services.MemberService
}

func ListenAndServe(service *services.MemberService) {
	lis, err := net.Listen("tcp", configs.App.Domain+":"+configs.App.GRPC_Port)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterMembersServiceServer(s, &MemberServer{service: service})

	log.Printf("gRPC Server started on port %s", configs.App.GRPC_Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}

func (s *MemberServer) CreateMember(ctx context.Context, req *proto.CreateMemberRequest) (*proto.CreateMemberResponse, error) {
	member, err := s.service.CreateMember(req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	return &proto.CreateMemberResponse{
		Member: &proto.Member{
			Id:    uint64(member.ID),
			Name:  member.Name,
			Email: member.Email,
		},
	}, nil
}
