package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"

	pb "genshin-test/proto/character"
	character_service "genshin-test/services/character"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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

	pb.RegisterCharacterServiceServer(s, &character_service.Server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
