package tests

import (
	"context"
	"fmt"
	pb_character "genshin-grpc/proto/character"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// MockStream implements grpc.BidiStreamingServer with generic parameters
type MockStream struct {
	grpc.ServerStream
	Ctx      context.Context // used for passing db connection
	RecvChan chan *pb_character.StreamRequest
	SendChan chan *pb_character.StreamResponse
}

// NewMockStream creates a new instance of MockStream
func NewMockStream(context context.Context) *MockStream {
	return &MockStream{
		Ctx:      context,
		RecvChan: make(chan *pb_character.StreamRequest, 1),
		SendChan: make(chan *pb_character.StreamResponse, 1),
	}
}

func (m *MockStream) Send(resp *pb_character.StreamResponse) error {
	m.SendChan <- resp
	return nil
}

func (m *MockStream) Recv() (*pb_character.StreamRequest, error) {
	req, ok := <-m.RecvChan
	if !ok {
		return nil, fmt.Errorf("channel closed")
	}
	return req, nil
}

// Mock context functions
func (m *MockStream) Context() context.Context {
	return m.Ctx
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
