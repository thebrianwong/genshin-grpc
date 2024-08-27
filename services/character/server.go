package character

import (
	pb "genshin-grpc/proto/character"

	"github.com/jackc/pgx/v5"
)

type Server struct {
	pb.UnimplementedCharacterServiceServer
	// pass db connection to server instead of context
	DB *pgx.Conn
}
