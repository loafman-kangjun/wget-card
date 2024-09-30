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
		return mvc.Response{Code: 404} // 商品未找到，返回404
	}

	return mvc.View{
		Name: "item.html",
		Data: iris.Map{"product": product},
	}
}

// 处理提交订单请求
func (c *ProductController) PostItemBy(id int, ctx iris.Context) mvc.Result {
	email := ctx.FormValue("email")
	if email == "" {
		return mvc.Response{Code: 400, Text: "邮箱不能为空"} // 检查邮箱
	}

	// 获取产品
	var product models.Product
	if err := c.DB.First(&product, id).Error; err != nil {
		return mvc.Response{Code: 404} // 商品未找到，返回404
	}

	// 创建订单
	order := models.Order{
		ProductID: uint(id),
		Email:     email,
	}
	c.DB.Create(&order)

	return mvc.Response{
		Code: 200,
		Text: "订单已提交",
	}
}

// 处理查询订单的请求
func (c *ProductController) GetOrders(ctx iris.Context) mvc.Result {
	email := ctx.URLParam("email")
	if email == "" {
		return mvc.Response{Code: 400, Text: "邮箱不能为空"} // 检查邮箱
	}

	var orders []models.Order
	if err := c.DB.Where("email = ?", email).Find(&orders).Error; err != nil || len(orders) == 0 {
		// 如果没有找到订单，返回适当的提示
		return mvc.View{
			Name: "orders.html",
			Data: iris.Map{
				"orders": nil, 
				"message": "未找到相关订单", // 中文提示
			},
		}
	}

	return mvc.View{
		Name: "orders.html",
		Data: iris.Map{
			"orders": orders, 
			"message": "",
		},
	}
}