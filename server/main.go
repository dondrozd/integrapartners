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
	Server *echo.Echo
	DB     *sql.DB
}

func main() {
	app := App{}
	app.Initialize("postgres", "postgres", "postgres")
	app.Run()
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Server = echo.New()
}

func (a *App) Run() {
	a.Server.Static("/", "assets")
	userDao := new(dao.UserDAO)
	userDao.Init(a.DB)
	controller.RegisterNewUserResource(userDao, a.Server)
	a.Server.Logger.Fatal(a.Server.Start(":1323"))
}
