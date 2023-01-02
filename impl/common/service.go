package common

import "college-learning-asynchronous-gamification/gcs"

type Service interface {
	UploadFile(req UploadFileReq, res *UploadFileRes) error
}

type service struct {
	fileDownloader gcs.DownloaderClient
	fileUploader   gcs.UploaderClient
}

func NewService(fileDownloader gcs.DownloaderClient, fileUploader gcs.UploaderClient) Service {
	return &service{
		fileDownloader: fileDownloader,
		fileUploader:   fileUploader,
	}
}

type UploadFileReq struct{}

type UploadFileRes struct {
	URL string `json:"url"`
}
