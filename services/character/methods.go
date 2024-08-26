package character

import (
	"context"

	"genshin-test/keys"
	pb_character "genshin-test/proto/character"
	pb_common "genshin-test/proto/common"

	"github.com/jackc/pgx/v5"
)

func (*Server) GetCharacter(ctx context.Context, in *pb_character.GetCharacterRequest) (*pb_character.GetCharacterResponse, error) {
	id := in.Id

	// db connection established during server initialization
	// and passed to the gRPC server to be used throughout the app
	// via context
	db := (ctx.Value(keys.DBSession)).(*pgx.Conn)

	// Scan into these variables because there is this error
	// if you try to scan directly into the proto Character struct
	// can't scan into dest[0]: cannot scan varchar (OID 1043) in text format into string
	var name string
	var gender string
	var height string
	var element string

	err := (*db).QueryRow(
		context.Background(),
		"SELECT character.name, gender, height, element.name FROM genshin.character INNER JOIN genshin.element ON character.element_id=element.id WHERE character.id=$1",
		&id,
	).Scan(&name, &gender, &height, &element) //works

	if err != nil {
		return nil, err
	}

	character := pb_common.Character{
		Name:    name,
		Gender:  gender,
		Height:  height,
		Element: element,
	}

	// return as array even if only one value
	// so that client can always expect to work with an array
	response := []*pb_common.Character{&character}

	return &pb_character.GetCharacterResponse{Character: response}, nil
}
