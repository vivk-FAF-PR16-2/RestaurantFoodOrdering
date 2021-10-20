package dto

type Item struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	PreparationTime  int     `json:"preparation-time"`
	Complexity       int     `json:"complexity"`
	CookingApparatus *string `json:"cooking-apparatus"`
}
