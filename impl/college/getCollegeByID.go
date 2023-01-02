package college

func (c *service) GetCollegeByID(req GetCollegeByIDReq, res *GetCollegeByIDRes) error {
	collegeDb, err := c.accessor.GetCollegeByID(req.CollegeID)
	if err != nil {
		return err
	}

	*res = GetCollegeByIDRes{
		College: College{
			ID:       collegeDb.ID,
			Name:     collegeDb.Name,
			Address:  collegeDb.Address,
			ImageUrl: collegeDb.IconURL,
		},
	}

	return nil
}
