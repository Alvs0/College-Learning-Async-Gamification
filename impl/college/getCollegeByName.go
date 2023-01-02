package college

func (c *service) GetCollegeByName(req GetCollegeByNameReq, res *GetCollegeByNameRes) error {
	collegeDb, err := c.accessor.GetCollegeByName(req.Name)
	if err != nil {
		return err
	}

	*res = GetCollegeByNameRes{
		College: College{
			ID:       collegeDb.ID,
			Name:     collegeDb.Name,
			Address:  collegeDb.Address,
			ImageUrl: collegeDb.IconURL,
		},
	}

	return nil
}
