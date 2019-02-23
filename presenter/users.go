package presenter

import (
	"Nutrition/app/errors"
	"Nutrition/models/request"
	"net/http"
)

func (presenter Presenter) UsersSignIn(r *http.Request) (interface{}, error) {

	var params request.UsersSignIn

	if err := parseRequestBody(r, &params); err != nil {
		return nil, errors.BadRequest{}
	}

	result, err := presenter.controller.UsersSignIn(params)

	if err != nil {
		return nil, err
	}

	return result, nil

}
