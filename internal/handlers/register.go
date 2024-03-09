package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rizabekus/registration-api/internal/models"
	"github.com/Rizabekus/registration-api/pkg/loggers"
	"github.com/go-playground/validator"
)

func (handler *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to register a user")
	var newUser models.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
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
	err = validate.Struct(newUser)

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
			Error: "",
		}

		handler.Service.UserService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("Validation Error", err)
		return
	}
	loggers.DebugLog.Println("Validated the data", newUser)
	exists, err := handler.Service.UserService.CheckUserExistence(newUser.Email)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to check user existence: ", err)
		return
	}

	if exists {
		response := models.ResponseStructure{
			Field: "User already exists",
			Error: "Email is used",
		}

		handler.Service.UserService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("User does exist")
		return
	}
	loggers.DebugLog.Println("User does not exist")
	err = handler.Service.UserService.AddUser(newUser)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to add user: ", err)
		return
	}
	loggers.InfoLog.Println("Successfully registered a user")
}
