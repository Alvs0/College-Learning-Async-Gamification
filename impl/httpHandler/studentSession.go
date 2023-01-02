package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/session"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	"html/template"
	stdHttp "net/http"
	"path"
)

func StudentSessionHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "student_session.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
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

		listSessionReq := session.ListSessionReq{
			CollegeID: getCollegeReq.CollegeID,
		}

		var listSessionRes session.ListSessionRes
		if err := services.SessionService.ListSession(listSessionReq, &listSessionRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		getSessionReq := session.GetSessionByIDsReq{
			SessionIDs: listSessionRes.SessionIDs,
		}

		var getSessionRes session.GetSessionByIDsRes
		if err := services.SessionService.GetSessionByIDs(getSessionReq, &getSessionRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		sessionBytes, err := json.Marshal(getSessionRes.Sessions)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"sessionData":         string(sessionBytes),
			"collegeProfileImage": getCollegeRes.College.ImageUrl,
			"sessionUrl":          fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_SESSION_PATH),
			"rewardUrl":           fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_REWARD_PATH),
			"collegeDetailUrl":    fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_COLLEGE_DETAIL_PATH),
			"homeUrl":             fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_HOME_PATH),
			"loginUrl":            fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}
