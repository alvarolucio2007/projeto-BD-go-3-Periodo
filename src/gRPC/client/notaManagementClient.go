package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func (h *HubConexoes) doCreateNota(nota *models.Notas) (uint32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Nota.Create(ctx, &proto.CreateNotaRequest{UsuarioId: nota.UsuarioID, ProvaId: nota.ProvaID, NotaProva: nota.NotaProva})
	if err != nil {
		return 0, err
	}
	return res.NotaId, nil
}
