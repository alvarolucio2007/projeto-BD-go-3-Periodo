package grpcserver

import (
	"context"
	"log"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerProva) Create(ctx context.Context, in *proto.CreateProvaRequest) (*proto.CreateProvaResponse, error) {
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

func (s *ServerProva) ReadAll(ctx context.Context) (*proto.ProvaLista, error) {
	log.Printf("Função ler todas as provas foi chamada")
	provas, err := database.LerTodasProvas()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao procurar todas as provas: %v", err)
	}
	listaProvas := make([]*proto.Prova, 0, len(provas))
	for _, p := range provas {
		listaProvas = append(listaProvas, &proto.Prova{
			Id:           int32(p.ID),
			NomeProva:    p.NomeProva,
			TurmaProva:   p.TurmaProva,
			MateriaProva: p.MateriaProva,
			DataProva:    timestamppb.New(p.DataProva),
		})
	}
	return &proto.ProvaLista{
		Provas: listaProvas,
	}, nil
}
