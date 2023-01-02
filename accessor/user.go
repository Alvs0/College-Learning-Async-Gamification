package accessor

import (
	"fmt"
	"strings"
)

type User interface {
	UpsertUser(upsertInput UserUpsertInput) error
	ListUser(input UserListInput) ([]string, error)
	GetUsers(ids []string) ([]UserDb, error)
	GetUserByEmail(email string) ([]UserDb, error)
	DeleteUser(id string) error
}

type UserDb struct {
	ID              string `json:"id" db:"id"`
	CollegeID       string `json:"college_id" db:"college_id"`
	Name            string `json:"name" db:"name"`
	Email           string `json:"email" db:"email"`
	PhoneNumber     string `json:"phone_number" db:"phone_number"`
	BirthDate       string `json:"birth_date" db:"birth_date"`
	ProfileImageUrl string `json:"profile_image_url" db:"profile_image_url"`
	Password        string `json:"password" db:"password"`
	IsAdmin         bool   `json:"is_admin" db:"is_admin"`
}

type UserUpsertInput struct {
	ID              string `json:"id"`
	CollegeID       string `json:"collegeId"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phoneNumber"`
	BirthDate       string `json:"birthDate"`
	ProfileImageUrl string `json:"profileImageUrl"`
	Password        string `json:"password"`
	IsAdmin         bool   `json:"isAdmin"`
}

type UserListInput struct {
	CollegeID string `json:"collegeId"`
	IsAdmin   bool   `json:"isAdmin"`
}

func (a *accessor) UpsertUser(upsertInput UserUpsertInput) error {
	upsertQuery := `
		INSERT INTO "public".user(id, college_id, name, email, phone_number, birth_date, profile_image_url, password, is_admin)
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	queryInput := []interface{}{
		upsertInput.ID,
		upsertInput.CollegeID,
		upsertInput.Name,
		upsertInput.Email,
		upsertInput.PhoneNumber,
		upsertInput.BirthDate,
		upsertInput.ProfileImageUrl,
		upsertInput.Password,
		upsertInput.IsAdmin,
	}

	return a.sqlAdapter.Write(upsertQuery, queryInput)
}

type UserDeleteInput struct {
	ID string `json:"id"`
}

func (a *accessor) DeleteUser(id string) error {
	insertQuery := `DELETE FROM "public".user WHERE id = $1`

	queryInput := []interface{}{
		id,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}

func (a *accessor) GetUsers(ids []string) ([]UserDb, error) {
	getQuery := fmt.Sprintf(`SELECT * FROM "public".user WHERE id IN ('%s')`, strings.Join(ids, "', '"))

	var userDbs []UserDb
	if err := a.sqlAdapter.Read(getQuery, nil, &userDbs); err != nil {
		return nil, err
	}

	return userDbs, nil
}

func (a *accessor) GetUserByEmail(email string) ([]UserDb, error) {
	getQuery := `SELECT * FROM "public".user WHERE email = $1`

	queryInput := []interface{}{
		email,
	}

	var userDbs []UserDb
	if err := a.sqlAdapter.Read(getQuery, queryInput, &userDbs); err != nil {
		return nil, err
	}

	return userDbs, nil
}

func (a *accessor) ListUser(input UserListInput) ([]string, error) {
	queryInput := []interface{}{
		input.IsAdmin,
	}

	var getQuery string
	if input.CollegeID == "" {
		getQuery = `SELECT id FROM "public".user WHERE is_admin = $1`
	} else {
		getQuery = `SELECT id FROM "public".user WHERE is_admin = $1 AND college_id = $2`
		queryInput = append(queryInput, input.CollegeID)
	}

	var userDbs []string
	if err := a.sqlAdapter.Read(getQuery, queryInput, &userDbs); err != nil {
		return nil, err
	}

	return userDbs, nil
}
