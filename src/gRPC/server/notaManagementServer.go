package grpcserver

import (
	"context"
	"log"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerNota) Create(ctx context.Context, in *proto.CreateNotaRequest) (*proto.CreateNotaResponse, error) {
	log.Printf("Função create nota foi chamado com %v\n", in)
	modeloNota := models.Notas{UsuarioID: in.UsuarioId, ProvaID: in.ProvaId, NotaProva: in.NotaProva}
	id, err := database.CriarEntradaNotas(modeloNota)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao criar nota : %v", err)
	}
	return &proto.CreateNotaResponse{
		NotaId: id,
	}, nil
}
