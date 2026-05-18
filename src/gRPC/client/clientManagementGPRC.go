// Package grpcclient gerencia os clients do gRPC.
package grpcclient

import (
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HubGeral struct {
	User      proto.UsuariosServiceClient
	Prova     proto.ProvaServiceClient
	Nota      proto.NotaServiceClient
	LeftJoin  proto.LeftJoinServiceClient
	InnerJoin proto.InnerJoinServiceClient
	Dashboard proto.DashboardServiceClient
	Conn      *grpc.ClientConn
}

func ConnectAll(addr string) (*HubGeral, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &HubGeral{
		User:      proto.NewUsuariosServiceClient(conn),
		Prova:     proto.NewProvaServiceClient(conn),
		Nota:      proto.NewNotaServiceClient(conn),
		InnerJoin: proto.NewInnerJoinServiceClient(conn),
		LeftJoin:  proto.NewLeftJoinServiceClient(conn),
		Dashboard: proto.NewDashboardServiceClient(conn),
		Conn:      conn,
	}, nil
}
