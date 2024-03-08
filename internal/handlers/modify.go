package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rizabekus/registration-api/internal/models"
	errortypes "github.com/Rizabekus/registration-api/pkg/errors"
	"github.com/Rizabekus/registration-api/pkg/loggers"
	"github.com/Rizabekus/registration-api/pkg/validators"
	"github.com/go-playground/validator"
)

func (handler *Handlers) Modify(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("logged-in")
	if err != nil {
		if err == http.ErrNoCookie {
			response := models.ResponseStructure{
				Field: "You are not logged in",
				Error: "No permission to modify",
			}
			handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)
			return

		} else {
			response := models.ResponseStructure{
				Field: "Internal Server Error",
				Error: err.Error(),
			}
			handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

			return
		}
	}
	user_id, err := handler.Service.UserService.GetID(cookie.Value)
	if err != nil {
		if err == errortypes.ErrNoUserID {
			response := models.ResponseStructure{
				Field: "You are logged in not properly",
				Error: "No permission to modify",
			}
			handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)
			return
		} else {
			response := models.ResponseStructure{
				Field: "Internal Server Error",
				Error: err.Error(),
			}
			handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

			return
		}
	}
	var userModifications models.ModifyUser
	err = json.NewDecoder(r.Body).Decode(&userModifications)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to decode JSON",
			Error: err.Error(),
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("Failed to decode JSON")
		return
	}
	loggers.DebugLog.Println("Received data in JSON format")
	validate := validator.New()
	validate.RegisterValidation("ascii_validator", validators.ASCIIValidator)
	validate.RegisterValidation("date_of_birth_validator", validators.DateOfBirthValidator)
	validate.RegisterValidation("phone_validator", validators.PhoneValidator)
	validate.RegisterValidation("cyrillic_or_latin_validator", validators.CyrillicOrLatinValidator)
	err = validate.Struct(userModifications)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			response := models.ResponseStructure{
				Field: "Internal Server Error",
				Error: err.Error(),
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
	// UserData, err := handler.Service.UserService.GetUserDataByID(user_id)
	err = handler.Service.UserService.UpdateUser(user_id, userModifications)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: err.Error(),
		}
		handler.Service.UserService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Internal Server Error")
		return
	}
}
