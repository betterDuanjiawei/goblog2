package article

import (
	"goblog2/app/models"
	"goblog2/pkg/route"
)

type Article struct {
	models.BaseModel

	Title, Body string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
