package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/session"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	stdHttp "net/http"
	"path"
	"strings"
)

func SessionHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "session.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		vars := mux.Vars(req)

		id, ok := vars["id"]
		if !ok {
			return nil, http.WrapError(http.NOT_FOUND_CODE, err)
		}

		getSessionReq := session.GetSessionByIDsReq{
			SessionIDs: []string{id},
		}

		var getSessionRes session.GetSessionByIDsRes
		if err := services.SessionService.GetSessionByIDs(getSessionReq, &getSessionRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		var data = map[string]interface{}{
			"videoLink":  fmt.Sprintf("%v?controls=0", getSessionRes.Sessions[0].Link),
			"videoId":    strings.ReplaceAll(getSessionRes.Sessions[0].Link, "https://www.youtube.com/embed/", ""),
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
