package request

type RecordsCreate struct {
	Token         string  `json:"token" validate:"required,min=36,max=36"`
	Name          string  `json:"name" validate:"required,min=1,max=150"`
	Calories      int     `json:"calories" validate:"omitempty"`
	Proteins      float64 `json:"proteins" validate:"omitempty"`
	Carbohydrates float64 `json:"carbohydrates" validate:"omitempty"`
	Fats          float64 `json:"fats" validate:"omitempty"`
	Salt          float64 `json:"salt" validate:"omitempty"`
	Sugar         float64 `json:"sugar" validate:"omitempty"`
}

type RecordsDelete struct {
	Token string `json:"token" validate:"required,min=36,max=36"`
	Id    string `json:"id" validate:"required,min=36,max=36"`
}

type RecordsList struct {
	Token    string `json:"token" validate:"required,min=36,max=36"`
	DateFrom string `json:"date_from" validate:"required,datetime"`
	DateTo   string `json:"date_to" validate:"required,datetime"`
}
