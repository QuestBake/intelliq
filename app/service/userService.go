package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"

	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/repo"
)

//AddNewUser adds new user
func AddNewUser(user *model.User) *dto.AppResponseDto {
	if !utility.IsValidMobile(user.Mobile) {
		return utility.GetErrorResponse(common.MSG_MOBILE_MIN_LENGTH_ERROR)
	}
	user.UserName = utility.GenerateUserName(user.FullName, user.Mobile)
	if len(user.UserName) == 0 {
		return utility.GetErrorResponse(common.MSG_FULL_NAME_ERROR)
	}
	user.FullName = strings.Title(user.FullName)
	user.Password = utility.EncryptData(common.TEMP_PWD_PREFIX + user.Mobile)
	user.CreateDate = time.Now().UTC()
	user.LastModifiedDate = user.CreateDate
	userRepo := repo.NewUserRepository()
	err := userRepo.Save(user)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_SAVE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_SAVE_SUCCESS)
}

//UpdateUser updates existing user
func UpdateUser(user *model.User) *dto.AppResponseDto {
	if !utility.IsPrimaryIDValid(user.UserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	if len(strings.Split(user.FullName, " ")) < common.FULLNAME_MIN_LENGTH {
		return utility.GetErrorResponse(common.MSG_FULL_NAME_ERROR)
	}
	user.Email = strings.ToLower(user.Email)
	user.LastModifiedDate = time.Now().UTC()
	userRepo := repo.NewUserRepository()
	err := userRepo.Update(user)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_UPDATE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_UPDATE_SUCCESS)
}

//FetchAllSchoolAdmins gets all users with role school admin for a group
func FetchAllSchoolAdmins(groupID string) *dto.AppResponseDto {
	if utility.IsStringIDValid(groupID) {
		userRepo := repo.NewUserRepository()
		users, err := userRepo.FindAllSchoolAdmins(bson.ObjectIdHex(groupID))
		if err != nil {
			fmt.Println(err.Error())
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
		return utility.GetSuccessResponse(users)
	}
	return utility.GetErrorResponse(common.MSG_INVALID_ID)
}

//FetchAllTeachers gets all teachers within school
func FetchAllTeachers(schoolID string) *dto.AppResponseDto {
	if utility.IsStringIDValid(schoolID) {
		userRepo := repo.NewUserRepository()
		users, err := userRepo.FindAllSchoolTeachers(
			bson.ObjectIdHex(schoolID), nil)
		if err != nil {
			fmt.Println(err.Error())
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
		return utility.GetSuccessResponse(users)
	}
	return utility.GetErrorResponse(common.MSG_INVALID_ID)
}

//FetchAllTeachersUnderReviewer gets all teachers under a reviewer
func FetchAllTeachersUnderReviewer(schoolID string, reviewerID string) *dto.AppResponseDto {
	if utility.IsStringIDValid(schoolID) && utility.IsStringIDValid(reviewerID) {
		userRepo := repo.NewUserRepository()
		users, err := userRepo.FindAllteachersUnderReviewer(
			bson.ObjectIdHex(schoolID), bson.ObjectIdHex(reviewerID))
		if err != nil {
			fmt.Println(err.Error())
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
		return utility.GetSuccessResponse(users)
	}
	return utility.GetErrorResponse(common.MSG_INVALID_ID)
}

//FetchSelectedTeachers gets all teachers within school for specific role
func FetchSelectedTeachers(schoolID string, roleType string) *dto.AppResponseDto {
	role, errs := strconv.Atoi(roleType)
	if errs != nil || role < common.MIN_VALID_ROLE || role > common.MAX_VALID_ROLE {
		return utility.GetErrorResponse(common.MSG_NO_ROLE)
	}
	if utility.IsStringIDValid(schoolID) {
		userRepo := repo.NewUserRepository()
		users, err := userRepo.FindAllSchoolTeachers(
			bson.ObjectIdHex(schoolID), enums.UserRole(role))
		if err != nil {
			fmt.Println(err.Error())
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
		return utility.GetSuccessResponse(users)
	}
	return utility.GetErrorResponse(common.MSG_INVALID_ID)
}

//TransferUserRole transfers user roles
func TransferUserRole(roleType string, fromUserID string, toUserID string) (*dto.AppResponseDto, []string) {
	role, errs := strconv.Atoi(roleType)
	if errs != nil || role < common.MIN_VALID_ROLE ||
		role > common.MAX_VALID_ROLE {
		return utility.GetErrorResponse(common.MSG_NO_ROLE), nil
	}
	if !utility.IsStringIDValid(fromUserID) || !utility.IsStringIDValid(toUserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID), nil
	}
	userRepo := repo.NewUserRepository()
	msg, mobiles, err := userRepo.TransferRole(enums.UserRole(role),
		bson.ObjectIdHex(fromUserID), bson.ObjectIdHex(toUserID))
	if err != nil || len(msg) > 0 {
		if len(msg) > 0 {
			return utility.GetErrorResponse(msg), nil
		}
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg), nil
		}
		return utility.GetErrorResponse(common.MSG_UPDATE_ERROR), nil
	}
	return utility.GetSuccessResponse(common.MSG_UPDATE_SUCCESS), mobiles
}

//RemoveUserFromSchool transfers user roles
func RemoveUserFromSchool(schoolID string, userID string) *dto.AppResponseDto {
	if !utility.IsStringIDValid(schoolID) || !utility.IsStringIDValid(userID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	userRepo := repo.NewUserRepository()
	err := userRepo.RemoveSchoolTeacher(bson.ObjectIdHex(schoolID), bson.ObjectIdHex(userID))
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REMOVE_USER_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_REMOVE_USER_SUCCESS)
}

//AddBulkUser adds new users
func AddBulkUser(users model.Users) *dto.AppResponseDto {
	var userList []interface{}
	for _, user := range users {
		user.Password = utility.EncryptData(common.TEMP_PWD_PREFIX + user.Mobile)
		user.UserName = utility.GenerateUserName(user.FullName, user.Mobile)
		user.CreateDate = time.Now().UTC()
		user.LastModifiedDate = time.Now()
		userList = append(userList, user)
	}
	userRepo := repo.NewUserRepository()
	err := userRepo.BulkSave(userList)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_SAVE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_SAVE_SUCCESS)
}

