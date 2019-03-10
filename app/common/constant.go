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
	MSG_REMOVE_SUCCESS             string = "User Removed Successfully !!"
	MSG_REMOVE_ERROR               string = "Could not remove user !!"
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

	MSG_DUPLICATE_RECORD string = "Duplicate Record"
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

	MOBILE_LENGTH    = 10
	USERNAME_STR_LEN = 3

	DEF_REQUESTS_LIMIT = 60
	DEF_REQUESTS_PAGE  = 0
	QUES_BATCH_SIZE    = 10000

	GROUP_CODE_PREFIX = "GP_"
)
