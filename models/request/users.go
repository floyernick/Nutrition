package request

type UsersSignUp struct {
	Name     string `json:"name" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type UsersSignIn struct {
	Name     string `json:"name" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type UsersUpdateNutrientsRates struct {
	Token         string   `json:"token" validate:"required,min=36,max=36"`
	Calories      *int     `json:"calories" validate:"omitempty,min=1"`
	Proteins      *float64 `json:"proteins" validate:"omitempty,min=1"`
	Carbohydrates *float64 `json:"carbohydrates" validate:"omitempty,min=1"`
	Fats          *float64 `json:"fats" validate:"omitempty,min=1"`
	Salt          *float64 `json:"salt" validate:"omitempty,min=1"`
	Sugar         *float64 `json:"sugar" validate:"omitempty,min=1"`
}
