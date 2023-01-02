package reward

import (
	"college-learning-asynchronous-gamification/accessor"
	"github.com/google/uuid"
)

func (c *service) UpsertReward(req UpsertRewardReq, res *UpsertRewardRes) error {
	id := uuid.New()
	upsertInput := accessor.RewardDb{
		ID:             id.String(),
		CollegeID:      req.Reward.CollegeID,
		Name:           req.Reward.Name,
		Description:    req.Reward.Description,
		ImageUrl:       req.Reward.ImageUrl,
		Quantity:       req.Reward.Quantity,
		RequiredPoints: req.Reward.RequiredPoints,
		MinimalLevel:   req.Reward.MinimalLevel,
		IsActive:       req.Reward.IsActive,
	}

	if err := c.accessor.UpsertReward(upsertInput); err != nil {
		return err
	}

	return nil
}
