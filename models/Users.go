package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

//Users contains user details
type Users struct {
	UserID     int    `json:"ID"`
	Firstname  string `json:"Firstname"`
	Lastname   string `json:"Lastname"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	IsParent   bool
	IsFaculty  bool
	IsActive   bool
	LastLogin  string
	DateJoined string
}

// Students contains details of students
type Students struct {
	StudentID int    `json:"StudentID"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"lname"`
	Rollno    int
	Class     int `json:"class"`
	UserID    int `json:"uid"`
}

// Faculty contains details of faculty
type Faculty struct {
	FacultyID int
	ExamID    int
	UserID    int
}

//Parents consists of details of parents
type Parents struct {
	ParentID  int
	StudentID int
	UserID    int
}

//AddUser connects to db
func (user *Users) AddUser(conn *pgx.Conn) error {
	var err error
	CreateUsersTable := `CREATE TABLE IF NOT EXISTS users
(
	UserID     SERIAL PRIMARY KEY,  
	Firstname  TEXT,
	Lastname   TEXT,
	Email      TEXT NOT NULL UNIQUE,
	Password   TEXT,
	IsParent   BOOLEAN DEFAULT FALSE,
	IsFaculty  BOOLEAN DEFAULT FALSE,
	IsActive   BOOLEAN DEFAULT FALSE,
	DateJoined DATE DEFAULT CURRENT_DATE
)`
	_, err = conn.Exec(context.Background(), CreateUsersTable)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec(context.Background(), "insert into users (Firstname, Lastname, Email, Password) values($1,$2,$3,crypt($4, gen_salt('bf')))", user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}
