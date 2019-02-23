package controller

import (
	"Nutrition/app/errors"
	"Nutrition/models/request"
	"Nutrition/models/response"
	"Nutrition/tools/validator"
)

func (controller Controller) ProductsSearch(params request.ProductsSearch) (response.ProductsSearch, error) {

	var result response.ProductsSearch

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	_, err := controller.storage.GetSession(params.Token)

	if err != nil {
		return result, errors.SessionNotFound{}

	}

	products, err := controller.storage.GetProductsByNamePattern(params.Name)

	if err != nil {
		return result, errors.InternalError{}
	}

	result.Products = make([]response.ProductsSearchProduct, 0, len(products))

	for _, product := range products {
		resultProduct := response.ProductsSearchProduct{
			Id:            product.Id,
			Name:          product.Name,
			ImageUrl:      product.ImageUrl,
			Description:   product.Description,
			Calories:      product.Calories,
			Proteins:      product.Proteins,
			Carbohydrates: product.Carbohydrates,
			Fats:          product.Fats,
			Salt:          product.Salt,
			Sugar:         product.Sugar,
		}
		result.Products = append(result.Products, resultProduct)
	}

	return result, nil

}
