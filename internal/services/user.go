package services

import (
	"encoding/json"
	"net/http"

	"github.com/Rizabekus/registration-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	storage models.UserStorage
}

func CreateUserService(storage models.UserStorage) *UserService {
	return &UserService{storage: storage}
}
func (us *UserService) SendResponse(response models.ResponseStructure, w http.ResponseWriter, statusCode int) {
	responseJSON, err := json.Marshal(response)
	if err != nil {

		internalError := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "Failed to marshal JSON response",
		}
		internalErrorJSON, _ := json.Marshal(internalError)

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(internalErrorJSON), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
func (us *UserService) AddUser(UserData models.RegisterUser) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(UserData.Password), 1)
	if err != nil {
		return err
	}
	UserData.Password = string(pwd)
	return us.storage.AddUser(UserData)
}
func (us *UserService) CheckUserExistence(email string) (bool, error) {
	return us.storage.CheckUserExistence(email)
}

func (us *UserService) CheckCredentials(LoginInstance models.LoginUser) (bool, error) {
	UserData, err := us.storage.GetUserByEmail(LoginInstance.Email)
	if err != nil {
		return false, err
	}
	if bcrypt.CompareHashAndPassword([]byte(UserData.Password), []byte(LoginInstance.Password)) == nil {
		return true, nil
	} else {
		return false, nil
	}
}
func (us *UserService) GetUserByEmail(email string) (models.User, error) {
	return us.storage.GetUserByEmail(email)
}
func (us *UserService) CreateSession(id int, uuid string) error {
	return us.storage.CreateSession(id, uuid)
}
func (us *UserService) GetID(cookie string) (int, error) {
	return us.storage.GetID(cookie)
}

func (us *UserService) GetUserDataByID(user_id int) (models.User, error) {
	return us.storage.GetUserDataByID(user_id)
}
func (us *UserService) UpdateUser(user_id int, modifications models.ModifyUser) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(modifications.Password), 1)
	if err != nil {
		return err
	}
	modifications.Password = string(pwd)
	return us.storage.UpdateUser(user_id, modifications)
}
