package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	stdHttp "net/http"
	"path"
)

const SUPERADMIN_EMAIL = "asd@asd"

func LoginHandler(generalData GeneralData) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "loginPage.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"postLoginUrl": fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
			"registerUrl":  fmt.Sprintf("%v%v", generalData.HostUrl, view.REGISTER_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}

type UserCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	LoggedUser string `json:"loggedUser"`
	RedirectTo string `json:"redirectTo"`
	Message    string `json:"message"`
}

func LoginPostHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		decoder := json.NewDecoder(req.Body)

		var userCredential UserCredential
		if err := decoder.Decode(&userCredential); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		session, err := generalData.SessionStore.Get(req, "user")
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		if userCredential.Email == SUPERADMIN_EMAIL {
			session.Values["user"] = SUPERADMIN_EMAIL
			session.Save(req, res)

			outputByte, err := json.Marshal(LoginOutput{
				LoggedUser: SUPERADMIN_EMAIL,
				RedirectTo: fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_HOME_PATH),
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusOK)
			res.Write(outputByte)

			return nil, nil
		}

		getUserByEmailReq := user.GetUserByEmailReq{
			Email: userCredential.Email,
		}

		var getUserByEmailRes user.GetUserByEmailRes
		if err := services.UserService.GetUserByEmail(getUserByEmailReq, &getUserByEmailRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		if getUserByEmailRes.Results == nil {
			outputByte, err := json.Marshal(LoginOutput{
				LoggedUser: "",
				Message:    "No user found. Please create a new one",
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusNotFound)
			res.Write(outputByte)

			return nil, nil
		}

		err = bcrypt.CompareHashAndPassword([]byte(getUserByEmailRes.Results[0].EncryptedPassword), []byte(userCredential.Password))
		if err != nil {
			outputByte, err := json.Marshal(LoginOutput{
				Message: "Incorrect username or password",
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusNotFound)
			res.Write(outputByte)

			return nil, nil
		}

		session.Values["user"] = userCredential.Email
		session.Save(req, res)
		if getUserByEmailRes.Results[0].IsAdmin {
			outputByte, err := json.Marshal(LoginOutput{
				LoggedUser: userCredential.Email,
				RedirectTo: fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_HOME_PATH),
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusOK)
			res.Write(outputByte)

			return nil, nil
		}

		outputByte, err := json.Marshal(LoginOutput{
			LoggedUser: userCredential.Email,
			RedirectTo: fmt.Sprintf("%v%v", generalData.HostUrl, view.STUDENT_HOME_PATH),
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
