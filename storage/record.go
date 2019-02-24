package storage

import (
	"Nutrition/models"
	"Nutrition/tools/logger"
	"database/sql"
)

func (storage Service) StoreRecord(record models.Record) error {

	query := "INSERT INTO records(id, user_id, name, calories, proteins, carbohydrates, fats, salt, sugar, created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	_, err := storage.pool.Exec(query, record.Id, record.UserId, record.Name, record.Calories, record.Proteins, record.Carbohydrates, record.Fats, record.Salt, record.Sugar, record.Created)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) DeleteRecord(record models.Record) error {

	query := "DELETE FROM records WHERE id = $1"

	_, err := storage.pool.Exec(query, record.Id)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) GetRecord(id string) (models.Record, error) {

	var record models.Record

	query := "SELECT id, user_id, name, calories, proteins, carbohydrates, fats, salt, sugar, created FROM records WHERE id = $1"

	err := storage.pool.QueryRow(query, id).Scan(&record.Id, &record.UserId, &record.Name, &record.Calories, &record.Proteins, &record.Carbohydrates, &record.Fats, &record.Salt, &record.Sugar, &record.Created)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return record, err
	}

	return record, nil

}
