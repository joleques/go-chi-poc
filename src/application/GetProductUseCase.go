package application

import (
	"fmt"
	"github.com/joleques/go-chi-poc/src/domain"
)

func GetProductUseCase(id string, baseDados domain.BaseDados) Product {
	model, err := baseDados.Get(id)
	if err != nil {
		fmt.Println(err.Error())
		return Product{}
	}
	product := Product{Name: model.Nome, Obs: model.Observacao}
	return product
}

func SaveProductUseCase(dto Product, baseDados domain.BaseDados) Result {
	product, err := domain.NewProdutoModel(dto.Name, dto.Obs)
	if err != nil {
		return Result{StatusCod: 400, Message: err.Error()}
	}
	baseDados.Save(product)
	return Result{StatusCod: 200, Message: fmt.Sprintf("Produto %s salvo com sucesso", product.Id)}
}
