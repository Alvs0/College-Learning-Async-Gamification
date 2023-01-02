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

func SuperAdminAddCollegeHandler(generalData GeneralData) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "superadmin_add_college.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"adminUrl": fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_ADMIN_PATH),
			"homeUrl":  fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_HOME_PATH),
			"loginUrl": fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}

type CollegeData struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	ImageUrl string `json:"imageUrl"`
}

type AddCollegeOutput struct {
	RedirectTo string `json:"redirectTo"`
	Message    string `json:"message"`
}

func SuperAdminAddCollegePostHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		decoder := json.NewDecoder(req.Body)

		var collegeData CollegeData
		if err := decoder.Decode(&collegeData); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		addCollegeReq := college.UpsertCollegeReq{
			Name:     collegeData.Name,
			Address:  collegeData.Address,
			ImageUrl: collegeData.ImageUrl,
		}

		var addCollegeRes college.UpsertCollegeRes
		if err := services.CollegeService.UpsertCollege(addCollegeReq, &addCollegeRes); err != nil {
			outputByte, err := json.Marshal(AddCollegeOutput{
				Message: "Failed to add college",
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusOK)
			res.Write(outputByte)

			return nil, nil
		}

		outputByte, err := json.Marshal(AddCollegeOutput{
			RedirectTo: fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_COLLEGE_PATH),
		})
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(stdHttp.StatusOK)
		res.Write(outputByte)

		return nil, nil
	}
}