//UpdateBulkUsers adds new users
func UpdateBulkUsers(users model.Users) *dto.AppResponseDto {
	userRepo := repo.NewUserRepository()
	err := userRepo.BulkUpdate(users)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_SAVE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_SAVE_SUCCESS)
}

//AuthenticateUser authenicate give user and returns authenicated user details.
func AuthenticateUser(user *model.User) *dto.AppResponseDto {
	if utility.IsValidMobile(user.Mobile) {
		userRepo := repo.NewUserRepository()
		loggedUser, err := userRepo.FindOne("mobile", user.Mobile)
		if err != nil {
			return utility.GetErrorResponse(common.MSG_INVALID_CREDENTIALS_MOBILE)
		}
		if utility.ComparePasswords(loggedUser.Password, user.Password) {
			loggedUser.Password = ""
			return utility.GetSuccessResponse(loggedUser)
		}
		return utility.GetErrorResponse(common.MSG_INVALID_CREDENTIALS_PWD)
	}
	return utility.GetErrorResponse(common.MSG_MOBILE_MIN_LENGTH_ERROR)
}

//FetchUserByMobileOrID fetch user info by mobile or ID
func FetchUserByMobileOrID(key string, val string) *dto.AppResponseDto {
	var value interface{} = val
	if key == common.PARAM_KEY_ID {
		if utility.IsStringIDValid(val) {
			value = bson.ObjectIdHex(val)
		} else {
			return utility.GetErrorResponse(common.MSG_INVALID_ID)
		}
	} else if key == common.PARAM_KEY_MOBILE {
		if !utility.IsValidMobile(val) {
			return utility.GetErrorResponse(common.MSG_INVALID_MOBILE)
		}
	}
	userRepo := repo.NewUserRepository()
	user, err := userRepo.FindOne(key, value)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	user.Password = ""
	user.Days = nil
	return utility.GetSuccessResponse(user)
}

