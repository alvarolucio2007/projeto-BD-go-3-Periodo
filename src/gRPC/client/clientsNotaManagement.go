package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func (h *HubGeral) DoCreateNota(nota *models.Notas) (uint32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Nota.Create(ctx, &proto.CreateNotaRequest{UsuarioId: nota.UsuarioID, ProvaId: nota.ProvaID, NotaProva: nota.NotaProva})
	if err != nil {
		return 0, err
	}
	return res.NotaId, nil
}

func (h *HubGeral) DoReadNota(username string) ([]*models.InnerJoinType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Nota.Read(ctx, &proto.ReadNotaRequest{Username: username})
	if err != nil {
		return nil, err
	}
	if res == nil || res.Response == nil {
		return []*models.InnerJoinType{}, nil
	}
	response := make([]*models.InnerJoinType, 0, len(res.Response))
	for _, n := range res.Response {
		response = append(response, &models.InnerJoinType{
			Username:      n.Username,
			NomeProva:     n.NomeProva,
			NotaProva:     n.NotaProva,
			DataAplicacao: n.DataProva.AsTime(),
		})
	}
	return response, nil
}

func (h *HubGeral) DoUpdateNota(nota *models.Notas) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := h.Nota.Update(ctx, &proto.UpdateNotaRequest{NotaId: nota.ID, ValorNota: nota.NotaProva, AlunoId: nota.UsuarioID, ProvaId: nota.ProvaID})
	if err != nil {
		return err
	}
	return nil
}

func (h *HubGeral) DoDeleteNota(id uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := h.Nota.Delete(ctx, &proto.DeleteNotaRequest{NotaId: id})
	if err != nil {
		return err
	}
	return nil
}
