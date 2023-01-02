package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"html/template"
	stdHttp "net/http"
	"path"
)

func AdminCollegeDetailHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "admin_college_detail.html")
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

		imageUrl := getCollegeRes.College.ImageUrl
		if getCollegeRes.College.ImageUrl == "" {
			imageUrl = "https://storage.googleapis.com/college-async-gamification/not_found.png"
		}

		var data = map[string]interface{}{
			"collegeId":           getCollegeRes.College.ID,
			"collegeName":         getCollegeRes.College.Name,
			"collegeAddress":      getCollegeRes.College.Address,
			"collegeImageUrl":     imageUrl,
			"collegeProfileImage": imageUrl,
			"sessionUrl":          fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_SESSION_PATH),
			"rewardUrl":           fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_REWARD_PATH),
			"studentUrl":          fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_STUDENT_PATH),
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
