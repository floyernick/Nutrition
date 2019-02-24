package response

type RecordsCreate struct {
	Id string `json:"id"`
}

type RecordsDelete struct{}

type RecordsList struct {
	Records []RecordsListRecord         `json:"records"`
	Stats   map[string]RecordsListStats `json:"stats"`
}

type RecordsListRecord struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Calories      int     `json:"calories"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
	Fats          float64 `json:"fats"`
	Salt          float64 `json:"salt"`
	Sugar         float64 `json:"sugar"`
	Created       string  `json:"created"`
}

type RecordsListStats struct {
	Calories      int     `json:"calories"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
	Fats          float64 `json:"fats"`
	Salt          float64 `json:"salt"`
	Sugar         float64 `json:"sugar"`
}
