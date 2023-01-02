package accessor

type UserPoint interface {
	InsertUserPoint(userPointDb UserPointDb) error
	UpdateUserPoint(userPointDb UserPointDb) error
	GetUserPointByUserID(id string) (UserPointDb, error)
}

type UserPointDb struct {
	UserID    string `json:"user_id"`
	CollegeID string `json:"college_id"`
	Point     int    `json:"point"`
}

func (a *accessor) GetUserPointByUserID(id string) (UserPointDb, error) {
	getQuery := `SELECT * FROM "public".user_point WHERE user_id = $1`

	queryInput := []interface{}{
		id,
	}

	var userPointDb UserPointDb
	if err := a.sqlAdapter.Read(getQuery, queryInput, &userPointDb); err != nil {
		return UserPointDb{}, err
	}

	return userPointDb, nil
}

func (a *accessor) InsertUserPoint(userPointDb UserPointDb) error {
	upsertQuery := `
		INSERT INTO "public".user_point(user_id, college_id, point)
        VALUES($1, $2, $3)
	`

	queryInput := []interface{}{
		userPointDb.UserID,
		userPointDb.CollegeID,
		userPointDb.Point,
	}

	return a.sqlAdapter.Write(upsertQuery, queryInput)
}

func (a *accessor) UpdateUserPoint(userPointDb UserPointDb) error {
	upsertQuery := `
		UPDATE "public".user_point
		SET point = $1
		WHERE user_id = $2 AND college_id = $3
	`

	queryInput := []interface{}{
		userPointDb.Point,
		userPointDb.UserID,
		userPointDb.CollegeID,
	}

	return a.sqlAdapter.Write(upsertQuery, queryInput)
}