package domain

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

type BaseDados interface {
	Save(produto *ProdutoModel)
	Get(id string) (*ProdutoModel, error)
}

type ProdutoModel struct {
	Id         string `json:"id"`
	Nome       string `json:"nome"`
	Observacao string `json:"observacao"`
}

func NewProdutoModel(nome string, obs string) (*ProdutoModel, error) {
	id := uuid.NewV4()
	model := &ProdutoModel{Id: id.String(), Nome: nome, Observacao: obs}
	err := model.verify()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (model ProdutoModel) verify() error {
	if model.Nome == "" {
		return errors.New("Nome é obrigatório")
	}
	return nil
}
