package domain

type Type struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type TypeItem struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name    string `json:"label"`
	Value    string `json:"value"`
}