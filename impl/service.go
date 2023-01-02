package impl

import (
	"college-learning-asynchronous-gamification/accessor"
	"college-learning-asynchronous-gamification/gcs"
	"college-learning-asynchronous-gamification/impl/college"
	"college-learning-asynchronous-gamification/impl/common"
	"college-learning-asynchronous-gamification/impl/reward"
	"college-learning-asynchronous-gamification/impl/session"
	"college-learning-asynchronous-gamification/impl/user"
)

type Services struct {
	UserService    user.Service
	CollegeService college.Service
	RewardService  reward.Service
	SessionService session.Service
	CommonService  common.Service
}

func NewServices(accessor accessor.Accessor, fileDownloader gcs.DownloaderClient, fileUploader gcs.UploaderClient) Services {
	return Services{
		UserService:    user.NewService(accessor),
		CollegeService: college.NewService(accessor),
		RewardService:  reward.NewService(accessor),
		SessionService: session.NewService(accessor),
		CommonService:  common.NewService(fileDownloader, fileUploader),
	}
}
