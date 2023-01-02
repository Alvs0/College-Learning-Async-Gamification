package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	stdHttp "net/http"
)

func SuperAdminAdminDeleteHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		decoder := json.NewDecoder(req.Body)

		var deleteInput DeleteInput
		if err := decoder.Decode(&deleteInput); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		deleteAdminByIDReq := user.DeleteUserReq{
			ID: deleteInput.ID,
		}

		var deleteAdminByIDRes user.DeleteUserRes
		if err := services.UserService.DeleteUser(deleteAdminByIDReq, &deleteAdminByIDRes); err != nil {
			outputByte, err := json.Marshal(DeleteOutput{
				Message: "Failed to Delete Admin",
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusOK)
			res.Write(outputByte)

			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		outputByte, err := json.Marshal(DeleteOutput{
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
