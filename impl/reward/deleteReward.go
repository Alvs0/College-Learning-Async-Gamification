package reward

func (c *service) DeleteReward(req DeleteRewardReq, res *DeleteRewardRes) error {
	if err := c.accessor.DeleteReward(req.RewardID); err != nil {
		return err
	}

	return nil
}
