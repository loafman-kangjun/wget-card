package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12"
	"github.com/loafman-kangjun/wget-card/models"
)

type CardController struct {
	Cards []models.Card
}

func (c *CardController) Get() mvc.Result {
	c.Cards = []models.Card{
		{"Card 1", "This is the content of card 1."},
		{"Card 2", "This is the content of card 2."},
		{"Card 3", "This is the content of card 3."},
	}
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"cards": c.Cards}, // 修改为 iris.Map
	}
}