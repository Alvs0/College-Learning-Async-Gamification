package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	"html/template"
	stdHttp "net/http"
	"path"
)

func SuperAdminCollegeHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "superadmin_college.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		listCollegeReq := college.ListCollegeReq{}
		var listCollegeRes college.ListCollegeRes
		if err := services.CollegeService.ListCollege(listCollegeReq, &listCollegeRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		getCollegesReq := college.GetCollegeByIDsReq{
			CollegeIDs: listCollegeRes.CollegeIDs,
		}

		var getCollegesRes college.GetCollegeByIDsRes
		if err := services.CollegeService.GetCollegeByIDs(getCollegesReq, &getCollegesRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		collegesBytes, err := json.Marshal(getCollegesRes.Colleges)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"collegeData":   string(collegesBytes),
			"addCollegeUrl": fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_ADD_COLLEGE_PATH),
			"homeUrl":       fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_HOME_PATH),
			"loginUrl":      fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}
