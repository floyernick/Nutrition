package presenter

import (
	"Nutrition/app/errors"
	"Nutrition/models/request"
	"net/http"
)

func (presenter Presenter) RecordsCreate(r *http.Request) (interface{}, error) {

	var params request.RecordsCreate

	if err := parseRequestBody(r, &params); err != nil {
		return nil, errors.BadRequest{}
	}

	result, err := presenter.controller.RecordsCreate(params)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (presenter Presenter) RecordsDelete(r *http.Request) (interface{}, error) {

	var params request.RecordsDelete

	if err := parseRequestBody(r, &params); err != nil {
		return nil, errors.BadRequest{}
	}

	result, err := presenter.controller.RecordsDelete(params)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (presenter Presenter) RecordsList(r *http.Request) (interface{}, error) {

	var params request.RecordsList

	if err := parseRequestBody(r, &params); err != nil {
		return nil, errors.BadRequest{}
	}

	result, err := presenter.controller.RecordsList(params)

	if err != nil {
		return nil, err
	}

	return result, nil

}
