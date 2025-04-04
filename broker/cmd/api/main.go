package main

import (
	"broker/internal/configs"
	"broker/internal/handlers"
	"broker/internal/routes"
	"broker/internal/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	configs.LoadConfig()

	// 建立 gRPC 客戶端連接
	gRPCConn, err := grpc.NewClient(
		configs.MemberService.Domain+":"+configs.MemberService.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer gRPCConn.Close()

	memberService := services.NewMemberService(gRPCConn)
	memberHandler := handlers.NewMemberHandler(memberService)

	router := routes.Api(memberHandler)

	log.Println("Server is running: " + configs.App.Domain + ":" + configs.App.Port)

	router.Run(":" + configs.App.Port)
}
