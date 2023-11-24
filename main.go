package main

import (
	"webserver/config"
	"webserver/controller"
	"webserver/middleware"

	_ "webserver/docs"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Api Document Employee
// @version 1.0
// @Description Api Document Employee
// @termsOfService http://localhost
// @contact.name pegadaian
// @contact.email pegadaian.id
// @license.name pegadaian 1.0
// @license.url
// @host localhost:9000
// @BasePath /
func main() {
	// err := godotenv.Load()
	config.Connect()
	r := echo.New()
	r.Use(echoMid.CSRFWithConfig(echoMid.CSRFConfig{
		TokenLookup: "header:" + config.CSRFTokenHeader,
		ContextKey:  config.CSRFKey,
	}))

	r.GET("/index", controller.Index)
	r.POST("/sayhello", controller.SayHello)
	// r.GET("/",controller.HelloWorld)
	// r.GET("/json", controller.JsonMap)
	// r.GET("/page1", controller.Page)
	// r.Any("/user", controller.User)

	emm := r.Group("/employee")
	emm.Use(middleware.Authentication())
	emm.PUT("/", controller.UpdateEmployee)
	emm.DELETE("/:id", controller.DeleteEmployee)

	itm := r.Group("/item")
	itm.Use(middleware.Authentication())
	itm.POST("/", controller.CreateItem)

	r.POST("/register", controller.CreateEmployee)
	r.POST("/login", controller.UserLogin)
	//route for swagger
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	r.Logger.Fatal(r.Start(":9000"))
}