//ResetPassword resets user password
func ResetPassword(pwdDTO *dto.PasswordDto) *dto.AppResponseDto {
	if len(pwdDTO.NewPwd) < common.PWD_MIN_LENGTH {
		return utility.GetErrorResponse(common.MSG_PWD_MIN_LENGTH_ERROR)
	}
	if pwdDTO.ForgotPwd {
		if !utility.IsValidMobile(pwdDTO.Mobile) {
			return utility.GetErrorResponse(common.MSG_MOBILE_MIN_LENGTH_ERROR)
		}
		pwdDTO.NewPwd = utility.EncryptData(pwdDTO.NewPwd)
		userRepo := repo.NewUserRepository()
		err := userRepo.UpdateMobilePwd("mobile", "password", pwdDTO.Mobile, pwdDTO.NewPwd)
		if err != nil {
			fmt.Println(err.Error())
			errorMsg := utility.GetErrorMsg(err)
			if len(errorMsg) > 0 {
				return utility.GetErrorResponse(errorMsg)
			}
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
	} else {
		userRepo := repo.NewUserRepository()
		user, err := userRepo.FindOne("_id", pwdDTO.UserID)
		if err != nil {
			fmt.Println(err)
			errorMsg := utility.GetErrorMsg(err)
			if len(errorMsg) > 0 {
				return utility.GetErrorResponse(errorMsg)
			}
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
		if !utility.ComparePasswords(user.Password, pwdDTO.OldPwd) {
			return utility.GetErrorResponse(common.MSG_INVALID_CREDENTIALS_PWD)
		}
		userRepo = repo.NewUserRepository()
		pwdDTO.NewPwd = utility.EncryptData(pwdDTO.NewPwd)
		updateErr := userRepo.UpdateMobilePwd("_id", "password", user.UserID, pwdDTO.NewPwd)
		if updateErr != nil {
			fmt.Println(updateErr.Error())
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
	}
	return utility.GetSuccessResponse(common.MSG_PWD_RESET_SUCCESS)
}

//UpdateMobile updates user mobile number
func UpdateMobile(pwdDTO *dto.PasswordDto) *dto.AppResponseDto {
	if !utility.IsValidMobile(pwdDTO.Mobile) {
		return utility.GetErrorResponse(common.MSG_MOBILE_MIN_LENGTH_ERROR)
	}
	userRepo := repo.NewUserRepository()
	err := userRepo.UpdateMobilePwd("_id", "mobile", pwdDTO.UserID, pwdDTO.Mobile)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(common.MSG_MOBILE_UPDATE_SUCCESS)
}

//SendOTP send 6-digit OTP to user mobile
func SendOTP(mobile string, forgotPassword bool) (*dto.AppResponseDto, string) {
	if !utility.IsValidMobile(mobile) {
		return utility.GetErrorResponse(common.MSG_MOBILE_MIN_LENGTH_ERROR), ""
	}
	if forgotPassword {
		userRepo := repo.NewUserRepository()
		_, err := userRepo.FindOne("mobile", mobile)
		if err != nil {
			return utility.GetErrorResponse(common.MSG_INVALID_CREDENTIALS_MOBILE), ""
		}
	}
	otp := utility.GenerateRandom(common.OTP_LOWER_BOUND, common.OTP_UPPER_BOUND)
	fmt.Println("OTP=> ", otp)
	return utility.GetSuccessResponse("OTP sent successfully !!"), strconv.Itoa(otp)
}

//UpdateSchedule updates user's time-table
func UpdateSchedule(user *model.User) *dto.AppResponseDto {
	if !utility.IsPrimaryIDValid(user.UserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	userRepo := repo.NewUserRepository()
	err := userRepo.UpdateSchedule(user)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(common.MSG_SCHEDULE_UPDATE_SUCCESS)
}
