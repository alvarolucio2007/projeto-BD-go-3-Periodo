package grpcclient

import (
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HubConexoes struct {
	User  proto.UsuariosServiceClient
	Prova proto.ProvaServiceClient
	Nota  proto.NotaServiceClient
	Conn  *grpc.ClientConn
}

func ConnectAll(addr string) (*HubConexoes, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &HubConexoes{
		User:  proto.NewUsuariosServiceClient(conn),
		Prova: proto.NewProvaServiceClient(conn),
		Nota:  proto.NewNotaServiceClient(conn),
		Conn:  conn,
	}, nil
}
