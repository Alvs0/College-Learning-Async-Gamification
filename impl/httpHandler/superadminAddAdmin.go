package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	"html/template"
	stdHttp "net/http"
	"path"
)

func SuperAdminAddAdminHandler(generalData GeneralData) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "superadmin_add_admin.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"collegeUrl": fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_COLLEGE_PATH),
			"homeUrl":    fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_HOME_PATH),
			"loginUrl":   fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}

type AdminData struct {
	CollegeName string `json:"collegeName"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Birthdate   string `json:"birthDate"`
	ImageUrl    string `json:"imageUrl"`
}

func SuperAdminAddAdminPostHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		decoder := json.NewDecoder(req.Body)

		var adminData AdminData
		if err := decoder.Decode(&adminData); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		getCollegeReq := college.GetCollegeByNameReq{
			Name: adminData.CollegeName,
		}

		var getCollegeRes college.GetCollegeByNameRes
		if err := services.CollegeService.GetCollegeByName(getCollegeReq, &getCollegeRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		if getCollegeRes.College.ID == "" {
			outputByte, err := json.Marshal(Output{
				Message: "Failed to add college. Make sure inputted college exist!",
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusOK)
			res.Write(outputByte)

			return nil, nil
		}

		addUserReq := user.UpsertUserReq{User: user.UserObj{
			CollegeID:         getCollegeRes.College.ID,
			Name:              adminData.Name,
			Email:             adminData.Email,
			PhoneNumber:       adminData.PhoneNumber,
			BirthDate:         adminData.Birthdate,
			ProfileImageUrl:   adminData.ImageUrl,
			EncryptedPassword: adminData.Password,
			IsAdmin:           true,
		}}

		var addUserRes user.UpsertUserRes
		if err := services.UserService.UpsertUser(addUserReq, &addUserRes); err != nil {
			outputByte, err := json.Marshal(AddCollegeOutput{
				Message: "Failed to add admin",
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
			RedirectTo: fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_ADMIN_PATH),
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
