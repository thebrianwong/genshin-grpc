package character_test

import (
	"context"
	"genshin-grpc/keys"
	pb_character "genshin-grpc/proto/character"
	pb_common "genshin-grpc/proto/common"
	"genshin-grpc/services/character"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func TestGetCharacter(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

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

	if actual.Name != expected.Name {
		t.Error("Expected:", actual.Name, "| Got:", expected.Name)
	}
	if actual.Gender != expected.Gender {
		t.Error("Expected:", actual.Gender, "| Got:", expected.Gender)
	}
	if actual.Height != expected.Height {
		t.Error("Expected:", actual.Height, "| Got:", expected.Height)
	}
	if actual.Element != expected.Element {
		t.Error("Expected:", actual.Element, "| Got:", expected.Element)
	}
}
