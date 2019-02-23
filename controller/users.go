package controller

import (
	"Nutrition/app/errors"
	"Nutrition/models"
	"Nutrition/models/request"
	"Nutrition/models/response"
	"Nutrition/tools/passwords"
	"Nutrition/tools/uuid"
	"Nutrition/tools/validator"
)

func (controller Controller) UsersSignIn(params request.UsersSignIn) (response.UsersSignIn, error) {

	var result response.UsersSignIn

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	user, err := controller.storage.GetUserByName(params.Name)

	if err != nil {
		return result, errors.InternalError{}
	}

	if !passwords.Validate(params.Password, user.Password) {
		return result, errors.InvalidCredentials{}
	}

	session := models.Session{
		Id:     uuid.Generate(),
		UserId: user.Id,
	}

	err = controller.storage.StoreSession(session)

	if err != nil {
		return result, errors.InternalError{}
	}

	result = response.UsersSignIn{
		Token: session.Id,
	}

	return result, nil

}
