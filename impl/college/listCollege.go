package college

func (c *service) ListCollege(req ListCollegeReq, res *ListCollegeRes) error {
	collegeIDs, err := c.accessor.ListCollege()
	if err != nil {
		return err
	}

	*res = ListCollegeRes{
		CollegeIDs: collegeIDs,
	}

	return nil
}

