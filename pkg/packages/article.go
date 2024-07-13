package packages

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
