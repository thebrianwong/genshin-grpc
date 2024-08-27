package middleware

import (
	"context"
	"genshin-grpc/keys"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

// following are copied from https://fale.io/blog/2022/03/21/inject-db-connections-in-golang-grpc-api
// they allow the db connection to be used in service methods via context

// better to pass via dependency injection into the server
// but keeping this here for future reference

func DBUnaryServerInterceptor(session *pgx.Conn) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, keys.DBSession, session), req)
	}
}

func DBStreamServerInterceptor(session *pgx.Conn) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = context.WithValue(stream.Context(), keys.DBSession, session)
		return handler(srv, wrapped)
	}
}
