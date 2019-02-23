package storage

import (
	"Nutrition/models"
	"Nutrition/tools/logger"
	"database/sql"
)

func (storage Service) GetProductsByNamePattern(name string) ([]models.Product, error) {

	products := make([]models.Product, 0)

	query := "SELECT id, name, image_url, description, calories, proteins, carbohydrates, fats, salt, sugar FROM products WHERE name ILIKE '%' || $1 || '%'"

	rows, err := storage.pool.Query(query, name)

	if err != nil && err != sql.ErrNoRows {
		logger.Warning(err.Error())
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.ImageUrl, &product.Description, &product.Calories, &product.Proteins, &product.Carbohydrates, &product.Fats, &product.Salt, &product.Sugar)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil

}
