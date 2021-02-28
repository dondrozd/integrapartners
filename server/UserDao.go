package main

import (
	"database/sql"
	"log"
)

type UserDAO struct {
	DB *sql.DB
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