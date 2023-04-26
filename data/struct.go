package data

type Recipe struct {
	Id          string            `json:"Id"`
	RecipeName  string            `json:"RecipeName"`
	Source      string            `json:"Source"`
	PrepTime    string            `json:"PrepTime"`
	CookTime    string            `json:"CookTime"`
	ServingSize int               `json:"ServingSize"`
	Ingredients map[string]string `json:"Ingredients"`
	Directions  map[int]string    `json:"Directions"`
	Tags        []string          `json:"Tags"`
}

var ThisVariable Recipe
