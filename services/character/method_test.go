package character_test

import (
	"context"
	"genshin-grpc/keys"
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
