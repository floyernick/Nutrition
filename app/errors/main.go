package errors

type InternalError struct{}

func (err InternalError) Error() string {
	return "internal error"
}

type BadRequest struct{}

func (err BadRequest) Error() string {
	return "bad request"
}

type InvalidParams struct{}

func (err InvalidParams) Error() string {
	return "invalid params"
}
