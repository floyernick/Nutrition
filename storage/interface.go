package storage

import "Nutrition/models"

type Storage interface {
	sessionStorage
}

type sessionStorage interface {
	StoreSession(models.Session) error
	DeleteSession(models.Session) error
	GetSession(string) (models.Session, error)
}
