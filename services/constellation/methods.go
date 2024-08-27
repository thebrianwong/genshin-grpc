package constellation

import (
	"context"

	pb_common "genshin-grpc/proto/common"
	pb_constellation "genshin-grpc/proto/constellation"
	"genshin-grpc/utils"
)

func (s *Server) GetConstellation(ctx context.Context, in *pb_constellation.GetConstellationRequest) (*pb_constellation.GetConstellationResponse, error) {
	id := in.Id

	// db := (ctx.Value(keys.DBSession)).(*pgx.Conn)
	db := s.DB

	var constellationName string
	var level uint32
	var characterName string
	var gender string
	var height string
	var element string

	err := (*db).QueryRow(
		context.Background(),
		`
			SELECT constellation.name, level, character.name, gender, height, element.name FROM genshin.constellation 
			INNER JOIN genshin.character ON constellation.character_id=character.id 
			INNER JOIN genshin.element ON character.element_id=element.id 
			WHERE constellation.id=$1
		`,
		&id,
	).Scan(&constellationName, &level, &characterName, &gender, &height, &element)

	if err != nil {
		return nil, err
	}

	character := pb_common.Character{
		Name:    characterName,
		Gender:  gender,
		Height:  height,
		Element: utils.ElementStringToEnum(element),
	}

	constellation := pb_common.Constellation{
		Name:      constellationName,
		Level:     level,
		Character: &character,
	}

	response := []*pb_common.Constellation{&constellation}

	return &pb_constellation.GetConstellationResponse{Constellation: response}, nil
}
