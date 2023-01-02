package college

func (c *service) GetCollegeByIDs(req GetCollegeByIDsReq, res *GetCollegeByIDsRes) error {
	if req.CollegeIDs == nil {
		return nil
	}

	collegeDbs, err := c.accessor.GetCollegeByIDs(req.CollegeIDs)
	if err != nil {
		return err
	}

	var colleges []College
	for _, collegeDb := range collegeDbs {
		colleges = append(colleges, College{
			ID:       collegeDb.ID,
			Name:     collegeDb.Name,
			Address:  collegeDb.Address,
			ImageUrl: collegeDb.IconURL,
		})
	}

	*res = GetCollegeByIDsRes{
		Colleges: colleges,
	}

	return nil
}
