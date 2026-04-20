package grpcserver

import (
	"context"
	"database/sql"
	"log"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerInnerJoin) InnerJoin(ctx context.Context, in *emptypb.Empty) (*proto.InnerJoinResponse, error) {
	log.Printf("Função inner join foi chamada")
	sliceInnerJoin, err := database.InnerJoin()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao dar Inner Join: %v", err)
	}
	listaInner := make([]*proto.NotaDetalhada, 0, len(sliceInnerJoin))
	for _, p := range sliceInnerJoin {
		listaInner = append(listaInner, &proto.NotaDetalhada{
			Username:  p.Username,
			NomeProva: p.NomeProva,
			NotaProva: p.NotaProva,
			DataProva: timestamppb.New(p.DataAplicacao),
		})
	}
	return &proto.InnerJoinResponse{
		Itens: listaInner,
	}, nil
}

func nullStringToPointer(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func nullFloatToPointer(nf sql.NullFloat64) *float32 {
	if !nf.Valid {
		return nil
	}
	f := float32(nf.Float64)
	return &f
}

func (s *ServerLeftJoin) LeftJoin(ctx context.Context, in *emptypb.Empty) (*proto.LeftJoinResponse, error) {
	log.Printf("Função left join foi chamada")
	provas, err := database.LeftJoin()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao dar left join: %v", err)
	}
	listaLeft := make([]*proto.LeftJoin, 0, len(provas))
	for _, p := range provas {
		listaLeft = append(listaLeft, &proto.LeftJoin{
			Username:  p.Username,
			NomeProva: nullStringToPointer(p.NomeProva),
			NotaProva: nullFloatToPointer(p.NotaProva),
			DataProva: timestamppb.New(p.DataAplicacao),
		})
	}
	return &proto.LeftJoinResponse{
		Resultado: listaLeft,
	}, nil
}
