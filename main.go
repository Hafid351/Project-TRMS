package main

import (
	"trms/app/routes"
	"trms/app/services"

	/*
		"github.com/casbin/casbin/v2"
		gormadapter "github.com/casbin/gorm-adapter/v3"
	*/
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func NomorTabel(index int, page int, perPage int) int {
	return (page-1)*perPage + index + 1
}

func main() {
	services.ConnectDB()
	/*
		a, _ := gormadapter.NewAdapterByDB(services.DB.Db)
		e, _ := casbin.NewEnforcer("casbin/model.conf", a)
		e.LoadPolicy()
	*/
	engine := html.New("./app/views", ".html")
	engine.AddFunc("nomorTabel", NomorTabel)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	routes.Handlers(app)

	app.Listen(":3080")
	//log.Fatal(app.ListenAndServe("localhost:3000", nil))
}
