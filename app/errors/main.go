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

type InvalidCredentials struct{}

func (err InvalidCredentials) Error() string {
	return "invalid credentials"
}

type UserNameIsTaken struct{}

func (err UserNameIsTaken) Error() string {
	return "username is taken"
}

type SessionNotFound struct{}

func (err SessionNotFound) Error() string {
	return "session not found"
}

type UserNotFound struct{}

func (err UserNotFound) Error() string {
	return "user not found"
}
