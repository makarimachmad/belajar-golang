package models

type Article struct {
	Id 			int		`json:"id"`
	SourceId	string 	`json:"sourceid"`
	SourceName 	string	`json:"sourcename"`
	Author		string	`json:"author"`
	Title 		string	`json:"title"`
	Description	string	`json:"description"`
	Url 		string	`json:"url"`
	UrlToImage	string	`json:"urlToImage"`
	PublishedAt	string	`json:"publishedAt"`
	Content 	string	`json:"content"`
}