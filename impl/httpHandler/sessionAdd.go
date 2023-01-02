package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	stdHttp "net/http"
)

type SessionAddPointInput struct {
	Point int `json:"points"`
}

func SessionAddPointHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		decoder := json.NewDecoder(req.Body)

		var sessionAddPointInput SessionAddPointInput
		if err := decoder.Decode(&sessionAddPointInput); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		userEmail, isValid := validateUser(generalData, req)
		if !isValid {
			stdHttp.Redirect(res, req, fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH), stdHttp.StatusSeeOther)
		}

		getUserReq := user.GetUserByEmailReq{
			Email: userEmail,
		}

		var getUserRes user.GetUserByEmailRes
		if err := services.UserService.GetUserByEmail(getUserReq, &getUserRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		getCollegeReq := college.GetCollegeByIDReq{
			CollegeID: getUserRes.Results[0].CollegeID,
		}

		var getCollegeRes college.GetCollegeByIDRes
		if err := services.CollegeService.GetCollegeByID(getCollegeReq, &getCollegeRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		updateUserPointReq := user.UpdateUserPointReq{
			UserID:    getUserRes.Results[0].ID,
			CollegeID: getUserRes.Results[0].CollegeID,
			Point:     sessionAddPointInput.Point,
		}

		var updateUserPointRes user.UpdateUserPointRes
		if err := services.UserService.UpdateUserPoint(updateUserPointReq, &updateUserPointRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}
