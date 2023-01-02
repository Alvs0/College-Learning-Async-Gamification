package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/session"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	stdHttp "net/http"
	"path"
)

func StudentSessionDetailHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "student_session_detail.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		vars := mux.Vars(req)

		id, ok := vars["id"]
		if !ok {
			return nil, http.WrapError(http.NOT_FOUND_CODE, err)
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

		getSessionReq := session.GetSessionByIDsReq{
			SessionIDs: []string{id},
		}

		var getSessionRes session.GetSessionByIDsRes
		if err := services.SessionService.GetSessionByIDs(getSessionReq, &getSessionRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		session := getSessionRes.Sessions[0]
		imageUrl := session.ImageUrl
		if session.ImageUrl == "" {
			imageUrl = "https://storage.googleapis.com/college-async-gamification/not_found.png"
		}

		var data = map[string]interface{}{
			"sessionId":           session.ID,
			"sessionName":         session.Name,
			"sessionLink":         session.Link,
			"sessionImageUrl":     imageUrl,
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
