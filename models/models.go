package models

import (
	"time"
)

type Articles struct {
	Id      int       `form:"-"`
	Title   string    `form:"title" required`
	Content string    `orm:";type(text)" form:"content,textarea"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (a *Articles) TableName() string {
	return "articles"
}
