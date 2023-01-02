package reward

import "college-learning-asynchronous-gamification/accessor"

func (s *service) UseReward(req UseRewardReq, res *UseRewardRes) error {
	userDbs, err := s.accessor.GetUsers([]string{req.StudentID})
	if err != nil {
		return err
	}

	userDb := userDbs[0]

	userRewardDbs, err := s.accessor.GetUserReward(req.StudentID, req.RewardID, userDb.CollegeID)
	if err != nil {
		return err
	}

	if userRewardDbs == nil {
		*res = UseRewardRes{
			Success: false,
			Message: "You don't have this reward",
		}

		return nil
	}

	afterQuantity := userRewardDbs[0].Quantity - 1
	if afterQuantity > 0 {
		if err := s.accessor.UpdateUserReward(accessor.UserRewardDb{
			ID:        req.StudentID,
			CollegeID: userDb.CollegeID,
			RewardID:  req.RewardID,
			Quantity:  userRewardDbs[0].Quantity - 1,
		}); err != nil {
			return err
		}
	} else {
		if err := s.accessor.DeleteUserReward(req.StudentID, req.RewardID, userDb.CollegeID); err != nil {
			return err
		}
	}

	
	*res = UseRewardRes{
		Success: true,
		Message: "",
	}

	return nil
}
