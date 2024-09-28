package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/loafman-kangjun/wget-card/models"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func (c *ProductController) Get() mvc.Result {
	var products []models.Product
	c.DB.Find(&products)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"products": products},
	}
}
