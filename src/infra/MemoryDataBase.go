package infra

import (
	"errors"
	"fmt"
	"github.com/joleques/go-chi-poc/src/domain"
)

type MemoryDataBase struct {
	Data map[string]*domain.ProdutoModel
}

func NewMemoryDataBase() *MemoryDataBase {
	data := map[string]*domain.ProdutoModel{}
	return &MemoryDataBase{Data: data}
}

func (dataBase MemoryDataBase) Status() {
	fmt.Println("Base ok!")
}

func (dataBase *MemoryDataBase) Save(produto *domain.ProdutoModel) {
	dataBase.Data[produto.Id] = produto
}

func (dataBase *MemoryDataBase) Get(id string) (*domain.ProdutoModel, error) {
	model := dataBase.Data[id]
	if model == nil {
		return nil, errors.New("Não foi encontardo o produto")
	}
	return model, nil
}
