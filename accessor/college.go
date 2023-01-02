package accessor

import (
	"fmt"
	"strings"
)

type College interface {
	UpsertCollege(collegeDb CollegeDb) error
	GetCollegeByID(id string) (CollegeDb, error)
	GetCollegeByName(name string) (CollegeDb, error)
	GetCollegeByIDs(ids []string) ([]CollegeDb, error)
	ListCollege() ([]string, error)
	DeleteCollege(id string) error
}

type CollegeDb struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	IconURL string `json:"icon_url" db:"icon_url"`
}

func (a *accessor) UpsertCollege(collegeDb CollegeDb) error {
	insertQuery := `INSERT INTO "public".college(id, name, address, icon_url) VALUES($1, $2, $3, $4)`

	queryInput := []interface{}{
		collegeDb.ID,
		collegeDb.Name,
		collegeDb.Address,
		collegeDb.IconURL,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}

func (a *accessor) GetCollegeByID(id string) (CollegeDb, error) {
	getQuery := `SELECT * FROM "public".college WHERE id = $1`

	queryInput := []interface{}{
		id,
	}

	var collegeDb CollegeDb
	if err := a.sqlAdapter.Read(getQuery, queryInput, &collegeDb); err != nil {
		return CollegeDb{}, err
	}

	return collegeDb, nil
}

func (a *accessor) GetCollegeByName(name string) (CollegeDb, error) {
	getQuery := `SELECT * FROM "public".college WHERE name = $1`

	queryInput := []interface{}{
		name,
	}

	var collegeDb CollegeDb
	if err := a.sqlAdapter.Read(getQuery, queryInput, &collegeDb); err != nil {
		return CollegeDb{}, err
	}

	return collegeDb, nil
}

func (a *accessor) GetCollegeByIDs(ids []string) ([]CollegeDb, error) {
	getQuery := fmt.Sprintf(`SELECT * FROM "public".college WHERE id IN ('%s')`, strings.Join(ids, "', '"))

	var collegeDbs []CollegeDb
	if err := a.sqlAdapter.Read(getQuery, nil, &collegeDbs); err != nil {
		return nil, err
	}

	return collegeDbs, nil
}

func (a *accessor) ListCollege() ([]string, error) {
	getQuery := `SELECT id FROM "public".college`

	queryInput := []interface{}{}

	var collegeIds []string
	if err := a.sqlAdapter.Read(getQuery, queryInput, &collegeIds); err != nil {
		return nil, err
	}

	return collegeIds, nil
}

func (a *accessor) DeleteCollege(id string) error {
	insertQuery := `DELETE FROM "public".college WHERE id = $1`

	queryInput := []interface{}{
		id,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}
