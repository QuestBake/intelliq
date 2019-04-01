package common

const ( // app msgs
	MSG_SAVE_SUCCESS               string = "Data Saved Successfully !!"
	MSG_SAVE_ERROR                 string = "Could not save data !!"
	MSG_UPDATE_SUCCESS             string = "Data Updated Successfully !!"
	MSG_UPDATE_ERROR               string = "Could not save data !!"
	MSG_DELETE_SUCCESS             string = "Data removed Successfully !!"
	MSG_DELETE_ERROR               string = "Could not remove data !!"
	MSG_REQUEST_FAILED             string = "Could not process request. Try later !!"
	MSG_INVALID_ID                 string = "Invalid Identifier"
	MSG_INVALID_GROUP              string = "Invalid Group Code"
	MSG_INVALID_MOBILE             string = "Invalid Mobile Number"
	MSG_BAD_INPUT                  string = "Bad request data"
	MSG_INVALID_STATE              string = "Invalid State !!"
	MSG_NO_ROLE                    string = "No such role Exists !!"
	MSG_NO_STATUS                  string = "No such status Exists !!"
	MSG_NO_RECORD                  string = "No record found !!"
	MSG_INSUFFICIENT_USER_COUNT    string = "Expected 2 users, but found "
	MSG_REMOVE_USER_SUCCESS        string = "User Removed Successfully !!"
	MSG_REMOVE_USER_ERROR          string = "Could not remove user !!"
	MSG_REMOVE_TEMPLATE_SUCCESS    string = "Template Removed Successfully !!"
	MSG_REMOVE_TEMPLATE_ERROR      string = "Could not remove template !!"
	MSG_REMOVE_TEST_SUCCESS        string = "Test Paper Removed Successfully !!"
	MSG_REMOVE_TEST_ERROR          string = "Could not remove test paper !!"
	MSG_LOGIN_SUCCESS              string = "Login Sucess !!"
	MSG_INVALID_CREDENTIALS_MOBILE string = "Mobile number not registered with us !!"
	MSG_INVALID_CREDENTIALS_PWD    string = "Invalid Password!!"
	MSG_QUES_SUBMIT_SUCCESS        string = "Successfully submitted for review !!"
	MSG_UNATHORIZED_ACCESS         string = "Unauthorized access to DB !!"
	MSG_QUES_STATUS_SUCCESS        string = "Successfully updated the status !!"
	MSG_QUES_STATUS_ERROR          string = "Status could not be updated !!"
	MSG_QUES_REMOVE_SUCCESS        string = "Successfully removed question !!"
	MSG_QUES_REMOVE_ERROR          string = "Question could not be removed !!"
	MSG_CORRUPT_DATA               string = "Corrupt Criteria Data"
	MSG_DUPLICATE_RECORD           string = "Duplicate Record"
	MSG_NAME_MIN_LENGTH_ERROR      string = "Name should be atleast 3 characters"
	MSG_MOBILE_MIN_LENGTH_ERROR    string = "Mobile should be 10 digits"
	MSG_PWD_MIN_LENGTH_ERROR       string = "Password should be atleast 8 characters"
	MSG_PWD_RESET_SUCCESS          string = "Password changed successfully !!"
	MSG_MOBILE_UPDATE_SUCCESS      string = "Mobile number changed successfully !!"
	MSG_USER_AUTH_ERROR            string = "User Authentication Error !!"
	MSG_USER_SESSION_ERROR         string = "Session Not Found !!"
	MSG_USER_FORGERY_ERROR         string = "Detected user forgery !! !!"
)

const ( // codes
	ERR_CODE_DUPLICATE int = 11000
)

const ( // congfig params
	PARAM_KEY_ID     = "_id"
	PARAM_KEY_CODE   = "code"
	PARAM_KEY_MOBILE = "mobile"

	DATE_TIME_FORMAT = "02 Jan,2006 03:04:05 PM"
	DATE_FORMAT      = "02 Jan,2006" // 01= Month , 02 = Date

	TEMP_PWD_PREFIX = "TP_"
	USERNAME_PREFIX = "user@"

	MIN_VALID_ROLE = 0
	MAX_VALID_ROLE = 4

	MIN_VALID_STD = 1
	MAX_VALID_STD = 12

	MOBILE_LENGTH       = 10
	USERNAME_MIN_LENGTH = 3
	PWD_MIN_LENGTH      = 8

	DEF_REQUESTS_LIMIT = 60
	DEF_REQUESTS_PAGE  = 0
	QUES_BATCH_SIZE    = 10000

	OTP_LENGTH      = 6
	OTP_UPPER_BOUND = 989899
	OTP_LOWER_BOUND = 101010

	GROUP_CODE_PREFIX           = "GP_"
	CACHE_STORE_KEY             = "CACHE_STORE"
	REQUEST_OTP_SESSION_ID_KEY  = "RQST_OTP_SESS_ID"
	RESPONSE_OTP_SESSION_ID_KEY = "RESP_OTP_SESS_ID"

	COOKIE_SESSION_KEY  = "c_user"
	COOKIE_XSRF_KEY     = "XSRF-TOKEN"
	HEADER_XSRF_KEY     = "X-Xsrf-Token"
	CORS_REQUEST_METHOD = "OPTIONS"

	USER_SESSION_TIMEOUT    = 120 //minutes
	COOKIE_SESSION_TIMEOUT  = 180
	CACHE_OTP_TIMEOUT       = 5
	CACHE_OBJ_LONG_TIMEOUT  = 60
	CACHE_OBJ_SHORT_TIMEOUT = 30

	APP_NAME     = "IntelliQ"
	APP_PORT     = ":8080"
	CACHE_PORT   = ":6379"
	CACHE_DOMAIN = "localhost"

	PRIVATE_KEY_FILEPATH = "/Users/lionheart/.ssh/appKey.priv"
	SSL_CERT_FILEPATH    = "/Users/lionheart/.ssh/ssl.crt"
	SSL_KEY_FILEPATH     = "/Users/lionheart/.ssh/sslKey.key"
)
