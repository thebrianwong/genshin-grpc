package character

import (
	"context"
	"log"

	pb_character "genshin-grpc/proto/character"
	pb_common "genshin-grpc/proto/common"
	"genshin-grpc/utils"
)

func (s *Server) StreamData(stream pb_character.CharacterService_StreamDataServer) error {
	// test on how you would call the database using context
	// context can be extracted from stream since it isn't passed as an argument for stream methods

	// ctx := stream.Context()
	// db := (ctx.Value(keys.DBSession)).(*pgx.Conn)
	db := s.DB
	// // var name string
	// (*db).QueryRow(
	// 	context.Background(),
	// 	`
	// 		SELECT name FROM genshin.character
	// 		WHERE character.id=1
	// 	`,
	// ).Scan(&name)

	// create a channel to take a request, do something with it, then respond
	msgChannel := make(chan *pb_character.StreamResponse)

	// create a channel to handle goroutine errors
	errChannel := make(chan error)

	// listen for a request from the client
	// use a goroutine to not block the main goroutine
	go func() {
		for {
			req, err := stream.Recv()
			if err != nil {
				// Handle any error that occurred during receiving request
				errChannel <- err
				return
			}

			var name string
			var gender string
			var height string
			var element string

			err = (*db).QueryRow(
				context.Background(),
				`
					SELECT character.name, gender, height, element.name FROM genshin.character
					INNER JOIN genshin.element ON character.element_id=element.id
					WHERE character.id=$1
				`,
				req.CharacterId,
			).Scan(&name, &gender, &height, &element)
			if err != nil {
				errChannel <- err
				return
			}

			character := pb_common.Character{
				Name:    name,
				Gender:  gender,
				Height:  height,
				Element: utils.ElementStringToEnum(element),
			}

			response := []*pb_common.Character{&character}

			msgChannel <- &pb_character.StreamResponse{Message: character.Name + " says: " + req.Message, Character: response}
		}
	}()

	// test for sending response on regular interval
	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Second)
	// 		msgChannel <- &pb_character.StreamResponse{Message: "The current datetime is: " + time.Now().Format("January 2, 2006 | 3:04:05PM")}
	// 	}
	// }()

	for {
		select {
		case <-stream.Context().Done():
			// Handle client disconnection or cancellation
			log.Println("Client disconnected")
			return nil
		case update := <-msgChannel:
			// handle error from sending stream response
			if err := stream.Send(update); err != nil {
				return err
			}
			log.Printf("Sent update: %s", update.Message)
		case err := <-errChannel:
			// handle error from receiving stream request
			if err != nil {
				return err
			}
		}
	}
}

func (s *Server) GetCharacter(ctx context.Context, in *pb_character.GetCharacterRequest) (*pb_character.GetCharacterResponse, error) {
	id := in.Id

	// db connection established during server initialization
	// and passed to the gRPC server to be used throughout the app

	// via context
	// db := (ctx.Value(keys.DBSession)).(*pgx.Conn)

	// via dependency injection into server
	db := s.DB

	// Scan into these variables because there is this error
	// if you try to scan directly into the proto Character struct
	// can't scan into dest[0]: cannot scan varchar (OID 1043) in text format into string
	var name string
	var gender string
	var height string
	var element string

	err := (*db).QueryRow(
		context.Background(),
		`
			SELECT character.name, gender, height, element.name FROM genshin.character
			INNER JOIN genshin.element ON character.element_id=element.id
			WHERE character.id=$1
		`,
		&id,
	).Scan(&name, &gender, &height, &element) //works

	if err != nil {
		return nil, err
	}

	character := pb_common.Character{
		Name:    name,
		Gender:  gender,
		Height:  height,
		Element: utils.ElementStringToEnum(element),
	}

	// return as array even if only one value
	// so that client can always expect to work with an array
	response := []*pb_common.Character{&character}

	return &pb_character.GetCharacterResponse{Character: response}, nil
}
