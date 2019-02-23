package models

type Product struct {
	Id            string
	Name          string
	ImageUrl      string
	Description   string
	Calories      int
	Proteins      float64
	Carbohydrates float64
	Fats          float64
	Salt          float64
	Sugar         float64
}
