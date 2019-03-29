package controller

import (
	"intelliq/app/cachestore"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/security"
	"net/http"

	"github.com/gin-gonic/gin"

	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
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
	res := service.TransferUserRole(roleType, fromUserID, toUserID)
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

//ListUserByMobileOrID get user info by id or mobile number
func ListUserByMobileOrID(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	if len(key) == 0 || len(val) == 0 ||
		(key != common.PARAM_KEY_ID &&
			key != common.PARAM_KEY_MOBILE) {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.FetchUserByMobileOrID(key, val)
	ctx.JSON(http.StatusOK, res)
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

//UpdateUserMobile updates user mobile no post OTP verification
func UpdateUserMobile(ctx *gin.Context) {
	var pwdDto dto.PasswordDto
	err := ctx.BindJSON(&pwdDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateMobile(&pwdDto)
	ctx.JSON(http.StatusOK, res)
}

//ForgotPasswordOTP sends OTP to existing mobile number
func ForgotPasswordOTP(ctx *gin.Context) {
	mobile := ctx.Param("mobile")
	res, otp := service.SendOTP(mobile, true)
	if res.Status == enums.Status.SUCCESS {
		if createOTPSession(ctx, otp) {
			sessionToken := security.GenerateToken(
				"Contact", mobile,
				common.CACHE_OTP_TIMEOUT)
			if len(sessionToken) > 0 {
				security.SetCookie(ctx, sessionToken,
					common.CACHE_OTP_TIMEOUT)
				ctx.JSON(http.StatusOK, res)
				return
			}
		}
		ctx.JSON(http.StatusOK, utility.GetErrorResponse(
			"Could not create otp session!! Try later ..."))
	}
}

//SendUserOTP sends OTP to new mobile number
func SendUserOTP(ctx *gin.Context) {
	mobile := ctx.Param("mobile")
	res, otp := service.SendOTP(mobile, false)
	if res.Status == enums.Status.SUCCESS {
		if createOTPSession(ctx, otp) {
			ctx.JSON(http.StatusOK, res)
		} else {
			ctx.JSON(http.StatusOK, utility.GetErrorResponse(
				"Could not create otp session!! Try later ..."))
		}
	}
}

func createOTPSession(ctx *gin.Context, otp string) bool {
	OTPSessionID := cachestore.GenerateSessionID(ctx)
	if len(OTPSessionID) > 0 {
		cachestore.SetCache(ctx, OTPSessionID, otp,
			common.CACHE_OTP_TIMEOUT)
		ctx.Writer.Header().Set(common.RESPONSE_OTP_SESSION_ID_KEY,
			OTPSessionID)
		return true
	}
	return false
}

//VerifyOTP verifies OTP
func VerifyOTP(ctx *gin.Context) {
	userOTP := ctx.Param("otp")
	if len(userOTP) != 6 {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		otpSessionID := ctx.Request.Header.Get(common.REQUEST_OTP_SESSION_ID_KEY)
		if cachestore.CheckCache(ctx, otpSessionID) {
			sessionOTP := cachestore.GetCache(ctx, otpSessionID).(string)
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
