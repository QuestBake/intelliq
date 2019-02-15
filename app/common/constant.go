package common

const ( // app msgs
	MSG_SAVE_SUCCESS            string = "Data Saved Successfully !!"
	MSG_SAVE_ERROR              string = "Could not save data !!"
	MSG_UPDATE_SUCCESS          string = "Data Updated Successfully !!"
	MSG_UPDATE_ERROR            string = "Could not save data !!"
	MSG_REQUEST_FAILED          string = "Could not process request. Try later !!"
	MSG_INVALID_ID              string = "Invalid Identifier"
	MSG_BAD_INPUT               string = "Bad request params"
	MSG_NO_ROLE                 string = "No such role Exists !!"
	MSG_NO_RECORD               string = "No such record exists !!"
	MSG_INSUFFICIENT_USER_COUNT string = "Expected 2 users, but found "
	MSG_REMOVE_SUCCESS          string = "User Removed Successfully !!"
	MSG_REMOVE_ERROR            string = "Could not remove user !!"

	MSG_DUPLICATE_RECORD string = "Duplicate Record"
)

const ( // codes
	ERR_CODE_DUPLICATE int = 11000
)

const ( // congfig params
	PARAM_KEY_ID   = "_id"
	PARAM_KEY_CODE = "code"

	DATE_TIME_FORMAT = "02 Jan,2006 03:04:05 PM"
	DATE_FORMAT      = "02 Jan,2006" // 01= Month , 02 = Date

	TEMP_PWD_PREFIX = "TP_"

	MIN_VALID_ROLE = 0
	MAX_VALID_ROLE = 4
)
