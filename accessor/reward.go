package accessor

import (
	"fmt"
	"strings"
)

type Reward interface {
	UpsertReward(rewardDb RewardDb) error
	ListReward(input RewardListInput) ([]string, error)
	GetRewards(ids []string) ([]RewardDb, error)
	DeleteReward(id string) error
	UpdateRewardQuantity(rewardDb RewardDb) error
}

type RewardDb struct {
	ID             string `json:"id"`
	CollegeID      string `json:"college_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ImageUrl       string `json:"image_url"`
	Quantity       int    `json:"quantity"`
	RequiredPoints int    `json:"required_points"`
	MinimalLevel   int    `json:"minimal_level"`
	IsActive       bool   `json:"is_active"`
}

type RewardListInput struct {
	CollegeID string `json:"collegeId"`
}

func (a *accessor) ListReward(input RewardListInput) ([]string, error) {
	queryInput := []interface{}{
		input.CollegeID,
	}

	getQuery := `SELECT id FROM "public".reward WHERE college_id = $1`

	var rewardIds []string
	if err := a.sqlAdapter.Read(getQuery, queryInput, &rewardIds); err != nil {
		return nil, err
	}

	return rewardIds, nil
}

func (a *accessor) GetRewards(ids []string) ([]RewardDb, error) {
	getQuery := fmt.Sprintf(`SELECT * FROM "public".reward WHERE id IN ('%s')`, strings.Join(ids, "', '"))

	var rewardDbs []RewardDb
	if err := a.sqlAdapter.Read(getQuery, nil, &rewardDbs); err != nil {
		return nil, err
	}

	return rewardDbs, nil
}

func (a *accessor) UpsertReward(rewardDb RewardDb) error {
	insertQuery := `INSERT INTO "public".reward(id, college_id, name, description, quantity, minimal_level, required_points, image_url, is_active) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	queryInput := []interface{}{
		rewardDb.ID,
		rewardDb.CollegeID,
		rewardDb.Name,
		rewardDb.Description,
		rewardDb.Quantity,
		rewardDb.MinimalLevel,
		rewardDb.RequiredPoints,
		rewardDb.ImageUrl,
		rewardDb.IsActive,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}

func (a *accessor) UpdateRewardQuantity(rewardDb RewardDb) error {
	insertQuery := `UPDATE "public".reward SET quantity = $1 WHERE id = $2 AND college_id = $3`

	queryInput := []interface{}{
		rewardDb.Quantity,
		rewardDb.ID,
		rewardDb.CollegeID,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}

func (a *accessor) DeleteReward(id string) error {
	insertQuery := `DELETE FROM "public".reward WHERE id = $1`

	queryInput := []interface{}{
		id,
	}

	if err := a.sqlAdapter.Write(insertQuery, queryInput); err != nil {
		return err
	}

	return nil
}
