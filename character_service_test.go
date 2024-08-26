package main

import (
	"context"
	"genshin-grpc/middleware"
	character_pb "genshin-grpc/proto/character"
	character_service "genshin-grpc/services/character"
	"genshin-grpc/tests"
	"genshin-grpc/utils"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	utils.LoadEnvVars("")
	conn := utils.ConnectToDb()

	lis = bufconn.Listen(bufSize)

	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			middleware.DBStreamServerInterceptor(conn),
		),
		grpc.ChainUnaryInterceptor(
			middleware.DBUnaryServerInterceptor(conn),
		),
	)

	character_pb.RegisterCharacterServiceServer(s, &character_service.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCharacterService(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.NewClient("passthrough://bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := character_pb.NewCharacterServiceClient(conn)
	resp, err := client.GetCharacter(ctx, &character_pb.GetCharacterRequest{Id: "1"})
	if err != nil {
		t.Fatalf("GetCharacter failed: %v", err)
	}

	tests.Compare(t, resp.Character[0].Name, "Amber")
	tests.Compare(t, resp.Character[0].Gender, "Female")
}
