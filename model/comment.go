package model

type Comment struct {
	Id        string `json:"id"`
	Pid       string `json:"pid"`
	ArticleId string `json:"articleId"`
	Time      string `json:"time"`
	Content   string `json:"content"`
	User
}

func (c *Comment) TableName() string {
	return "comment"
}