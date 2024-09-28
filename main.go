package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/loafman-kangjun/wget-card/controllers"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	mvc.New(app.Party("/")).Handle(new(controllers.CardController))

	app.StaticWeb("/public", "./public")

	app.Listen(":8080")
}
