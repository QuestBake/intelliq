package approuter

import (
	"github.com/gin-gonic/gin"

	"intelliq/app/controller"
)

var mrouter *gin.Engine

//AddRouters adding routes
func AddRouters(router *gin.Engine) {
	mrouter = router
	addMetaRouters()
	addUserRouters()
	addSchoolRouters()
	addGroupRouters()
	addQuestionRouters()
	addQuestionRequestRouters()
	addQuestionPaperRouters()
}

func addMetaRouters() {
	metaRoutes := mrouter.Group("/meta")
	{
		metaRoutes.GET("/read", controller.ReadMetaData)
		metaRoutes.POST("/add", controller.AddMetaData)
		metaRoutes.PUT("/update", controller.UpdateMetaData)
		metaRoutes.DELETE("/remove", controller.RemoveMetaData)
	}
}

func addUserRouters() {
	userRoutes := mrouter.Group("/user")
	{
		userRoutes.POST("/add", controller.AddNewUser)
		userRoutes.PUT("/update", controller.UpdateUserProfile)
		userRoutes.GET("/all/admins/:groupId", controller.ListAllSchoolAdminsUnderGroup)
		userRoutes.GET("/all/school/:schoolId", controller.ListAllTeachersUnderSchool)
		userRoutes.GET("/all/reviewer/:schoolId/:reviewerId", controller.ListAllTeachersUnderReviewer)
		userRoutes.GET("/all/school/:schoolId/:roleType", controller.ListSelectedTeachers)
		userRoutes.PUT("/role/transfer/:roleType/:fromUser/:toUser", controller.TransferRole)
		userRoutes.DELETE("/remove/:schoolId/:userId", controller.RemoveUserFromSchool)
		userRoutes.POST("/bulk/add", controller.AddBulkUsers)
		userRoutes.POST("/bulk/update", controller.UpdateBulkUsers)
		userRoutes.POST("/login", controller.AuthenticateUser)
		userRoutes.GET("/logout", controller.Logout)
		userRoutes.GET("/info/:key/:val", controller.ListUserByMobileOrID)
		userRoutes.GET("/forgot/pwd/:mobile", controller.ForgotPasswordOTP)
		userRoutes.GET("/new/mobile/:mobile", controller.UpdateMobileOTP)
		userRoutes.GET("/otp/verify/:otp", controller.VerifyOTP)
		userRoutes.POST("/reset/pwd", controller.ResetUserPassword)
		userRoutes.POST("/update/mobile", controller.UpdateUserMobile)
	}
}

func addSchoolRouters() {
	schoolRoutes := mrouter.Group("/school")
	{
		schoolRoutes.POST("/add", controller.AddNewSchool)
		schoolRoutes.GET("/all/:key/:val", controller.ListAllSchools)
		schoolRoutes.PUT("/update", controller.UpdateSchoolProfile)
		schoolRoutes.GET("info/:key/:val", controller.ListSchoolByCodeOrID)
	}
}

func addGroupRouters() {
	groupRoutes := mrouter.Group("/group")
	{
		groupRoutes.POST("/add", controller.AddNewGroup)
		groupRoutes.PUT("/update", controller.UpdateGroup)
		groupRoutes.GET("/all/:restrict", controller.ListAllGroups)
		groupRoutes.GET("/info/:key/:val", controller.ListGroupByCodeOrID)
	}
}

func addQuestionRouters() {
	quesRoutes := mrouter.Group("/question")
	{
		quesRoutes.GET("/:groupCode/:quesId", controller.FindQuestion) // find particular ques
		quesRoutes.POST("/all", controller.GetQuestionsFromBank)       // all approved ques from bank
		quesRoutes.POST("/suggestions", controller.GetQuestionSuggestions)
		quesRoutes.POST("/filter", controller.GetFilteredQuestions)
		quesRoutes.DELETE("/all/:groupCode", controller.RemoveObsoleteQuestions) // removes all obsolete ques
	}
}

func addQuestionRequestRouters() {
	requestRoutes := mrouter.Group("/question/request")
	{
		requestRoutes.POST("/add", controller.RequestAdd)              // new ques request
		requestRoutes.PUT("/update", controller.RequestUpdate)         // update approved ques || update approved rejected ques || update newly unapproved rejected ques
		requestRoutes.DELETE("/remove", controller.RequestRemoval)     //remove approved ques request || remove rejected ques request
		requestRoutes.PUT("/approve", controller.ApproveRequest)       // approve action by reviewer
		requestRoutes.PUT("/reject", controller.RejectRequest)         // reject action by reviewer
		requestRoutes.POST("/all", controller.GetReviewerRequests)     // all requests for reviewer
		requestRoutes.POST("/all/self", controller.GetTeacherRequests) // all requests for teacher
	}
}

func addQuestionPaperRouters() {
	paperRoutes := mrouter.Group("/paper")
	{
		paperRoutes.POST("/generate", controller.GenerateQuestionPaper) // creates ques paper
		paperRoutes.POST("/draft", controller.DraftTestPapers)
		paperRoutes.POST("/save", controller.SaveTestPapers)

		paperRoutes.GET("/templates/:groupCode/:teacherId", controller.GetTemplateSuggestions)
		paperRoutes.GET("/template/:groupCode/:templateId", controller.FindTemplate)
		paperRoutes.DELETE("/template/:groupCode/:testId", controller.RemoveTemplate)

		paperRoutes.GET("/release/:groupCode/:teacherId", controller.GetReleasePaperSuggestions)
		paperRoutes.GET("/drafts/:groupCode/:teacherId", controller.GetDraftSuggestions)
		paperRoutes.GET("/draft/:groupCode/:testId", controller.FindDraft)

		paperRoutes.DELETE("/remove/:groupCode/:testId", controller.RemoveDraft)

	}
}
