package models

type Record struct {
	Id            string
	UserId        string
	Name          string
	Calories      int
	Proteins      float64
	Carbohydrates float64
	Fats          float64
	Salt          float64
	Sugar         float64
	Created       string
}

func (record *Record) Exists() bool {
	return record.Id != ""
}

func (record *Record) BelongsTo(user User) bool {
	return record.UserId == user.Id
}
