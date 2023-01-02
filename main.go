package main

import (
	"college-learning-asynchronous-gamification/accessor"
	"college-learning-asynchronous-gamification/config"
	"college-learning-asynchronous-gamification/gcs"
	"college-learning-asynchronous-gamification/http"
	"college-learning-asynchronous-gamification/impl"
	"college-learning-asynchronous-gamification/impl/httpHandler"
	"college-learning-asynchronous-gamification/pageHandler/view"
	"github.com/gorilla/sessions"
	"log"
	stdHttp "net/http"
)

func main() {
	runner := http.NewDaemonRunner()
	runner.Run(&component{}, &config.ConfigMap{})
}

type component struct {
	server http.HttpServer
}

func (c *component) Instantiate(configMap interface{}) (err error) {
	cm, ok := configMap.(*config.ConfigMap)
	if !ok {
		log.Fatal("error casting config map")
	}

	defaultServer := http.NewHttpServer()

	sqlAdapter := accessor.NewSqlAdapter(cm.MySQLConfig)
	accessor := accessor.NewAccessor(sqlAdapter)

	fileDownloader, err := gcs.NewDownloader(cm.GCSConfig)
	if err != nil {
		log.Fatal("error creating downloader")
	}

	key := []byte("secret-key")
	sessionStore := sessions.NewCookieStore(key)

	fileUploader, err := gcs.NewUploader(cm.GCSConfig)
	if err != nil {
		log.Fatal("error creating uploader")
	}

	services := impl.NewServices(accessor, fileDownloader, fileUploader)

	c.server = defaultServer

	generalData := httpHandler.GeneralData{
		HostUrl:      cm.HostUrl,
		SessionStore: sessionStore,
	}

	fileServer := stdHttp.FileServer(stdHttp.Dir("./static/"))
	c.server.AddStaticHandler(view.STATIC_PATH, stdHttp.StripPrefix("/static/", fileServer))

	// COMMON HTTP HANDLER
	c.server.AddHttpHandler(
		view.INDEX_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.IndexHandler(generalData),
		),
	)

	c.server.AddHttpHandler(
		view.LOGIN_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.LoginHandler(generalData),
		),
	)

	c.server.AddHttpHandler(
		view.LOGIN_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.LoginPostHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.REGISTER_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.RegisterHandler(generalData),
		),
	)

	c.server.AddHttpHandler(
		view.REGISTER_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.RegisterPostHandler(generalData, services),
		),
	)

	// SUPERADMIN HTTP HANDLER
	c.server.AddHttpHandler(
		view.SUPERADMIN_HOME_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SuperAdminHomeHandler(generalData),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_COLLEGE_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SuperAdminCollegeHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_ADD_COLLEGE_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SuperAdminAddCollegeHandler(generalData),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_ADD_COLLEGE_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.SuperAdminAddCollegePostHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_COLLEGE_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SuperAdminCollegeDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_DELETE_COLLEGE_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.SuperAdminCollegeDeleteHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_ADMIN_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SuperAdminAdminHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_ADD_ADMIN_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SuperAdminAddAdminHandler(generalData),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_ADD_ADMIN_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.SuperAdminAddAdminPostHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_ADMIN_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SuperAdminAdminDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SUPERADMIN_DELETE_ADMIN_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.SuperAdminAdminDeleteHandler(generalData, services),
		),
	)

	// ADMIN HTTP HANDLER
	c.server.AddHttpHandler(
		view.ADMIN_HOME_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminHomeHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_COLLEGE_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminCollegeDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_REWARD_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminRewardHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_ADD_REWARD_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminAddRewardHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_ADD_REWARD_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.AdminAddRewardPostHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_REWARD_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminRewardDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_DELETE_REWARD_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.AdminRewardDeleteHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_STUDENT_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminStudentHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_STUDENT_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminStudentDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_SESSION_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminSessionHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_ADD_SESSION_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminAddSessionHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_ADD_SESSION_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.AdminAddSessionPostHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_SESSION_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.AdminSessionDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.ADMIN_DELETE_SESSION_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.AdminSessionDeleteHandler(generalData, services),
		),
	)

	// STUDENT HTTP HANDLER
	c.server.AddHttpHandler(
		view.STUDENT_HOME_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.StudentHomeHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_COLLEGE_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.StudentCollegeDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_REWARD_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.StudentRewardHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_REWARD_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.StudentRewardDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_REWARD_CLAIM_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.StudentClaimRewardHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_MY_REWARD_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.StudentMyRewardHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_USE_REWARD_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.StudentUseRewardHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_SESSION_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.StudentSessionHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.STUDENT_SESSION_DETAIL_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.StudentSessionDetailHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SESSION_PATH,
		stdHttp.MethodGet,
		http.NewHandler(
			httpHandler.SessionHandler(generalData, services),
		),
	)

	c.server.AddHttpHandler(
		view.SESSION_ADD_POINT,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.SessionAddPointHandler(generalData, services),
		),
	)

	// COMMON HTTP HANDLER
	c.server.AddHttpHandler(
		view.UPLOAD_FILE_PATH,
		stdHttp.MethodPost,
		http.NewHandler(
			httpHandler.UploadImagePostHandler(generalData, fileUploader),
		),
	)

	return nil
}

func (c *component) Start() (err error) {
	c.server.Start()
	return
}

func (c *component) Stop() (err error) {
	c.server.Stop()
	return
}
