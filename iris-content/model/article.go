package model

type Article struct {
	Model
	Title        string   `json:"title"`
	Remark       string   `json:"remark"`
	Pic          string   `json:"pic"`
	PublishState string   `json:"publishState"`
	PublishTime  JsonTime `json:"publishTime"`
}

type ArticleDetail struct {
	ArticleId string `json:"articleId"`
	Content   string `json:"content"`
}
