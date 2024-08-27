package character_test

import (
	"context"
	"fmt"
	"genshin-grpc/keys"
	pb_character "genshin-grpc/proto/character"
	pb_common "genshin-grpc/proto/common"
	"genshin-grpc/services/character"
	test_utils "genshin-grpc/tests"
	"genshin-grpc/utils"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestGetCharacter(t *testing.T) {
	utils.LoadEnvVars("../../.env")
	conn := utils.ConnectToDb()

	s := character.Server{}
	req := &pb_character.GetCharacterRequest{Id: "1"}
	ctx := context.Background()
	ctx = context.WithValue(ctx, keys.DBSession, conn)
	resp, err := s.GetCharacter(ctx, req)
	if err != nil {
		t.Error("Error calling GetCharacter()")
	}

	actual := resp.Character[0]
	expected := pb_common.Character{
		Name:    "Amber",
		Gender:  "Female",
		Height:  "Medium",
		Element: pb_common.Element_PYRO,
	}

	test_utils.Compare(t, expected.Name, actual.Name)
	test_utils.Compare(t, expected.Gender, actual.Gender)
	test_utils.Compare(t, expected.Height, actual.Height)
	test_utils.Compare(t, expected.Element, actual.Element)
}

// MockStream implements grpc.BidiStreamingServer with generic parameters
type MockStream struct {
	grpc.ServerStream
	recvChan chan *pb_character.StreamRequest
	sendChan chan *pb_character.StreamResponse
}

// NewMockStream creates a new instance of MockStream
func NewMockStream() *MockStream {
	return &MockStream{
		recvChan: make(chan *pb_character.StreamRequest, 1),
		sendChan: make(chan *pb_character.StreamResponse, 1),
	}
}

func (m *MockStream) Send(resp *pb_character.StreamResponse) error {
	m.sendChan <- resp
	return nil
}

func (m *MockStream) Recv() (*pb_character.StreamRequest, error) {
	req, ok := <-m.recvChan
	if !ok {
		return nil, fmt.Errorf("channel closed")
	}
	return req, nil
}

// Mock context functions
func (m *MockStream) Context() context.Context {
	return context.Background()
}

func (m *MockStream) SendHeader(metadata.MD) error {
	return nil
}

func (m *MockStream) SetHeader(metadata.MD) error {
	return nil
}

func (m *MockStream) SetTrailer(metadata.MD) {}

func (m *MockStream) SendMsg(msg interface{}) error {
	return nil
}

func (m *MockStream) RecvMsg(msg interface{}) error {
	return nil
}

func TestStreamData(t *testing.T) {
	utils.LoadEnvVars("../../.env")
	conn := utils.ConnectToDb()

	// create mock stream to be used by server
	mockStream := NewMockStream()

	s := character.Server{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, keys.DBSession, conn)

	// establish the stream within a goroutine to not block
	// the stream has to listen on a goroutine or else
	// the below code can't be executed
	go func() {
		s.StreamData(mockStream)
	}()

	// send a request to the stream
	mockStream.recvChan <- &pb_character.StreamRequest{Message: "test data"}
	close(mockStream.recvChan)

	// Check the response from the server
	resp := <-mockStream.sendChan

	actual := resp.Message
	expected := "You sent a message: test data"
	test_utils.Compare(t, expected, actual)
}
