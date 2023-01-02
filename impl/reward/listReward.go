package reward

import "college-learning-asynchronous-gamification/accessor"

func (c *service) ListReward(req ListRewardReq, res *ListRewardRes) error {
	rewardIDs, err := c.accessor.ListReward(accessor.RewardListInput{
		CollegeID: req.CollegeID,
	})
	if err != nil {
		return err
	}

	*res = ListRewardRes{
		RewardIDs: rewardIDs,
	}

	return nil
}
