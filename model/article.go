package model

type Article struct {
	Model
	Title       string   `json:"title"`
	Remark      string   `json:"remark"`
	Pic         string   `json:"pic"`
	Content     string   `json:"content"`
	PulishState string   `json:"publishState"`
	PublishTime JsonTime `json:"publishTime"`
}
