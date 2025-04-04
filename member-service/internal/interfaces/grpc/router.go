package grpc

import (
	"log"
	"member-service/internal/configs"
	"member-service/internal/domain/member"
	"member-service/internal/infrastructure/proto"
	"net"

	"google.golang.org/grpc"
)

type MemberServer struct {
	proto.UnimplementedMembersServiceServer
	service *member.MemberService
}

func Listen(service *member.MemberService) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+configs.App.GRPC_Port)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterMembersServiceServer(s, &MemberServer{service: service})

	log.Printf("gRPC Server started on port %s", configs.App.GRPC_Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
