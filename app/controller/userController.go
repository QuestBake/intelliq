package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"intelliq/app/cachestore"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/security"
	"intelliq/app/service"
)

//AddNewUser adds new user
func AddNewUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddNewUser(&user)
	ctx.JSON(http.StatusOK, res)
}

//UpdateUserProfile updates user profile
func UpdateUserProfile(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateUser(&user)
	if res.Status == enums.Status.SUCCESS && res.Body != nil {
		if cachestore.CheckCache(ctx, user.Mobile) {
			cachestore.SetCache(ctx, user.Mobile, user,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
		if cachestore.CheckCache(ctx, user.UserID.Hex()) {
			cachestore.SetCache(ctx, user.UserID.Hex(), user,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
	}
	ctx.JSON(http.StatusOK, res)
}

//ListAllSchoolAdminsUnderGroup fetches all users with role schooladmin
func ListAllSchoolAdminsUnderGroup(ctx *gin.Context) {
	groupID := ctx.Param("groupId")
	res := service.FetchAllSchoolAdmins(groupID)
	ctx.JSON(http.StatusOK, res)
}

//ListAllTeachersUnderSchool fetches all users within school
func ListAllTeachersUnderSchool(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	res := service.FetchAllTeachers(schoolID)
	ctx.JSON(http.StatusOK, res)
}

//ListAllTeachersUnderReviewer fetches all users within school
func ListAllTeachersUnderReviewer(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	reviewerID := ctx.Param("reviewerId")
	res := service.FetchAllTeachersUnderReviewer(schoolID, reviewerID)
	ctx.JSON(http.StatusOK, res)
}

//ListSelectedTeachers fetches all users within school for specific role (teacher/apporver)
func ListSelectedTeachers(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	roleType := ctx.Param("roleType")
	res := service.FetchSelectedTeachers(schoolID, roleType)
	ctx.JSON(http.StatusOK, res)
}

//TransferRole transfers role from one user to another
func TransferRole(ctx *gin.Context) {
	roleType := ctx.Param("roleType")
	fromUserID := ctx.Param("fromUser")
	toUserID := ctx.Param("toUser")
	if len(roleType) == 0 || len(fromUserID) == 0 || len(toUserID) == 0 {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res, mobiles := service.TransferUserRole(roleType, fromUserID, toUserID)
	if res.Status == enums.Status.SUCCESS {
		if mobiles != nil {
			cachestore.RemoveCache(ctx, mobiles[0])
			cachestore.RemoveCache(ctx, mobiles[1])
		}
		cachestore.RemoveCache(ctx, fromUserID)
		cachestore.RemoveCache(ctx, toUserID)
	}
	ctx.JSON(http.StatusOK, res)
}

//RemoveUserFromSchool removes user from current school
func RemoveUserFromSchool(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	userID := ctx.Param("userId")
	res := service.RemoveUserFromSchool(schoolID, userID)
	ctx.JSON(http.StatusOK, res)
}

//AddBulkUsers adds new users
func AddBulkUsers(ctx *gin.Context) {
	var users model.Users
	err := ctx.BindJSON(&users)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddBulkUser(users)
	ctx.JSON(http.StatusOK, res)
}

//UpdateBulkUsers updates bulk users
func UpdateBulkUsers(ctx *gin.Context) {
	var users model.Users
	err := ctx.BindJSON(&users)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateBulkUsers(users)
	ctx.JSON(http.StatusOK, res)
}

//ListUserByMobileOrID get user info by id/mobile/username
func ListUserByMobileOrID(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	if len(key) == 0 || len(val) == 0 ||
		(key != common.PARAM_KEY_ID &&
			key != common.PARAM_KEY_MOBILE && key != common.PARAM_KEY_USERNAME) {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if cachestore.CheckCache(ctx, val) {
		cacheVal := cachestore.GetCache(ctx, val).(string)
		var user model.User
		json.Unmarshal([]byte(cacheVal), &user)
		ctx.JSON(http.StatusOK, utility.GetSuccessResponse(user))
	} else {
		res := service.FetchUserByMobileOrID(key, val)
		if res.Status == enums.Status.SUCCESS && res.Body != nil {
			cachestore.SetCache(ctx, val, res.Body,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
		ctx.JSON(http.StatusOK, res)
	}
}

//ResetUserPassword resets user password either forgotten or renew
func ResetUserPassword(ctx *gin.Context) {
	var pwdDto dto.PasswordDto
	err := ctx.BindJSON(&pwdDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.ResetPassword(&pwdDto)
	ctx.JSON(http.StatusOK, res)
}

//UpdateUserMobile updates user mobile num post OTP verification
func UpdateUserMobile(ctx *gin.Context) {
	var pwdDto dto.PasswordDto
	err := ctx.BindJSON(&pwdDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateMobile(&pwdDto)
	if res.Status == enums.Status.SUCCESS && res.Body != nil {
		if cachestore.CheckCache(ctx, pwdDto.UserID.Hex()) {
			cacheVal := cachestore.GetCache(ctx, pwdDto.UserID.Hex()).(string)
			var user model.User
			err := json.Unmarshal([]byte(cacheVal), &user)
			if err == nil {
				oldMobile := user.Mobile
				if len(oldMobile) == common.MOBILE_LENGTH {
					user.Mobile = pwdDto.Mobile
					cachestore.SetCache(ctx, pwdDto.UserID.Hex(), user,
						common.CACHE_OBJ_LONG_TIMEOUT, true)
					cachestore.RemoveCache(ctx, oldMobile)
				}
			}
		}
	}
	ctx.JSON(http.StatusOK, res)
}

//ForgotPasswordOTP sends OTP to existing mobile number
func ForgotPasswordOTP(ctx *gin.Context) {
	mobile := ctx.Param("mobile")
	res, otp := service.SendOTP(mobile, true)
	if res.Status == enums.Status.SUCCESS {
		if ok, sessID := createOTPSession(ctx, otp); ok {
			sessionToken := security.GenerateToken(
				"Contact", mobile,
				common.CACHE_OTP_TIMEOUT)
			res.Body = sessID
			if len(sessionToken) > 0 {
				security.SetCookie(ctx, sessionToken,
					common.CACHE_OTP_TIMEOUT)
				security.SetSecureCookie(ctx, sessionToken)
				ctx.JSON(http.StatusOK, res)
				return
			}
		}
		ctx.JSON(http.StatusOK, utility.GetErrorResponse(
			"Could not create otp session!! Try later ..."))
	}
}

//UpdateMobileOTP sends OTP to new mobile number
func UpdateMobileOTP(ctx *gin.Context) {
	mobile := ctx.Param("mobile")
	res, otp := service.SendOTP(mobile, false)
	if res.Status == enums.Status.SUCCESS {
		if ok, sessID := createOTPSession(ctx, otp); ok {
			res.Body = sessID
			ctx.JSON(http.StatusOK, res)
		} else {
			ctx.JSON(http.StatusOK, utility.GetErrorResponse(
				"Could not create otp session!! Try later ..."))
		}
	}
}

func createOTPSession(ctx *gin.Context, otp string) (bool, string) {
	OTPSessionID := cachestore.GenerateSessionID(ctx)
	if len(OTPSessionID) > 0 {
		cachestore.SetCache(ctx, OTPSessionID, otp,
			common.CACHE_OTP_TIMEOUT, false)
		ctx.Writer.Header().Set(common.RESPONSE_OTP_SESSION_ID_KEY,
			OTPSessionID)
		return true, OTPSessionID
	}
	return false, ""
}

//VerifyOTP verifies OTP
func VerifyOTP(ctx *gin.Context) {
	userOTP := ctx.Param("otp")
	if len(userOTP) != common.OTP_LENGTH {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		otpSessionID := ctx.Request.Header.Get(common.REQUEST_OTP_SESSION_ID_KEY)
		if cachestore.CheckCache(ctx, otpSessionID) {
			sessionOTP := cachestore.GetCache(ctx, otpSessionID)
			if sessionOTP == userOTP {
				cachestore.RemoveCache(ctx, otpSessionID)
				ctx.JSON(http.StatusOK, utility.GetSuccessResponse(
					"OTP Verified !!"))
			} else {
				ctx.JSON(http.StatusOK, utility.GetErrorResponse(
					"Incorrect OTP !!"))
			}
		} else {
			ctx.JSON(http.StatusOK, utility.GetErrorResponse(
				"Session Expired !!"))
		}
	}
}

//AuthenticateUser authenticate and returns AppResponse object
func AuthenticateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AuthenticateUser(&user)
	if res.Status == enums.Status.SUCCESS {
		user := res.Body.(*model.User)
		sessionToken := security.GenerateToken(
			"UserID", user.UserID.Hex(),
			common.USER_SESSION_TIMEOUT)
		xsrfToken := security.GenerateToken(
			"NONCE", utility.GenerateUUID(),
			common.USER_SESSION_TIMEOUT)
		if len(sessionToken) > 0 && len(xsrfToken) > 0 {
			security.SetCookie(ctx, sessionToken,
				common.COOKIE_SESSION_TIMEOUT)
			security.SetSecureCookie(ctx, xsrfToken)
		} else {
			res = utility.GetErrorResponse(
				"Could not create session!! Try later ...")
		}
	}
	ctx.JSON(http.StatusOK, res)
}

//Logout logs out user and clear sessions
func Logout(ctx *gin.Context) {
	security.RemoveCookie(ctx)
	res := utility.GetSuccessResponse(
		"Logout Successful !!")
	ctx.JSON(http.StatusOK, res)
}

//UpdateUserSchedule updates user's time-table
func UpdateUserSchedule(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateSchedule(&user)
	ctx.JSON(http.StatusOK, res)
}
