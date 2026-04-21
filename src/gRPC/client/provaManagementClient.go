package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
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
