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
	"html/template"
	stdHttp "net/http"
	"path"
	"strconv"
)

func AdminAddRewardHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		var filepath = path.Join(getViewPath(), "admin_add_reward.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
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

		var data = map[string]interface{}{
			"collegeProfileImage": getCollegeRes.College.ImageUrl,
			"sessionUrl":          fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_SESSION_PATH),
			"rewardUrl":           fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_REWARD_PATH),
			"studentUrl":          fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_STUDENT_PATH),
			"collegeDetailUrl":    fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_COLLEGE_DETAIL_PATH),
			"homeUrl":             fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_HOME_PATH),
			"loginUrl":            fmt.Sprintf("%v%v", generalData.HostUrl, view.LOGIN_PATH),
		}

		err = tmpl.Execute(res, data)
		if err != nil {
			return nil, http.WrapError(http.INTERNAL_SERVER_ERROR_CODE, err)
		}

		return nil, nil
	}
}

type RewardData struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Quantity       string `json:"quantity"`
	MinimalLevel   string `json:"minimalLevel"`
	RequiredPoints string `json:"requiredPoints"`
	ImageUrl       string `json:"imageUrl"`
	IsActive       string `json:"isActive"`
}

func AdminAddRewardPostHandler(generalData GeneralData, services impl.Services) http.HandlerFunc {
	return func(req *stdHttp.Request, res stdHttp.ResponseWriter) (interface{}, *http.ErrorWrapper) {
		decoder := json.NewDecoder(req.Body)

		var rewardData RewardData
		if err := decoder.Decode(&rewardData); err != nil {
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

		quantityInt, _ := strconv.Atoi(rewardData.Quantity)
		minimalLevelInt, _ := strconv.Atoi(rewardData.MinimalLevel)
		requiredPointInt, _ := strconv.Atoi(rewardData.RequiredPoints)

		var isActive bool
		if rewardData.IsActive == "on" {
			isActive = true
		}

		addRewardReq := reward.UpsertRewardReq{
			Reward: reward.RewardObj{
				CollegeID:      getCollegeReq.CollegeID,
				Name:           rewardData.Name,
				ImageUrl:       rewardData.ImageUrl,
				Description:    rewardData.Description,
				Quantity:       quantityInt,
				RequiredPoints: requiredPointInt,
				MinimalLevel:   minimalLevelInt,
				IsActive:       isActive,
			},
		}

		var addRewardRes reward.UpsertRewardRes
		if err := services.RewardService.UpsertReward(addRewardReq, &addRewardRes); err != nil {
			outputByte, err := json.Marshal(Output{
				Message: "Failed to add reward",
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
			RedirectTo: fmt.Sprintf("%v%v", generalData.HostUrl, view.ADMIN_REWARD_PATH),
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
