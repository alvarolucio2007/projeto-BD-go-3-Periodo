package grpcserver

import (
	"log"
	"net"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/proto"
	"google.golang.org/grpc"
)

const addr string = "0.0.0.0:50051"

type ServerUser struct {
	proto.UnimplementedUsuariosServiceServer
}
type ServerProva struct {
	proto.UnimplementedProvaServiceServer
}
type ServerNota struct {
	proto.UnimplementedNotaServiceServer
}
type ServerLeftJoin struct {
	proto.UnimplementedLeftJoinServiceServer
}
type ServerInnerJoin struct {
	proto.UnimplementedInnerJoinServiceServer
}

func StartServerGeralGRPC() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Falha ao abrir a porta gRPC: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUsuariosServiceServer(s, &ServerUser{})
	proto.RegisterProvaServiceServer(s, &ServerProva{})
	proto.RegisterNotaServiceServer(s, &ServerNota{})
	proto.RegisterLeftJoinServiceServer(s, &ServerLeftJoin{})
	proto.RegisterInnerJoinServiceServer(s, &ServerInnerJoin{})
	log.Printf("Servidor Usuário gRPC rodando em %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir no servidor gRPC do usuário: %v", err)
	}
}
