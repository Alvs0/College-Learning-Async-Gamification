package accessor

import (
	"fmt"
	"strings"
)

type Session interface {
	UpsertSession(sessionDb SessionDb) error
	ListSession(input SessionListInput) ([]string, error)
	GetSessions(ids []string) ([]SessionDb, error)
	DeleteSession(id string) error
}

type SessionDb struct {
	ID        string `json:"id"`
	CollegeID string `json:"college_id"`
	Name      string `json:"name"`
	Link      string `json:"link"`
	ImageUrl  string `json:"image_url"`
}

type SessionListInput struct {
	CollegeID string `json:"collegeId"`
}

func (a *accessor) ListSession(input SessionListInput) ([]string, error) {
	queryInput := []interface{}{
		input.CollegeID,
	}

	getQuery := `SELECT id FROM "public".session WHERE college_id = $1`

	var sessionIds []string
	if err := a.sqlAdapter.Read(getQuery, queryInput, &sessionIds); err != nil {
		return nil, err
	}

	return sessionIds, nil
}

func (a *accessor) GetSessions(ids []string) ([]SessionDb, error) {
	getQuery := fmt.Sprintf(`SELECT * FROM "public".session WHERE id IN ('%s')`, strings.Join(ids, "', '"))

	var sessionDbs []SessionDb
	if err := a.sqlAdapter.Read(getQuery, nil, &sessionDbs); err != nil {
		return nil, err
	}

	return sessionDbs, nil
}

func (a *accessor) UpsertSession(sessionDb SessionDb) error {
	insertQuery := `INSERT INTO "public".session(id, college_id, name, link, image_url) VALUES($1, $2, $3, $4, $5)`

	queryInput := []interface{}{
		sessionDb.ID,
		sessionDb.CollegeID,
		sessionDb.Name,
		sessionDb.Link,
		sessionDb.ImageUrl,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}

func (a *accessor) DeleteSession(id string) error {
	insertQuery := `DELETE FROM "public".session WHERE id = $1`

	queryInput := []interface{}{
		id,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}
