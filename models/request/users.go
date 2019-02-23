package request

type UsersSignUp struct {
	Name     string `json:"name" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type UsersSignIn struct {
	Name     string `json:"name" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}
