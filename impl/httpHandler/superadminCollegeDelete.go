package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	stdHttp "net/http"
)

type DeleteInput struct {
	ID string `json:"id"`
}

type DeleteOutput struct {
	RedirectTo string `json:"redirectTo"`
	Message    string `json:"message"`
}

func SuperAdminCollegeDeleteHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		decoder := json.NewDecoder(req.Body)

		var deleteInput DeleteInput
		if err := decoder.Decode(&deleteInput); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		deleteCollegeByIDReq := college.DeleteCollegeReq{
			ID: deleteInput.ID,
		}

		var deleteCollegeByIDRes college.DeleteCollegeRes
		if err := services.CollegeService.DeleteCollege(deleteCollegeByIDReq, &deleteCollegeByIDRes); err != nil {
			outputByte, err := json.Marshal(DeleteOutput{
				Message: "Failed to Delete College",
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
