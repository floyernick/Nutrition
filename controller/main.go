package controller

import (
	"Nutrition/storage"
)

type Controller struct {
	storage storage.Storage
}

func Init(storage storage.Storage) Controller {
	return Controller{storage}
}
