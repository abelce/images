package command

type CreateImage struct {
	Url           string `json:"url"`
	Width         int `json:"width"`
	Height        int `json:"height"`
}