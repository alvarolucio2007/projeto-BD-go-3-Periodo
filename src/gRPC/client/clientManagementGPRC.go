// Package grpcclient gerencia os clients do gRPC.
package grpcclient

import (
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserConexao struct {
	User proto.UsuariosServiceClient
	Conn *grpc.ClientConn
}
type ProvaConexao struct {
	Prova proto.ProvaServiceClient
	Conn  *grpc.ClientConn
}
type NotaConexao struct {
	Nota proto.NotaServiceClient
	Conn *grpc.ClientConn
}
type LeftJoinConexao struct {
	LeftJoin proto.LeftJoinServiceClient
	Conn     *grpc.ClientConn
}
type InnerJoinConexao struct {
	InnerJoin proto.InnerJoinServiceClient
	Conn      *grpc.ClientConn
}

func NewGRPCConn(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
