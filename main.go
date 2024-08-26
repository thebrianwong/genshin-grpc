package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"

	character_pb "genshin-grpc/proto/character"
	character_service "genshin-grpc/services/character"

	constellation_pb "genshin-grpc/proto/constellation"
	constellation_service "genshin-grpc/services/constellation"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50052: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			DBStreamServerInterceptor(conn),
		),
		grpc.ChainUnaryInterceptor(
			DBUnaryServerInterceptor(conn),
		),
	)

	character_pb.RegisterCharacterServiceServer(s, &character_service.Server{})
	constellation_pb.RegisterConstellationServiceServer(s, &constellation_service.Server{})

	reflection.Register(s)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
