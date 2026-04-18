package grpcserver

import (
	"context"
	"fmt"
	"log"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerUser) Create(ctx context.Context, in *proto.UsuarioCreateRequest) (*proto.UsuarioCreateResponse, error) {
	log.Printf("função criar usuário foi chamada com %v\n", in)
	modeloUsuario := models.Usuario{Username: in.Username, Password: in.Password, Role: in.Role}
	id, err := database.CriarEntradaUsuario(modeloUsuario)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao inserir no postgres: %v", err)
	}
	return &proto.UsuarioCreateResponse{
		Message: fmt.Sprintf("ID: %d", id),
	}, nil
}

func (s *ServerUser) Read(ctx context.Context, in *proto.UsuarioReadRequest) (*proto.UsuarioReadResponse, error) {
	log.Printf("função ler usuário foi chamada com %v\n", in)
	user, err := database.ProcurarUsuario(in.Nome)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao procurar no postgres: %v", err)
	}
	listaProtobuf := make([]*proto.Usuario, 0, len(user))
	for _, u := range user {
		listaProtobuf = append(listaProtobuf, &proto.Usuario{
			Id:       uint32(u.ID),
			Username: u.Username,
			Password: u.Password,
			Role:     u.Role,
		})
	}
	return &proto.UsuarioReadResponse{
		Usuarios: listaProtobuf,
	}, nil
}

func (s *ServerUser) Update(ctx context.Context, in *proto.UsuarioUpdateRequest) (*emptypb.Empty, error) {
	log.Printf("função atualizar usuário foi chamada com %v\n", in)
	modeloUsuario := models.Usuario{Username: in.Username, Password: in.Password, Role: in.Role}
	err := database.UpdateUsuarios(in.Id, modeloUsuario)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erro ao atualizar no Postgres: %v", err)
	}
	return &emptypb.Empty{}, nil
}
