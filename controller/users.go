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

func (controller Controller) UsersSignUp(params request.UsersSignUp) (response.UsersSignUp, error) {

	var result response.UsersSignUp

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	user, err := controller.storage.GetUserByName(params.Name)

	if err != nil {
		return result, errors.InternalError{}
	}

	if user.Exists() {
		return result, errors.UserNameIsTaken{}
	}

	user = models.User{
		Id:       uuid.Generate(),
		Name:     params.Name,
		Password: passwords.Encode(params.Password),
	}

	err = controller.storage.StoreUser(user)

	if err != nil {
		return result, errors.InternalError{}
	}

	session := models.Session{
		Id:     uuid.Generate(),
		UserId: user.Id,
	}

	err = controller.storage.StoreSession(session)

	if err != nil {
		return result, errors.InternalError{}
	}

	result = response.UsersSignUp{
		Token: session.Id,
	}

	return result, nil

}

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
