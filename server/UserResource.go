package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserResource struct {
	UserDAO *UserDAO
}

func RegisterNewUserResource(userDAO *UserDAO, server *echo.Echo) UserResource {
	userResource := new(UserResource)
	userResource.UserDAO = userDAO
	userResource.initialize(server)
	return *userResource
}

func (resorce *UserResource) initialize(server *echo.Echo) {
	server.GET("/users", resorce.getUsers)
	server.GET("/users/:id", resorce.getUser)
	server.POST("/users", resorce.addUser)
	server.DELETE("/users/:id", resorce.deleteUser)
	server.PUT("/users/:id", resorce.updateUser)
}

func (a *UserResource) addUser(context echo.Context) error {
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

func (a *UserResource) getUsers(context echo.Context) error {
	users, err := a.UserDAO.GetUsers()

	if err != nil {
		log.Println("Error retrieving users:", err.Error())
		return context.String(http.StatusInternalServerError, "Error retrieving users")
	}

	return context.JSON(http.StatusOK, users)
}

func (a *UserResource) getUser(context echo.Context) error {
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

func (a *UserResource) deleteUser(context echo.Context) error {
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

func (a *UserResource) updateUser(context echo.Context) error {
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
