package constellation

import pb "genshin-grpc/proto/constellation"

type Server struct {
	pb.UnimplementedConstellationServiceServer
}
