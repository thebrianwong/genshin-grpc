package constellation

import (
	pb "genshin-grpc/proto/constellation"

	"github.com/jackc/pgx/v5"
)

type Server struct {
	pb.UnimplementedConstellationServiceServer
	DB *pgx.Conn
}
