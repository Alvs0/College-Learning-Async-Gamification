package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	stdHttp "net/http"
	"os"
	"path"
	"path/filepath"
)

type GeneralData struct {
	HostUrl      string
	SessionStore sessions.Store
}

type Output struct {
	RedirectTo string `json:"redirectTo"`
	Message    string `json:"message"`
}

func IndexHandler(generalData GeneralData) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"loginUrl": fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}

func getViewPath() string {
	currPath, _ := os.Getwd()
	pageHandlerPath := filepath.Join(currPath, "pageHandler")
	viewPath := filepath.Join(pageHandlerPath, "view")
	return viewPath
}
