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
)

func (s *ServerUser) Create(ctx context.Context, in *proto.UsuarioCreateRequest) (*proto.UsuarioCreateResponse, error) {
	log.Printf("função criar usuário foi chamada com %v\n", in)
	modeloUsuario := models.Usuario{Username: in.Username, Password: in.Password, Role: in.Role}
	id, err := database.CriarEntradaUsuario(modeloUsuario)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao inserir no postres: %v", err)
	}
	return &proto.UsuarioCreateResponse{
		Message: fmt.Sprintf("ID: %d", id),
	}, nil
}
