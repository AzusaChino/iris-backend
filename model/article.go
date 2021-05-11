package model

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Remark  string `json:"remark"`
	Pic     string `json:"pic"`
	Content string `json:"content"`
}

func (a *Article) TableName() string {
	return "article"
}
