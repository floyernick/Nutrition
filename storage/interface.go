package storage

import "Nutrition/models"

type Storage interface {
	sessionStorage
	userStorage
}

type sessionStorage interface {
	StoreSession(models.Session) error
	DeleteSession(models.Session) error
	GetSession(string) (models.Session, error)
}

type userStorage interface {
	StoreUser(models.User) error
	UpdateUser(models.User) error
	DeleteUser(models.User) error
	GetUser(string) (models.User, error)
}
