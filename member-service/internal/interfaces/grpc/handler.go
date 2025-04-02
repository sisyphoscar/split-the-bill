package grpc

import (
	"context"
	"log"
	"member-service/internal/infrastructure/proto"
)

func (s *MemberServer) CreateMember(ctx context.Context, req *proto.CreateMemberRequest) (*proto.CreateMemberResponse, error) {
	log.Println("CreateMember called with request:", req)
	member, err := s.service.Create(req.Name, req.Email)
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
