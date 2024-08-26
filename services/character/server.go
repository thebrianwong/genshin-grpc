package character

import pb "genshin-test/proto/character"

type Server struct {
	pb.UnimplementedCharacterServiceServer
}
