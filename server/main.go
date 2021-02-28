package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

type UserDAO struct {
	DB *sql.DB
}
type App struct {
	Server  *echo.Echo
	UserDAO *UserDAO
}

type User struct {
	ID         string `json:"id"         query:"id"`
	FirstName  string `json:"firstName"  query:"firstName"`
	LastName   string `json:"lastName"   query:"lastName"`
	Email      string `json:"email"      query:"email"`
	UserName   string `json:"userName"   query:"userName"`
	Status     string `json:"status"     query:"status"`
	Department string `json:"department" query:"department"`
}

func main() {

	app := App{}
	app.Initialize(
		"postgres",
		"postgres",
		"postgres")
	app.initializeEndPoints()

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

	a.UserDAO = new(UserDAO)
	a.UserDAO.Init(db)
}

func (a *App) initializeEndPoints() {
	a.Server.Static("/", "assets")
	a.Server.GET("/users", a.getUsers)
	a.Server.GET("/users/:id", a.getUser)
	a.Server.POST("/users", a.addUser)
	a.Server.DELETE("/users/:id", a.deleteUser)
	a.Server.PUT("/users/:id", a.updateUser)

}

func (a *App) Run() {
	a.Server.Logger.Fatal(a.Server.Start(":1323"))
}

func (a *App) addUser(context echo.Context) error {
	log.Println("add user")
	user := new(User)
	if err := context.Bind(user); err != nil {
		log.Println("Couldn't process new user", err.Error())
		return context.String(http.StatusUnprocessableEntity, "Couldn't process new user")
	}

	if err := a.UserDAO.InsertUser(user); err != nil {
		log.Println("Error creating user:", err.Error())
		return context.String(http.StatusInternalServerError, "Error adding user")
	}
	return context.String(http.StatusInternalServerError, "user added")
}

func (a *App) getUsers(context echo.Context) error {
	users, err := a.UserDAO.GetUsers()

	if err != nil {
		log.Println("Error retrieving users:", err.Error())
		return context.String(http.StatusInternalServerError, "Error retrieving users")
	}

	return context.JSON(http.StatusOK, users)
}

func (a *App) getUser(context echo.Context) error {
	stringID := context.Param("id")
	log.Println("get user: ", stringID)
	id, err := strconv.Atoi(stringID)
	if err != nil {
		log.Println("Error retrieving user:", stringID, err.Error())
		return context.String(http.StatusUnprocessableEntity, "bad id")
	}
	user, err := a.UserDAO.GetUser(id)
	if err != nil {
		log.Println("Error retrieving user:", stringID, err.Error())
		return context.String(http.StatusInternalServerError, "Error retrieving users")
	}
	return context.JSON(http.StatusOK, user)
}

func (a *App) deleteUser(context echo.Context) error {
	stringID := context.Param("id")
	log.Println("delete user: ", stringID)
	id, err := strconv.Atoi(stringID)
	if err != nil {
		log.Println("Error deleting bad id: ", stringID, err.Error())
		return context.String(http.StatusUnprocessableEntity, "bad id")
	}

	if err = a.UserDAO.DeleteUser(id); err != nil {
		log.Println("error deleting user: ", stringID, err.Error())
		return context.String(http.StatusInternalServerError, "something went wrong running the delete query")
	}

	return context.String(http.StatusOK, "deleted")
}

func (a *App) updateUser(context echo.Context) error {
	stringID := context.Param("id")
	log.Println("update user: ", stringID)
	id, err := strconv.Atoi(stringID)
	log.Println(id)
	if err != nil {
		log.Println("Error retrieving users:", err.Error())
		return context.String(http.StatusUnprocessableEntity, "bad id: "+stringID)
	}
	user := new(User)
	if err := context.Bind(user); err != nil {
		log.Println("Couldn't process user object", err.Error())
		return context.String(http.StatusUnprocessableEntity, "Couldn't process new user "+stringID)
	}

	if err = a.UserDAO.UpdateUser(id, user); err != nil {
		log.Println("Couldn't update user", err.Error())
		return context.String(http.StatusInternalServerError, "Couldn't update user "+stringID)
	}

	updatedUser, err := a.UserDAO.GetUser(id)
	if err != nil {
		return context.String(http.StatusInternalServerError, "error retrieving user during update "+stringID)
	}
	return context.JSON(http.StatusOK, updatedUser)
}

func (dao *UserDAO) Init(db *sql.DB) {
	dao.DB = db
}

func (dao *UserDAO) GetUsers() ([]User, error) {
	var users []User
	var err error

	rows, err := dao.DB.Query("select user_id, first_name, last_name, email, user_name, user_status from users")

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.UserName, &user.Status); err != nil {
			log.Println("error mapping user data from db", err.Error())
		}
		users = append(users, user)
	}
	return users, err
}

func (dao *UserDAO) GetUser(id int) (User, error) {
	var user User
	var err error
	rows, err := dao.DB.Query("select user_id, first_name, last_name, email, user_name, user_status from users where user_id = $1", id)

	if err != nil {
		log.Println("Error retrieving users:", err.Error())
		return user, err
	}

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.UserName, &user.Status); err != nil {
			log.Println("error mapping user data from db", err.Error())
		}
	}
	return user, err
}

func (dao *UserDAO) InsertUser(user *User) error {
	var err error
	_, err = dao.DB.Exec("INSERT INTO users (user_id, first_name, last_name, email, user_name, user_status) VALUES ($1, $2, $3, $4, $5, $6)", user.ID, user.FirstName, user.LastName, user.Email, user.UserName, user.Status)
	return err
}

func (dao *UserDAO) DeleteUser(id int) error {
	var err error
	_, err = dao.DB.Exec("delete from users where user_id = $1", id)
	return err
}

func (dao *UserDAO) UpdateUser(id int, user *User) error {
	var err error
	_, err = dao.DB.Exec("update users set first_name = $2, last_name = $3, email = $4, user_name = $5, user_status = $6 where user_id = $1", id, user.FirstName, user.LastName, user.Email, user.UserName, user.Status)
	return err
}
