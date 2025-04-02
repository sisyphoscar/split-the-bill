package main

import (
	"member-service/internal/configs"
	"member-service/internal/domain/member"
	"member-service/internal/infrastructure/database"
	"member-service/internal/infrastructure/repositories"
	"member-service/internal/interfaces/grpc"
	"member-service/internal/interfaces/http"
)

func main() {
	configs.LoadConfig()

	db := database.NewDB()
	defer database.CloseDB(db)

	repo := repositories.NewMemberRepository(db)
	service := member.NewMemberService(repo)

	go grpc.Listen(service)

	http.Listen()
}
