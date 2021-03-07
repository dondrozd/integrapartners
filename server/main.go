package main

import (
	"database/sql"
	"fmt"
	"log"
	"server/controller"
	"server/dao"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

type App struct {
	Server  *echo.Echo
	UserDAO *dao.UserDAO
}

func main() {
	app := App{}
	app.Initialize("postgres", "postgres", "postgres")
	app.Run()
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Server = echo.New()

	a.UserDAO = new(dao.UserDAO)
	a.UserDAO.Init(db)
}

func (a *App) Run() {
	a.Server.Static("/", "assets")
	controller.RegisterNewUserResource(a.UserDAO, a.Server)
	a.Server.Logger.Fatal(a.Server.Start(":1323"))
}
