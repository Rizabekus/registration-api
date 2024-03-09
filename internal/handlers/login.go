package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rizabekus/registration-api/internal/models"
	"github.com/Rizabekus/registration-api/pkg/loggers"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
)

func (handler *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to Login")
	var LoginInstance models.LoginUser
	err := json.NewDecoder(r.Body).Decode(&LoginInstance)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to decode JSON",
			Error: "",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("Failed to decode JSON")
		return
	}
	loggers.DebugLog.Println("Received data in JSON format")
	validate := validator.New()

	err = validate.Struct(LoginInstance)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			response := models.ResponseStructure{
				Field: "Internal Server Error",
				Error: "",
			}
			handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

			loggers.InfoLog.Println("Internal Server Error")
			return
		}
		firstValidationError := validationErrors[0]
		response := models.ResponseStructure{
			Field: fmt.Sprintf("Field: %s, Tag: %s\n", firstValidationError.Field(), firstValidationError.Tag()),
			Error: err.Error(),
		}

		handler.Service.UserService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("Validation Error")
		return
	}
	loggers.DebugLog.Println("Validated the data")

	ok, err := handler.Service.UserService.CheckCredentials(LoginInstance)
	if !ok {
		response := models.ResponseStructure{
			Field: "No user with such data",
			Error: "Wrong email/password",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.DebugLog.Println("Credentials are not found in database")
		return
	}
	if err != nil {

		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)
		fmt.Println("HELLO mazafaka")
		loggers.InfoLog.Println("Internal Server Error:", err)
		return
	}

	u2, err := uuid.NewV4()
	if err != nil {

		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to create UUID: ", err)
		return
	}
	loggers.DebugLog.Println("Created UUId instance")
	UserData, err := handler.Service.UserService.GetUserByEmail(LoginInstance.Email)

	if err != nil {

		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Internal Server Error: ", err)
		return
	}
	loggers.DebugLog.Println("Got user data")
	err = handler.Service.UserService.CreateSession(UserData.ID, u2.String())
	if err != nil {
		fmt.Println("HERE3")
		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Internal Server Error: ", err)
		return
	}
	loggers.DebugLog.Println("Added session into table")
	cookie := &http.Cookie{Name: "logged-in", Value: u2.String(), Expires: time.Now().Add(365 * 24 * time.Hour)}
	http.SetCookie(w, cookie)
	loggers.DebugLog.Println("Session is set")
	loggers.DebugLog.Println("User successfully logged-in")

}
