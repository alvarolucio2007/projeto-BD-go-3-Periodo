package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *HubConexoes) doCreateProva(prova *models.Provas) (uint32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Prova.Create(ctx, &proto.CreateProvaRequest{NomeProva: prova.NomeProva, TurmaProva: prova.NomeProva, MateriaProva: prova.MateriaProva, DataProva: timestamppb.New(prova.DataProva)})
	if err != nil {
		return 0, err
	}
	return res.IdProva, nil
}

func (h *HubConexoes) doReadAllProva() ([]*models.Provas, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Prova.ReadAll(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	result := make([]*models.Provas, 0, len(res.Provas))
	for _, u := range res.Provas {
		prova := models.Provas{
			ID:           uint32(u.Id),
			NomeProva:    u.NomeProva,
			TurmaProva:   u.TurmaProva,
			MateriaProva: u.MateriaProva,
			DataProva:    u.DataProva.AsTime(),
		}
		result = append(result, &prova)
	}
	return result, nil
}
