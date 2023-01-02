package reward

func (c *service) GetRewardByIDs(req GetRewardByIDsReq, res *GetRewardByIDsRes) error {
	rewardDbs, err := c.accessor.GetRewards(req.RewardIDs)
	if err != nil {
		return err
	}

	var rewardObjs []RewardObj
	for _, rewardDb := range rewardDbs {
		rewardObjs = append(rewardObjs, RewardObj{
			ID:             rewardDb.ID,
			CollegeID:      rewardDb.CollegeID,
			Name:           rewardDb.Name,
			ImageUrl:       rewardDb.ImageUrl,
			Description:    rewardDb.Description,
			Quantity:       rewardDb.Quantity,
			RequiredPoints: rewardDb.RequiredPoints,
			MinimalLevel:   rewardDb.MinimalLevel,
			IsActive:       rewardDb.IsActive,
		})
	}

	*res = GetRewardByIDsRes{
		Rewards: rewardObjs,
	}

	return nil
}
