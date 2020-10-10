package models


type Article struct {
	Id 			string	`json:"id"`
	Name 		string	`json:"name"`
	Author		string	`json:"author"`
	Title 		string	`json:"title"`
	Description	string	`json:"description"`
	Url 		string	`json:"url"`
	UrlToImage	string	`json:"urlToImage"`
	PublishedAt	string	`json:"publishedAt"`
	Content 	string	`json:"content"`
}