package view

const (
	// Common Path
	INDEX_PATH       = "/"
	LOGIN_PATH       = "/login"
	REGISTER_PATH    = "/register"
	STATIC_PATH      = "/static/"
	UPLOAD_FILE_PATH = "/upload_file"

	// Superadmin Path
	SUPERADMIN_HOME_PATH           = "/home_sa"
	SUPERADMIN_ADD_COLLEGE_PATH    = "/college_sa/add"
	SUPERADMIN_COLLEGE_PATH        = "/college_sa"
	SUPERADMIN_COLLEGE_DETAIL_PATH = "/college_detail_sa/{id}"
	SUPERADMIN_DELETE_COLLEGE_PATH = "/college_delete_sa"
	SUPERADMIN_ADMIN_PATH          = "/admin_sa"
	SUPERADMIN_ADD_ADMIN_PATH      = "/admin_sa/add"
	SUPERADMIN_ADMIN_DETAIL_PATH   = "/admin_detail_sa/{id}"
	SUPERADMIN_DELETE_ADMIN_PATH   = "/admin_delete_sa"

	// Admin Path
	ADMIN_HOME_PATH           = "/home_a"
	ADMIN_SESSION_PATH        = "/session_a"
	ADMIN_SESSION_DETAIL_PATH = "/session_detail_a/{id}"
	ADMIN_ADD_SESSION_PATH    = "/session_a/add"
	ADMIN_DELETE_SESSION_PATH = "/session_delete_a"
	ADMIN_REWARD_PATH         = "/reward_a"
	ADMIN_ADD_REWARD_PATH     = "/reward_a/add"
	ADMIN_REWARD_DETAIL_PATH  = "/reward_detail_a/{id}"
	ADMIN_STUDENT_PATH        = "/student_a"
	ADMIN_STUDENT_DETAIL_PATH = "/student_detail_a/{id}"
	ADMIN_COLLEGE_DETAIL_PATH = "/college_detail_a"
	ADMIN_DELETE_REWARD_PATH  = "/reward_delete_a"

	// Student Path
	STUDENT_HOME_PATH           = "/home_st"
	STUDENT_COLLEGE_DETAIL_PATH = "/college_detail_st"
	STUDENT_SESSION_PATH        = "/session_st"
	STUDENT_SESSION_DETAIL_PATH = "/session_detail_st/{id}"
	STUDENT_REWARD_PATH         = "/reward_st"
	STUDENT_REWARD_DETAIL_PATH  = "/reward_detail_st/{id}"
	STUDENT_REWARD_CLAIM_PATH   = "/reward_claim_st"
	STUDENT_MY_REWARD_PATH      = "/my_reward_st"
	STUDENT_USE_REWARD_PATH     = "/reward_use_st"

	// Session Path
	SESSION_PATH      = "/session/{id}"
	SESSION_ADD_POINT = "/session_add_point"
)
