package character_test

import (
	"context"
	pb_character "genshin-grpc/proto/character"
	pb_common "genshin-grpc/proto/common"
	"genshin-grpc/services/character"
	test_utils "genshin-grpc/tests"
	"genshin-grpc/utils"
	"testing"
)

func TestGetCharacter(t *testing.T) {
	utils.LoadEnvVars("../../.env")
	conn := utils.ConnectToDb()

	s := character.Server{DB: conn}
	req := &pb_character.GetCharacterRequest{Id: "1"}
	ctx := context.Background()

	// previously passing db connection via context
	// ctx = context.WithValue(ctx, keys.DBSession, conn)

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

func TestStreamData(t *testing.T) {
	utils.LoadEnvVars("../../.env")
	conn := utils.ConnectToDb()
	ctx := context.Background()

	// previously passing db connection via context
	// ctx = context.WithValue(ctx, keys.DBSession, conn)

	// create mock stream to be used by server
	mockStream := test_utils.NewMockStream(ctx)

	s := character.Server{DB: conn}

	// establish the stream within a goroutine to not block
	// the stream has to listen on a goroutine or else
	// the below code can't be executed
	go func() {
		s.StreamData(mockStream)
	}()

	// send a request to the stream
	mockStream.RecvChan <- &pb_character.StreamRequest{Message: "hello", CharacterId: "1"}
	close(mockStream.RecvChan)

	// Check the response from the server
	resp := <-mockStream.SendChan

	actual := resp.Message
	expected := "Amber says: hello"
	test_utils.Compare(t, expected, actual)
}
