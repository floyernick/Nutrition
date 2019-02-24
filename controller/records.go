package controller

import (
	"Nutrition/app/errors"
	"Nutrition/models"
	"Nutrition/models/request"
	"Nutrition/models/response"
	"Nutrition/tools/dates"
	"Nutrition/tools/uuid"
	"Nutrition/tools/validator"
)

func (controller Controller) RecordsCreate(params request.RecordsCreate) (response.RecordsCreate, error) {

	var result response.RecordsCreate

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	session, err := controller.storage.GetSession(params.Token)

	if err != nil {
		return result, errors.SessionNotFound{}

	}

	record := models.Record{
		Id:            uuid.Generate(),
		UserId:        session.UserId,
		Name:          params.Name,
		Calories:      params.Calories,
		Proteins:      params.Proteins,
		Carbohydrates: params.Carbohydrates,
		Fats:          params.Fats,
		Salt:          params.Salt,
		Sugar:         params.Sugar,
		Created:       dates.Generate(),
	}

	err = controller.storage.StoreRecord(record)

	if err != nil {
		return result, errors.InternalError{}
	}

	result = response.RecordsCreate{
		Id: record.Id,
	}

	return result, nil
}

func (controller Controller) RecordsDelete(params request.RecordsDelete) (response.RecordsDelete, error) {

	var result response.RecordsDelete

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	session, err := controller.storage.GetSession(params.Token)

	if err != nil {
		return result, errors.SessionNotFound{}

	}

	user, err := controller.storage.GetUser(session.UserId)

	record, err := controller.storage.GetRecord(params.Id)

	if err != nil {
		return result, errors.InternalError{}
	}

	if !record.Exists() {
		return result, errors.RecordNotFound{}
	}

	if !record.BelongsTo(user) {
		return result, errors.ActionNotAllowed{}
	}

	err = controller.storage.DeleteRecord(record)

	if err != nil {
		return result, errors.InternalError{}
	}

	return result, nil
}
