package model

import (
	"fmt"
	"time"
)

type JsonTime time.Time

const timeFormat = "2006-01-02 15:04:05"

// 定义序列化
func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("\"%s\"", time.Time(jsonTime).Format(timeFormat))
	return []byte(json), nil
}

// 定义反序列化
func (jsonTime JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	jsonTime = JsonTime(now)
	return
}

type Model struct {
	Id         string   `gorm:"primary_key" json:"id"`
	CreateUser string   `json:"createUser"`
	CreateTime JsonTime `json:"createTime"`
	UpdateUser string   `json:"updateUser"`
	UpdateTime JsonTime `json:"updateTime"`
	IsDelete   string   `json:"isDelete"`
}
