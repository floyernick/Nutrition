package storage

import "Nutrition/models"

type Storage interface {
	sessionStorage
	userStorage
	productStorage
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
	GetUserByName(string) (models.User, error)
}

type productStorage interface {
	GetProductsByNamePattern(string) ([]models.Product, error)
}
