package storage

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/Rizabekus/registration-api/internal/models"
	errortypes "github.com/Rizabekus/registration-api/pkg/errors"
)

type UserDB struct {
	DB *sql.DB
}

func CreateUserStorage(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}
func (udb *UserDB) CheckUserExistence(email string) (bool, error) {
	fmt.Println("email: ", email)
	fmt.Println("HELLLO")
	query := "SELECT EXISTS(SELECT * FROM users WHERE email = $1)"
	fmt.Println("HELLLO")
	var exists bool
	err := udb.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking existence: %v", err)
	}

	return exists, nil
}
func (udb *UserDB) AddUser(UserData models.RegisterUser) error {
	sqlStatement := `
		INSERT INTO users (email, hashed_password, usertype)
		VALUES ($1, $2, $3)
	`

	_, err := udb.DB.Exec(sqlStatement, UserData.Email, UserData.Password, "user")
	if err != nil {

		return fmt.Errorf("error executing SQL statement in AddPerson: %v", err)
	}

	return nil
}
func (udb *UserDB) GetUserByEmail(email string) (models.User, error) {
	var UserData models.User

	row := udb.DB.QueryRow("SELECT id, name, email, mobile_number, date_of_birth, hashed_password, usertype FROM users WHERE email = $1", email)

	err := row.Scan(&UserData.ID, &UserData.Name, &UserData.Email, &UserData.Mobile_number, &UserData.Date_of_birth, &UserData.Password, &UserData.Usertype)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserData, fmt.Errorf("person not found with Email %s", email)
		}
		return UserData, fmt.Errorf("failed to scan person data: %w", err)
	}

	return UserData, nil
}
func (udb *UserDB) CreateSession(id int, uuid string) error {
	stmt, err := udb.DB.Prepare("INSERT INTO session_cookies(user_id, session) VALUES($1, $2)")
	if err != nil {
		fmt.Println("qweqwe")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, uuid)
	if err != nil {

		return err
	}
	return nil
}

func (udb *UserDB) GetID(cookie string) (int, error) {
	st, err := udb.DB.Query("SELECT user_id FROM session_cookies WHERE session=($1)", cookie)
	if err != nil {
		fmt.Println("Here")
		return 0, err
	}
	defer st.Close()

	var user_id int
	found := false
	for st.Next() {
		found = true
		if err := st.Scan(&user_id); err != nil {
			fmt.Println("Here1")
			return 0, err
		}
	}
	if !found {
		fmt.Println("Here3")
		return 0, errortypes.ErrNoUserID
	}
	return user_id, nil
}

func (udb *UserDB) GetUserDataByID(user_id int) (models.User, error) {
	var UserData models.User

	row := udb.DB.QueryRow("SELECT id, name, email, mobile_number, date_of_birth, hashed_password, usertype FROM users WHERE id = $1", user_id)

	err := row.Scan(&UserData.ID, &UserData.Name, &UserData.Email, &UserData.Mobile_number, &UserData.Date_of_birth, &UserData.Password, &UserData.Usertype)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserData, fmt.Errorf("person not found with Email %d", user_id)
		}
		return UserData, fmt.Errorf("failed to scan person data: %w", err)
	}

	return UserData, nil
}

func (udb *UserDB) UpdateUser(userID int, userData models.ModifyUser) error {

	query := "UPDATE users SET "
	args := []interface{}{}
	var setStatements []string
	counter := 1
	if userData.Name != "" {

		setStatements = append(setStatements, "name = $"+strconv.Itoa(counter))
		counter++
		args = append(args, userData.Name)
	}

	if userData.Email != "" {
		setStatements = append(setStatements, "email = $"+strconv.Itoa(counter))
		counter++
		args = append(args, userData.Email)
	}

	if userData.Mobile_number != "" {
		setStatements = append(setStatements, "mobile_number = $"+strconv.Itoa(counter))
		counter++
		args = append(args, userData.Mobile_number)
	}

	if !userData.Date_of_birth.IsZero() {
		setStatements = append(setStatements, "date_of_birth = $"+strconv.Itoa(counter))
		counter++
		args = append(args, userData.Date_of_birth)
	}
	if userData.Password != "" {
		setStatements = append(setStatements, "hashed_password = $"+strconv.Itoa(counter))
		counter++
		args = append(args, userData.Password)
	}

	query += fmt.Sprintf("%s WHERE id = $"+strconv.Itoa(counter), strings.Join(setStatements, ", "))

	args = append(args, userID)

	_, err := udb.DB.Exec(query, args...)
	if err != nil {
		fmt.Println("Here4")
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
