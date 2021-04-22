package models

type Response struct {
	Status		string	`json:"status"`
	Totalresult	int		`json:"totalResults"`
	Data	 []Article	`json:"articles"`
}

type Article struct {
	Source 	  	Source	`json:"source"`
	Author		string	`json:"author"`
	Title 		string	`json:"title"`
	Description	string	`json:"description"`
	Url 		string	`json:"url"`
	UrlToImage	string	`json:"urlToImage"`
	PublishedAt	string	`json:"publishedAt"`
	Content 	string	`json:"content"`
}

type Source struct {
	Id 			string	`json:"id"`
	Name 		string	`json:"name"`
}