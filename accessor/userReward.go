package accessor

type UserReward interface {
	InsertUserReward(userRewardDb UserRewardDb) error
	UpdateUserReward(userRewardDb UserRewardDb) error
	GetUserReward(id string, rewardID string, collegeID string) ([]UserRewardDb, error)
	DeleteUserReward(id string, rewardID string, collegeID string) error
	ListUserReward(input ListUserRewardInput) ([]string, error)
}

type UserRewardDb struct {
	ID        string `json:"user_id"`
	CollegeID string `json:"college_id"`
	RewardID  string `json:"reward_id"`
	Quantity  int    `json:"quantity"`
}

type ListUserRewardInput struct {
	StudentID string `json:"studentId"`
	CollegeID string `json:"collegeId"`
}

func (a *accessor) GetUserReward(id string, rewardID string, collegeID string) ([]UserRewardDb, error) {
	getQuery := `SELECT * FROM "public".user_reward WHERE user_id = $1 AND college_id = $2 AND reward_id = $3`

	queryInput := []interface{}{
		id,
		collegeID,
		rewardID,
	}

	var userRewardDb []UserRewardDb
	if err := a.sqlAdapter.Read(getQuery, queryInput, &userRewardDb); err != nil {
		return nil, err
	}

	return userRewardDb, nil
}

func (a *accessor) InsertUserReward(userRewardDb UserRewardDb) error {
	upsertQuery := `
		INSERT INTO "public".user_reward(user_id, college_id, reward_id, quantity)
        VALUES($1, $2, $3, $4)
	`

	queryInput := []interface{}{
		userRewardDb.ID,
		userRewardDb.CollegeID,
		userRewardDb.RewardID,
		userRewardDb.Quantity,
	}

	return a.sqlAdapter.Write(upsertQuery, queryInput)
}

func (a *accessor) UpdateUserReward(userRewardDb UserRewardDb) error {
	upsertQuery := `
		UPDATE "public".user_reward
		SET quantity = $1
		WHERE user_id = $2 AND college_id = $3 AND reward_id = $4
	`

	queryInput := []interface{}{
		userRewardDb.Quantity,
		userRewardDb.ID,
		userRewardDb.CollegeID,
		userRewardDb.RewardID,
	}

	return a.sqlAdapter.Write(upsertQuery, queryInput)
}

func (a *accessor) DeleteUserReward(id string, rewardID string, collegeID string) error {
	insertQuery := `DELETE FROM "public".user_reward WHERE user_id = $1 AND college_id = $2 AND reward_id = $3`

	queryInput := []interface{}{
		id,
		collegeID,
		rewardID,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}

func (a *accessor) ListUserReward(input ListUserRewardInput) ([]string, error) {
	queryInput := []interface{}{
		input.StudentID,
		input.CollegeID,
	}

	getQuery := `SELECT reward_id FROM "public".user_reward WHERE user_id = $1 AND college_id = $2`

	var rewardIds []string
	if err := a.sqlAdapter.Read(getQuery, queryInput, &rewardIds); err != nil {
		return nil, err
	}

	return rewardIds, nil
}
