package grpcserver

import (
	"context"
	"log"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *ServerNota) Read(ctx context.Context, in *proto.ReadNotaRequest) (*proto.NotasResponse, error) {
	log.Printf("Função read nota foi chamada com %v\n", in)
	notas, err := database.BuscarNotas(in.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erro ao procurar todas as notas: %v\n", err)
	}
	listaNotas := make([]*proto.InnerJoin, 0, len(notas))
	for _, n := range notas {
		listaNotas = append(listaNotas, &proto.InnerJoin{
			Username:  n.Username,
			NomeProva: n.NomeProva,
			NotaProva: n.NotaProva,
			DataProva: timestamppb.New(n.DataAplicacao),
		})
	}
	return &proto.NotasResponse{
		Response: listaNotas,
	}, nil
}

func (s *ServerNota) Update(ctx context.Context, in *proto.UpdateNotaRequest) (*emptypb.Empty, error) {
	log.Printf("Função update nota foi chamada com %v\n", in)
	nota := models.Notas{UsuarioID: in.AlunoId, ProvaID: in.ProvaId, NotaProva: in.ValorNota}
	err := database.UpdateNotas(in.NotaId, nota)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erro ao dar update na nota: %v", err)
	}
	return &emptypb.Empty{}, nil
}
