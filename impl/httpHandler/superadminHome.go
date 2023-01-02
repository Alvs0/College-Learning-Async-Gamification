package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"html/template"
	stdHttp "net/http"
	"path"
)

func SuperAdminHomeHandler(generalData GeneralData) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "superadmin_home.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"collegeUrl": fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_COLLEGE_PATH),
			"adminUrl":   fmt.Sprintf("%v%v", generalData.HostUrl, view.SUPERADMIN_ADMIN_PATH),
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
