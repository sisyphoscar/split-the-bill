package grpc

import (
	"context"
	"log"
	"member-service/internal/configs"
	"member-service/internal/members"
	"net"

	"google.golang.org/grpc"
)

type MemberServer struct {
	members.UnimplementedMembersServiceServer
}

func ListenAndServe() {
	lis, err := net.Listen("tcp", configs.App.Domain+":"+configs.App.GRPC_Port)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	members.RegisterMembersServiceServer(s, &MemberServer{})

	log.Printf("gRPC Server started on port %s", configs.App.GRPC_Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}

func (s *MemberServer) GetMember(ctx context.Context, req *members.GetMemberRequest) (*members.GetMemberResponse, error) {
	grpcMember := &members.Member{
		Id:    "1",
		Name:  "John Doe",
		Email: "sss@sss.com",
	}

	return &members.GetMemberResponse{
		Member: grpcMember,
	}, nil
}

func (s *MemberServer) CreateMember(ctx context.Context, req *members.CreateMemberRequest) (*members.CreateMemberResponse, error) {
	grpcMember := &members.Member{
		Id:    "1",
		Name:  "John Doe",
		Email: "sss@sss.com",
	}

	return &members.CreateMemberResponse{
		Member: grpcMember,
	}, nil
}

func (s *MemberServer) UpdateMember(ctx context.Context, req *members.UpdateMemberRequest) (*members.UpdateMemberResponse, error) {
	grpcMember := &members.Member{
		Id:    "1",
		Name:  "John Doe",
		Email: "sss@sss.com",
	}

	return &members.UpdateMemberResponse{
		Member: grpcMember,
	}, nil
}

func (s *MemberServer) DeleteMember(ctx context.Context, req *members.DeleteMemberRequest) (*members.DeleteMemberResponse, error) {
	grpcMember := &members.Member{
		Id:    "1",
		Name:  "John Doe",
		Email: "sss@sss.com",
	}

	return &members.DeleteMemberResponse{
		Member: grpcMember,
	}, nil
}
