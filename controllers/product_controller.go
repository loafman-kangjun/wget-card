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

func (c *ProductController) GetItemBy(id int) mvc.Result {
	var product models.Product
	if err := c.DB.First(&product, id).Error; err != nil {
		return mvc.Response{Code: 404} // 如果没有找到商品返回404错误
	}

	return mvc.View{
		Name: "item.html",
		Data: iris.Map{"product": product},
	}
}
