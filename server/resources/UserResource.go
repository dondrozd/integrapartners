package resources

import (
	"log"
	"net/http"
	"server/daos"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserResource struct {
	UserDAO *daos.UserDAO
}

type EchoContextInterface interface {
	Bind(i interface{}) error
	String(code int, s string) error
	JSON(code int, i interface{}) error
	Param(name string) string
}

func RegisterNewUserResource(userDAO *daos.UserDAO, server *echo.Echo) *UserResource {
	userResource := new(UserResource)
	userResource.UserDAO = userDAO
	userResource.initialize(server)
	return userResource
}

func (resource *UserResource) initialize(server *echo.Echo) {
	server.GET("/api/users", func(context echo.Context) error { return resource.GetUsers(context) })
	server.GET("/api/users/:id", func(context echo.Context) error { return resource.GetUser(context) })
	server.POST("/api/users", func(context echo.Context) error { return resource.AddUser(context) })
	server.DELETE("/api/users/:id", func(context echo.Context) error { return resource.DeleteUser(context) })
	server.PUT("/api/users/:id", func(context echo.Context) error { return resource.UpdateUser(context) })
}

func (resource *UserResource) AddUser(context EchoContextInterface) error {
	log.Println("add user")
	user := new(daos.User)
	if err := context.Bind(user); err != nil {
		log.Println("Couldn't process new user", err.Error())
		return context.String(http.StatusUnprocessableEntity, "Couldn't process new user")
	}

	if err := resource.UserDAO.InsertUser(user); err != nil {
		log.Println("Error creating user:", err.Error())
		return context.String(http.StatusInternalServerError, "Error adding user")
	}
	return context.String(http.StatusInternalServerError, "user added")
}

func (resource *UserResource) GetUsers(context EchoContextInterface) error {
	log.Println("get users")

	users, err := resource.UserDAO.GetUsers()

	if err != nil {
		log.Println("Error retrieving users:", err.Error())
		return context.String(http.StatusInternalServerError, "Error retrieving users")
	}

	return context.JSON(http.StatusOK, users)
}

func (resource *UserResource) GetUser(context EchoContextInterface) error {
	stringID := context.Param("id")
	log.Println("get user: ", stringID)
	id, err := strconv.Atoi(stringID)
	if err != nil {
		log.Println("Error retrieving user:", stringID, err.Error())
		return context.String(http.StatusUnprocessableEntity, "bad id")
	}
	user, err := resource.UserDAO.GetUser(id)
	if err != nil {
		log.Println("Error retrieving user:", stringID, err.Error())
		return context.String(http.StatusInternalServerError, "Error retrieving users")
	}
	return context.JSON(http.StatusOK, user)
}

func (resource *UserResource) DeleteUser(context EchoContextInterface) error {
	stringID := context.Param("id")
	log.Println("delete user: ", stringID)
	id, err := strconv.Atoi(stringID)
	if err != nil {
		log.Println("Error deleting bad id: ", stringID, err.Error())
		return context.String(http.StatusUnprocessableEntity, "bad id")
	}

	if err = resource.UserDAO.DeleteUser(id); err != nil {
		log.Println("error deleting user: ", stringID, err.Error())
		return context.String(http.StatusInternalServerError, "something went wrong running the delete query")
	}

	return context.String(http.StatusOK, "deleted")
}

func (resource *UserResource) UpdateUser(context echo.Context) error {
	stringID := context.Param("id")
	log.Println("update user: ", stringID)
	id, err := strconv.Atoi(stringID)
	log.Println(id)
	if err != nil {
		log.Println("Error retrieving users:", err.Error())
		return context.String(http.StatusUnprocessableEntity, "bad id: "+stringID)
	}
	user := new(daos.User)
	if err := context.Bind(user); err != nil {
		log.Println("Couldn't process user object", err.Error())
		return context.String(http.StatusUnprocessableEntity, "Couldn't process new user "+stringID)
	}

	if err = resource.UserDAO.UpdateUser(id, user); err != nil {
		log.Println("Couldn't update user", err.Error())
		return context.String(http.StatusInternalServerError, "Couldn't update user "+stringID)
	}

	updatedUser, err := resource.UserDAO.GetUser(id)
	if err != nil {
		return context.String(http.StatusInternalServerError, "error retrieving user during update "+stringID)
	}
	return context.JSON(http.StatusOK, updatedUser)
}
