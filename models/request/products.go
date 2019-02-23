package request

type ProductsSearch struct {
	Token string `json:"token" validate:"required,min=36,max=36"`
	Name  string `json:"name" validate:"required,min=3"`
}
