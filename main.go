package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/loafman-kangjun/wget-card/controllers"
	"github.com/loafman-kangjun/wget-card/models"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	// Connect to PostgreSQL
	db, err := models.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema
	models.Migrate(db)

	// Inject the database connection into the controller
	productController := new(controllers.ProductController)
	productController.DB = db

	mvcApp := mvc.New(app.Party("/"))
	mvcApp.Handle(productController)

	// Serve static files
	app.HandleDir("/public", "./public")

	app.Listen(":8080")
}
