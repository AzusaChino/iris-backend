package model

type Comment struct {
	Model
	Pid       string `json:"pid"`
	ArticleId string `json:"articleId"`
	Content   string `json:"content"`
}
