package storage

import (
	"database/sql"

	"Nutrition/models"
	"Nutrition/tools/logger"
)

func (storage Service) StoreSession(session models.Session) error {

	query := "INSERT INTO sessions(id, user_id) VALUES($1, $2)"

	_, err := storage.pool.Exec(query, session.Id, session.UserId)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) DeleteSession(session models.Session) error {

	query := "DELETE FROM sessions WHERE id = $1"

	_, err := storage.pool.Exec(query, session.Id)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) GetSession(id string) (models.Session, error) {

	var session models.Session

	query := "SELECT id, user_id FROM sessions WHERE id = $1"

	err := storage.pool.QueryRow(query, id).Scan(&session.Id, &session.UserId)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return session, err
	}

	return session, nil
}
