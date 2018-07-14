package command

type UpdateArticle struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Markdowncontent string `json:"markdowncontent"`
	Private         string `json:"private"`
	Tags            string `json:"tags"`
	Status          string `json:"status"`
	Categories      string `json:"categories"`
	Type            string `json:"type"`  
	Description     string `json:"description"`
}