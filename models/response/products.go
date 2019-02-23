package response

type ProductsSearch struct {
	Products []ProductsSearchProduct `json:"products"`
}

type ProductsSearchProduct struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	ImageUrl      string  `json:"image_url"`
	Description   string  `json:"description"`
	Calories      int     `json:"calories"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
	Fats          float64 `json:"fats"`
	Salt          float64 `json:"salt"`
	Sugar         float64 `json:"sugar"`
}
