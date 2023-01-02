package session

func (c *service) GetSessionByIDs(req GetSessionByIDsReq, res *GetSessionByIDsRes) error {
	sessionDbs, err := c.accessor.GetSessions(req.SessionIDs)
	if err != nil {
		return err
	}

	var sessionObjs []SessionObj
	for _, sessionDb := range sessionDbs {
		sessionObjs = append(sessionObjs, SessionObj{
			ID:        sessionDb.ID,
			CollegeID: sessionDb.CollegeID,
			Name:      sessionDb.Name,
			Link:      sessionDb.Link,
			ImageUrl:  sessionDb.ImageUrl,
		})
	}

	*res = GetSessionByIDsRes{
		Sessions: sessionObjs,
	}

	return nil
}
