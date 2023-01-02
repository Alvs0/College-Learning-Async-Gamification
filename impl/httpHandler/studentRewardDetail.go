package httpHandler

import (
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/reward"
	"college-learning-asynchronous-gamification/impl/user"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	stdHttp "net/http"
	"path"
)

func StudentRewardDetailHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "student_reward_detail.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		vars := mux.Vars(req)

		id, ok := vars["id"]
		if !ok {
			return nil, http.WrapError(http.NOT_FOUND_CODE, err)
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

		getRewardReq := reward.GetRewardByIDsReq{
			RewardIDs: []string{id},
		}

		var getRewardRes reward.GetRewardByIDsRes
		if err := services.RewardService.GetRewardByIDs(getRewardReq, &getRewardRes); err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		reward := getRewardRes.Rewards[0]
		imageUrl := reward.ImageUrl
		if reward.ImageUrl == "" {
			imageUrl = "https://storage.googleapis.com/college-async-gamification/not_found.png"
		}

		var data = map[string]interface{}{
			"rewardId":             reward.ID,
			"rewardName":           reward.Name,
			"rewardDescription":    reward.Description,
			"rewardQuantity":       reward.Quantity,
			"rewardMinimalLevel":   reward.MinimalLevel,
			"rewardRequiredPoints": reward.RequiredPoints,
			"rewardImageUrl":       imageUrl,
			"rewardIsActive":       reward.IsActive,
			"collegeProfileImage": getCollegeRes.College.ImageUrl,
			"sessionUrl":          fmt.Sprintf("%v%v", generalData.HostUrl, view.STUDENT_SESSION_PATH),
			"rewardUrl":           fmt.Sprintf("%v%v", generalData.HostUrl, view.STUDENT_REWARD_PATH),
			"collegeDetailUrl":    fmt.Sprintf("%v%v", generalData.HostUrl, view.STUDENT_COLLEGE_DETAIL_PATH),
			"homeUrl":             fmt.Sprintf("%v%v", generalData.HostUrl, view.STUDENT_HOME_PATH),
			"loginUrl":            fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}
