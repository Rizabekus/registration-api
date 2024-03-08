package models

import (
	"database/sql"
	"net/http"
	"time"
)

type UserService interface {
	AddUser(UserData RegisterUser) error
	CheckUserExistence(email string) (bool, error)
	SendResponse(response ResponseStructure, w http.ResponseWriter, statusCode int)
	CheckCredentials(LoginInstance LoginUser) (bool, error)
	GetUserByEmail(email string) (User, error)
	CreateSession(id int, uuid string) error
	GetID(cookie string) (int, error)
	GetUserDataByID(user_id int) (User, error)
	UpdateUser(userID int, userData ModifyUser) error
}
type UserStorage interface {
	AddUser(UserData RegisterUser) error
	CheckUserExistence(email string) (bool, error)
	GetUserByEmail(email string) (User, error)
	CreateSession(id int, uuid string) error
	GetID(cookie string) (int, error)
	GetUserDataByID(user_id int) (User, error)
	UpdateUser(userID int, userData ModifyUser) error
}

type User struct {
	ID                int
	Name              sql.NullString
	Email             string
	Mobile_number     sql.NullString
	Date_of_birth     sql.NullTime
	Password          string
	Repeated_password string
	Usertype          string
}
type ModifyUser struct {
	Name          string    `json:"name,omitempty" validate:"omitempty,cyrillic_or_latin_validator,min=3,max=50"`
	Email         string    `json:"email,omitempty" validate:"omitempty,email"`
	Mobile_number string    `json:"mobile_number,omitempty" validate:"omitempty,phone_validator"`
	Date_of_birth time.Time `json:"date_of_birth,omitempty" time_format:"2006-01-02" validate:"omitempty,date_of_birth_validator"`
	Password      string    `json:"password,omitempty" validate:"omitempty,ascii_validator"`
	// Repeated_password string    `json:"repeated_password,omitempty" validate:"omitempty,eqfield=Password"`
}

type RegisterUser struct {
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required,min=9,max=50"`
	Repeated_password string `json:"repeated_password" validate:"required,min=9,max=50,eqfield=Password"`
}
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ResponseStructure struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
