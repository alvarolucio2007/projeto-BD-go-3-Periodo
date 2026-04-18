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

func (s *ServerNota) Create(ctx context.Context, in *proto.CreateProvaRequest) (*proto.CreateProvaResponse, error) {
	log.Printf("função criar prova foi chamada com %v\n", in)
	modeloProva := models.Provas{NomeProva: in.NomeProva, TurmaProva: in.TurmaProva, MateriaProva: in.MateriaProva, DataProva: in.DataProva.AsTime()}
	id, err := database.CriarEntradaProva(modeloProva)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao inserir a nota no Postgres: %v", err)
	}
	return &proto.CreateProvaResponse{
		IdProva: id,
	}, nil
}
