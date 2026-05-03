// Package grpcclient gerencia os clients do gRPC.
package grpcclient

import (
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HubConexoes struct {
	User      proto.UsuariosServiceClient
	Prova     proto.ProvaServiceClient
	Nota      proto.NotaServiceClient
	LeftJoin  proto.LeftJoinServiceClient
	InnerJoin proto.InnerJoinServiceClient
	Conn      *grpc.ClientConn
}
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

func ConnectAll(addr string) (*HubConexoes, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &HubConexoes{
		User:      proto.NewUsuariosServiceClient(conn),
		Prova:     proto.NewProvaServiceClient(conn),
		Nota:      proto.NewNotaServiceClient(conn),
		InnerJoin: proto.NewInnerJoinServiceClient(conn),
		LeftJoin:  proto.NewLeftJoinServiceClient(conn),
		Conn:      conn,
	}, nil
}
