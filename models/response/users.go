package response

type UsersSignUp struct {
	Token string `json:"token"`
}

type UsersSignIn struct {
	Token string `json:"token"`
}

type UsersUpdateNutrientsRates struct{}
