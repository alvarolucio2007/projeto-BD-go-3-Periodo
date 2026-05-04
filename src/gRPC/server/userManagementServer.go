package grpcserver

import (
	"context"
	"log"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
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
		Id: id,
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

func (s *ServerUser) ReadAll(ctx context.Context) (*proto.UsuarioReadResponse, error) {
	log.Printf("função ler todos usuários foi chamada")
	users, err := database.LerTodosUsuarios()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao mostrar todos os usuários no Postgres: %v", err)
	}
	listaProtobuf := make([]*proto.Usuario, 0, len(users))
	for _, u := range users {
		listaProtobuf = append(listaProtobuf, &proto.Usuario{
			Id:       u.ID,
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

func (s *ServerUser) Delete(ctx context.Context, in *proto.UsuarioDeleteRequest) (*emptypb.Empty, error) {
	log.Printf("função deletar usuário foi chamada com %v\n", in)
	err := database.DeleteUsuarios(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao deletar no Postgres: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *ServerUser) Auth(ctx context.Context, in *proto.UsuarioLoginRequest) (*proto.UsuarioLoginResponse, error) {
	log.Printf("Atutenticando usuário %v \n ", in)
	stats, msg, role, err := database.AutenticarUsuario(in.Username, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro crítico: %v", err)
	}
	return &proto.UsuarioLoginResponse{
		Status:   stats,
		Mensagem: msg,
		Role:     role,
	}, nil
}
