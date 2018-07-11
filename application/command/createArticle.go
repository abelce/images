package command

type CreateArticle struct {
	Title           string `json:"title"`
	Markdowncontent string `json:"markdowncontent"`
	Private         string `json:"private"`
	Tags            string `json:"tags"`
	Status          string `json:"status"`
	Categories      string `json:"categories"`
	Type            string `json:"type"`   //original
	Description     string `json:"description"`
}