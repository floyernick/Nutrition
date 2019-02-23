package models

type User struct {
	Id            string
	Name          string
	Password      string
	Calories      int
	Proteins      float64
	Carbohydrates float64
	Fats          float64
	Salt          float64
	Sugar         float64
}
