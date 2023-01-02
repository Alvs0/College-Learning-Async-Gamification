package college

func (c *service) DeleteCollege(req DeleteCollegeReq, res *DeleteCollegeRes) error {
	if err := c.accessor.DeleteCollege(req.ID); err != nil {
		return err
	}

	*res = DeleteCollegeRes{}

	return nil
}