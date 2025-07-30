package models

type Book struct { //published error
	ID          int    `json:"id" form:"id"`
	BookName    string `json:"bookName" form:"bookName"`
	Author      int    `json:"author" form:"author"`
	Quantity    int    `json:"quantity" form:"quantity"`
	Description string `json:"description" form:"description"`
	Published   string `json:"published" form:"published"`
	Page        int    `json:"page" form:"page"`
	Language    int    `json:"bookLanguage" form:"bookLanguage"`
}
