package reward

import (
	"college-learning-asynchronous-gamification/accessor"
)

type Service interface {
	UpsertReward(req UpsertRewardReq, res *UpsertRewardRes) error
	ListReward(req ListRewardReq, res *ListRewardRes) error
	GetRewardByIDs(req GetRewardByIDsReq, res *GetRewardByIDsRes) error
	DeleteReward(req DeleteRewardReq, res *DeleteRewardRes) error
	ClaimReward(req ClaimRewardReq, res *ClaimRewardRes) error
	UseReward(req UseRewardReq, res *UseRewardRes) error
}

type service struct {
	accessor accessor.Accessor
}

func NewService(accessor accessor.Accessor) Service {
	return &service{
		accessor: accessor,
	}
}

type ListRewardReq struct {
	CollegeID string `json:"collegeId"`
}

type ListRewardRes struct {
	RewardIDs []string `json:"rewardIds"`
}

type GetRewardByIDsReq struct {
	RewardIDs []string `json:"rewardIds"`
}

type GetRewardByIDsRes struct {
	Rewards []RewardObj `json:"rewards"`
}

type UpsertRewardReq struct {
	Reward RewardObj `json:"rewards"`
}

type UpsertRewardRes struct{}

type DeleteRewardReq struct {
	RewardID string `json:"rewardId"`
}

type DeleteRewardRes struct{}

type ClaimRewardReq struct {
	RewardID  string `json:"rewardId"`
	StudentID string `json:"studentId"`
	Quantity  int    `json:"quantity"`
}

type ClaimRewardRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UseRewardReq struct {
	RewardID  string `json:"rewardId"`
	StudentID string `json:"studentId"`
}

type UseRewardRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
