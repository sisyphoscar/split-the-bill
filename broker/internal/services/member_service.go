package services

import (
	proto "broker/internal/proto/member"
	"context"
	"log"

	"google.golang.org/grpc"
)

type MemberService struct {
	client proto.MembersServiceClient
}

func NewMemberService(gRPCConn *grpc.ClientConn) *MemberService {
	return &MemberService{
		client: proto.NewMembersServiceClient(gRPCConn),
	}
}

func (s *MemberService) CreateMember(name string, email string) (*proto.Member, error) {
	grpcReq := &proto.CreateMemberRequest{
		Name:  name,
		Email: email,
	}

	res, err := s.client.CreateMember(context.Background(), grpcReq)
	if err != nil {
		log.Printf("Failed to create member: %v", err)
		return nil, err
	}
	return res.Member, nil
}
