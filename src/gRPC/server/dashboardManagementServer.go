package grpcserver

import (
	"context"
	"log"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerDashboard) QuantidadeProvaAluno(ctx context.Context, in *proto.QuantidadeProvaAlunoRequest) (*proto.QuantidadeProvaAlunoResponse, error) {
	log.Printf("função dashboard prova aluno request foi chamada com %v\n", in.NomeBusca)
	res, err := database.LerQuantidadeProvaAluno(in.NomeBusca)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao fazer o relatório: %v", err)
	}
	return &proto.QuantidadeProvaAlunoResponse{
		Response: res,
	}, nil
}

func (s *ServerDashboard) QuantidadeNotaProvaAluno(ctx context.Context, in *proto.QuantidadeNotaProvaAlunoRequest) (*proto.QuantidadeNotaProvaAlunoResponse, error) {
	log.Printf("função dashboard nota prova aluno foi chamada com %v\n", in.NomeBusca)
	res, err := database.LerQuantidadeNotaProvaAluno(in.NomeBusca)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao fazer o relatório: %v", err)
	}
	resCorrect := make(map[string]*proto.EstatisticaAluno)
	for a, b := range res {
		c := &proto.EstatisticaAluno{QuantidadeProva: int64(b.QuantidadeProva), MediaProva: float32(b.MediaProvas)}
		resCorrect[a] = c
	}
	return &proto.QuantidadeNotaProvaAlunoResponse{
		Response: resCorrect,
	}, nil
}

func (s *ServerDashboard) MediaNotaMateria(ctx context.Context, in *proto.MediaNotaMateriaRequest) (*proto.MediaNotaMateriaResponse, error) {
	log.Printf("função media nota materia foi chamada com: %v\n", in.NomeCategoria)
	res, err := database.LerMediaNotaMateria(in.NomeCategoria)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao fazer o relatório: %v", err)
	}
	resCorrect := make(map[string]*proto.EstatisticaAluno)
	for a, b := range res {
		c := &proto.EstatisticaAluno{QuantidadeProva: int64(b.QuantidadeProva), MediaProva: float32(b.MediaProvas)}
		resCorrect[a] = c
	}
	return &proto.MediaNotaMateriaResponse{
		Response: resCorrect,
	}, nil
}

func (s *ServerDashboard) DistribuicaoStatusAluno(ctx context.Context, in *proto.DistribuicaoStatusAlunoRequest) (*proto.DistribuicaoStatusAlunoResponse, error) {
	log.Printf("função distribuicao status aluno foi chamado \n")
	res, err := database.LerDistribuicaoStatusAluno()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao fazer o relatório: %v", err)
	}
	return &proto.DistribuicaoStatusAlunoResponse{
		Response: res,
	}, nil
}
