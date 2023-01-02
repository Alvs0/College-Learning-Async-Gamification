package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	stdHttp "net/http"
	"path"
)

func SuperAdminCollegeDetailHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "superadmin_college_detail.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		vars := mux.Vars(req)

		id, ok := vars["id"]
		if !ok {
			return nil, http.WrapError(http.NOT_FOUND_CODE, err)
		}

		getCollegeByIDReq := college.GetCollegeByIDReq{
			CollegeID: id,
		}

		var getCollegeByIDRes college.GetCollegeByIDRes
		if err := services.CollegeService.GetCollegeByID(getCollegeByIDReq, &getCollegeByIDRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		imageUrl := getCollegeByIDRes.College.ImageUrl
		if getCollegeByIDRes.College.ImageUrl == "" {
			imageUrl = "https://storage.googleapis.com/college-async-gamification/not_found.png"
		}

		var data = map[string]interface{}{
			"collegeId":       getCollegeByIDRes.College.ID,
			"collegeName":     getCollegeByIDRes.College.Name,
			"collegeAddress":  getCollegeByIDRes.College.Address,
			"collegeImageUrl": imageUrl,
			"homeUrl":         fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_HOME_PATH),
			"loginUrl":        fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}
