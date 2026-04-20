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

func (s *ServerProva) Create(ctx context.Context, in *proto.CreateProvaRequest) (*proto.CreateProvaResponse, error) {
	log.Printf("função criar prova foi chamada com %v\n", in)
	modeloProva := models.Provas{NomeProva: in.NomeProva, TurmaProva: in.TurmaProva, MateriaProva: in.MateriaProva, DataProva: in.DataProva.AsTime()}
	id, err := database.CriarEntradaProva(modeloProva)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao inserir a prova no Postgres: %v", err)
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

func (s *ServerProva) Read(ctx context.Context, in *proto.ReadProvaRequest) (*proto.ProvaLista, error) {
	log.Printf("Função ler prova específica foi chamada com %v\n")
	provas, err := database.ProcurarProvaNome(in.NomeProva)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao procurar a prova: %v", err)
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

func (s *ServerProva) Update(ctx context.Context, in *proto.UpdateProvaRequest) (*emptypb.Empty, error) {
	log.Printf("Função update de prova foi chamada com %v\n")
	err := database.UpdateProvas(in.Id, models.Provas{NomeProva: in.NovoNome, TurmaProva: in.NovaTurma, MateriaProva: in.NovaMateria, DataProva: in.DataProva.AsTime()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erro ao editar a prova: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *ServerProva) Delete(ctx context.Context, in *proto.DeleteProvaRequest) (*emptypb.Empty, error) {
	log.Printf("Função delete de prova foi chamada com %v\n")
	err := database.DeleteProvas(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erro ao deletar prova: %v", err)
	}
	return &emptypb.Empty{}, nil
}
