package grpc

import (
	"context"
	"log"
	"member-service/internal/infrastructure/proto"
	"member-service/internal/interfaces/grpc/requests"
)

type CreateMemberRequestValidation struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func (s *MemberServer) CreateMember(ctx context.Context, req *proto.CreateMemberRequest) (*proto.CreateMemberResponse, error) {
	log.Println("CreateMember called with request:", req)

	validationReq := &requests.CreateMemberRequestValidation{
		Name:  req.Name,
		Email: req.Email,
	}
	err := validationReq.Validate()
	if err != nil {
		return nil, err
	}

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
