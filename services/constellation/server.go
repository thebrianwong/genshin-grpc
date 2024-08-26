package constellation

import pb "genshin-test/proto/constellation"

type Server struct {
	pb.UnimplementedConstellationServiceServer
}
