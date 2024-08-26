package character

import pb "genshin-grpc/proto/character"

type Server struct {
	pb.UnimplementedCharacterServiceServer
}
