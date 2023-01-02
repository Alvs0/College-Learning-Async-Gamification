package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/reward"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"encoding/json"
	"fmt"
	stdHttp "net/http"
)

type UseRewardInput struct {
	RewardID string `json:"rewardId"`
	ID       string `json:"id"`
}

func StudentUseRewardHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		//vars := mux.Vars(req)
		//
		//id, ok := vars["id"]
		//if !ok {
		//	return nil, http.WrapError(http.NOT_FOUND_CODE, nil)
		//}

		decoder := json.NewDecoder(req.Body)

		var userRewardInput UseRewardInput
		if err := decoder.Decode(&userRewardInput); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		userEmail, isValid := validateUser(generalData, req)
		if !isValid {
			stdHttp.Redirect(res, req, fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH), stdHttp.StatusSeeOther)
		}

		getUserReq := user.GetUserByEmailReq{
			Email: userEmail,
		}

		var getUserRes user.GetUserByEmailRes
		if err := services.UserService.GetUserByEmail(getUserReq, &getUserRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		getCollegeReq := college.GetCollegeByIDReq{
			CollegeID: getUserRes.Results[0].CollegeID,
		}

		var getCollegeRes college.GetCollegeByIDRes
		if err := services.CollegeService.GetCollegeByID(getCollegeReq, &getCollegeRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		useRewardReq := reward.UseRewardReq{
			RewardID:  userRewardInput.RewardID,
			StudentID: getUserRes.Results[0].ID,
		}

		var useRewardRes reward.UseRewardRes
		if err := services.RewardService.UseReward(useRewardReq, &useRewardRes); err != nil {
			outputByte, err := json.Marshal(Output{
				Message: "Failed to use reward",
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusOK)
			res.Write(outputByte)

			return nil, nil
		}

		if !useRewardRes.Success {
			outputByte, err := json.Marshal(Output{
				Message: useRewardRes.Message,
			})
			if err != nil {
				return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(stdHttp.StatusOK)
			res.Write(outputByte)

			return nil, nil
		}

		outputByte, err := json.Marshal(Output{
			RedirectTo: fmt.Sprintf("%v%v", generalData.HostUrl, view.STUDENT_MY_REWARD_PATH),
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
