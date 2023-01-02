package reward

import "college-learning-asynchronous-gamification/accessor"

func (s *service) ClaimReward(req ClaimRewardReq, res *ClaimRewardRes) error {
	userPoint, err := s.accessor.GetUserPointByUserID(req.StudentID)
	if err != nil {
		return err
	}

	rewardDbs, err := s.accessor.GetRewards([]string{req.RewardID})
	if err != nil {
		return err
	}

	rewardDb := rewardDbs[0]

	totalPoints := rewardDb.RequiredPoints * req.Quantity
	pointAfter := userPoint.Point - totalPoints
	if pointAfter < 0 {
		*res = ClaimRewardRes{
			Success: false,
			Message: "You don't have enough points!",
		}

		return nil
	}

	afterQuantity := rewardDb.Quantity - req.Quantity
	if afterQuantity < 0 {
		*res = ClaimRewardRes{
			Success: false,
			Message: "You cannot claim more than rewards quantity",
		}

		return nil
	}

	if afterQuantity > 0 {
		err = s.accessor.UpdateRewardQuantity(accessor.RewardDb{
			ID:             rewardDb.ID,
			CollegeID:      rewardDb.CollegeID,
			Name:           rewardDb.Name,
			Description:    rewardDb.Description,
			ImageUrl:       rewardDb.ImageUrl,
			Quantity:       afterQuantity,
			RequiredPoints: rewardDb.RequiredPoints,
			MinimalLevel:   rewardDb.MinimalLevel,
			IsActive:       rewardDb.IsActive,
		})
	} else {
		if err := s.accessor.DeleteReward(rewardDb.ID); err != nil {
			return nil
		}
	}

	err = s.accessor.UpdateUserPoint(accessor.UserPointDb{
		UserID:    req.StudentID,
		CollegeID: userPoint.CollegeID,
		Point:     pointAfter,
	})
	if err != nil {
		return err
	}

	userRewardDbs, err := s.accessor.GetUserReward(req.StudentID, req.RewardID, userPoint.CollegeID)
	if err != nil {
		return err
	}

	if userRewardDbs == nil {
		if err := s.accessor.InsertUserReward(accessor.UserRewardDb{
			ID:        req.StudentID,
			CollegeID: userPoint.CollegeID,
			RewardID:  req.RewardID,
			Quantity:  req.Quantity,
		}); err != nil {
			return err
		}
	} else {
		if err := s.accessor.UpdateUserReward(accessor.UserRewardDb{
			ID:        req.StudentID,
			CollegeID: userPoint.CollegeID,
			RewardID:  req.RewardID,
			Quantity:  userRewardDbs[0].Quantity + req.Quantity,
		}); err != nil {
			return err
		}
	}

	*res = ClaimRewardRes{
		Success: true,
		Message: "",
	}

	return nil
}
