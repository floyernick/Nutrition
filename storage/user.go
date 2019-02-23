package storage

import (
	"database/sql"

	"Nutrition/models"
	"Nutrition/tools/logger"
)

func (storage Service) StoreUser(user models.User) error {

	query := "INSERT INTO users(id, name, password, calories, proteins, carbohydrates, fats, salt, sugar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := storage.pool.Exec(query, user.Id, user.Name, user.Password, user.Calories, user.Proteins, user.Carbohydrates, user.Fats, user.Salt, user.Sugar)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) UpdateUser(user models.User) error {

	query := "UPDATE users SET name = $2, password = $3, calories = $4, proteins = $5, carbohydrates = $6, fats = $7, salt = $8, sugar = $9 WHERE id = $1"

	_, err := storage.pool.Exec(query, user.Id, user.Name, user.Password, user.Calories, user.Proteins, user.Carbohydrates, user.Fats, user.Salt, user.Sugar)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) DeleteUser(user models.User) error {

	query := "DELETE FROM users WHERE id = $1"

	_, err := storage.pool.Exec(query, user.Id)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) GetUser(id string) (models.User, error) {

	var user models.User

	query := "SELECT id, name, password, calories, proteins, carbohydrates, fats, salt, sugar FROM users WHERE id = $1"

	err := storage.pool.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Password, &user.Calories, &user.Proteins, &user.Carbohydrates, &user.Fats, &user.Salt, &user.Sugar)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (storage Service) GetUserByName(name string) (models.User, error) {

	var user models.User

	query := "SELECT id, name, password, calories, proteins, carbohydrates, fats, salt, sugar FROM users WHERE name = $1"

	err := storage.pool.QueryRow(query, name).Scan(&user.Id, &user.Name, &user.Password, &user.Calories, &user.Proteins, &user.Carbohydrates, &user.Fats, &user.Salt, &user.Sugar)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return user, err
	}

	return user, nil
}
