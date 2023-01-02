package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	stdHttp "net/http"
	"path"
)

func SuperAdminAdminDetailHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "superadmin_admin_detail.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		vars := mux.Vars(req)

		id, ok := vars["id"]
		if !ok {
			return nil, http.WrapError(http.NOT_FOUND_CODE, err)
		}

		getUserByIDReq := user.GetUserByIDsReq{
			UserIDs: []string{id},
		}

		var getUserByIDRes user.GetUserByIDsRes
		if err := services.UserService.GetUserByIDs(getUserByIDReq, &getUserByIDRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		user := getUserByIDRes.Users[0]
		imageUrl := user.ProfileImageUrl
		if user.ProfileImageUrl == "" {
			imageUrl = "https://storage.googleapis.com/college-async-gamification/not_found.png"
		}

		getCollegeReq := college.GetCollegeByIDReq{
			CollegeID: user.CollegeID,
		}

		var getCollegeRes college.GetCollegeByIDRes
		if err := services.CollegeService.GetCollegeByID(getCollegeReq, &getCollegeRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"adminId":          user.ID,
			"adminName":        user.Name,
			"adminEmail":       user.Email,
			"adminPhoneNumber": user.PhoneNumber,
			"adminBirthdate":   user.BirthDate,
			"adminImageUrl":    imageUrl,
			"adminCollegeName": getCollegeRes.College.Name,
			"homeUrl":          fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_HOME_PATH),
			"loginUrl":         fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}
