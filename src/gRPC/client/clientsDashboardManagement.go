package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func (h *HubGeral) DoLerQuantidadeProvaAluno(nome string) (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Dashboard.QuantidadeProvaAluno(ctx, &proto.QuantidadeProvaAlunoRequest{NomeBusca: nome})
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}

func (h *HubGeral) DoLerQuantidadeNotaProvaAluno(nome string) (map[string]models.EstatisticaAluno, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Dashboard.QuantidadeNotaProvaAluno(ctx, &proto.QuantidadeNotaProvaAlunoRequest{NomeBusca: nome})
	if err != nil {
		return nil, err
	}
	response := make(map[string]models.EstatisticaAluno)
	for s, e := range res.Response {
		val := models.EstatisticaAluno{QuantidadeProva: int(e.QuantidadeProva), MediaProvas: float64(e.MediaProva)}
		response[s] = val
	}
	return response, nil
}

func (h *HubGeral) DoLerMediaNotaMateria(nome string) (map[string]models.EstatisticaAluno, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Dashboard.MediaNotaMateria(ctx, &proto.MediaNotaMateriaRequest{NomeCategoria: nome})
	if err != nil {
		return nil, err
	}
	response := make(map[string]models.EstatisticaAluno)
	for s, e := range res.Response {
		val := models.EstatisticaAluno{QuantidadeProva: int(e.QuantidadeProva), MediaProvas: float64(e.MediaProva)}
		response[s] = val
	}
	return response, nil
}

func (h *HubGeral) DoLerDistribuicaoStatusAluno() (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.Dashboard.DistribuicaoStatusAluno(ctx, &proto.DistribuicaoStatusAlunoRequest{})
	if err != nil {
		return nil, err
	}
	return res.Response, nil
}
