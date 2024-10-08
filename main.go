package main

import (
	"log"
	"net"

	character_pb "genshin-grpc/proto/character"
	character_service "genshin-grpc/services/character"
	"genshin-grpc/utils"

	constellation_pb "genshin-grpc/proto/constellation"
	constellation_service "genshin-grpc/services/constellation"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	utils.LoadEnvVars("")
	conn := utils.ConnectToDb()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50052: %v", err)
	}

	s := grpc.NewServer(
	// previously passing db connection via context
	// instead of injecting it into service servers

	// grpc.ChainStreamInterceptor(
	// 	middleware.DBStreamServerInterceptor(conn),
	// ),
	// grpc.ChainUnaryInterceptor(
	// 	middleware.DBUnaryServerInterceptor(conn),
	// ),
	)

	character_pb.RegisterCharacterServiceServer(s, &character_service.Server{DB: conn})
	constellation_pb.RegisterConstellationServiceServer(s, &constellation_service.Server{DB: conn})

	reflection.Register(s)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
