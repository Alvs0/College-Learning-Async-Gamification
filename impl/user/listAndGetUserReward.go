package user

import (
	"college-learning-asynchronous-gamification/accessor"
)

func (s *service) ListAndGetUserReward(req ListAndGetUserRewardReq, res *ListAndGetUserRewardRes) error {
	rewardIds, err := s.accessor.ListUserReward(accessor.ListUserRewardInput{
		CollegeID: req.CollegeID,
		StudentID: req.UserID,
	})
	if err != nil {
		return err
	}

	rewardDbs, err := s.accessor.GetRewards(rewardIds)
	if err != nil {
		return err
	}

	rewardMap := make(map[string]RewardObj)
	for _, rewardDb := range rewardDbs {
		if _, ok := rewardMap[rewardDb.ID]; ok {
			continue
		}

		rewardMap[rewardDb.ID] = RewardObj{
			ID:             rewardDb.ID,
			CollegeID:      rewardDb.CollegeID,
			Name:           rewardDb.Name,
			ImageUrl:       rewardDb.ImageUrl,
			Description:    rewardDb.Description,
			Quantity:       rewardDb.Quantity,
			RequiredPoints: rewardDb.RequiredPoints,
			MinimalLevel:   rewardDb.MinimalLevel,
			IsActive:       rewardDb.IsActive,
		}
	}

	var rewards []RewardObj
	for _, rewardId := range rewardIds {
		userReward, err := s.accessor.GetUserReward(req.UserID, rewardId, req.CollegeID)
		if err != nil {
			continue
		}

		if _, ok := rewardMap[rewardId]; !ok {
			continue
		}

		rewards = append(rewards, RewardObj{
			ID:             rewardMap[rewardId].ID,
			CollegeID:      rewardMap[rewardId].CollegeID,
			Name:           rewardMap[rewardId].Name,
			ImageUrl:       rewardMap[rewardId].ImageUrl,
			Description:    rewardMap[rewardId].Description,
			Quantity:       userReward[0].Quantity,
			RequiredPoints: rewardMap[rewardId].RequiredPoints,
			MinimalLevel:   rewardMap[rewardId].MinimalLevel,
			IsActive:       rewardMap[rewardId].IsActive,
		})
	}

	useReward := UserReward{
		UserID:    req.UserID,
		CollegeID: req.CollegeID,
		Rewards:   rewards,
	}

	*res = ListAndGetUserRewardRes{
		UserReward: useReward,
	}

	return nil
}
