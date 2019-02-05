package approuter

import (
	"github.com/gin-gonic/gin"
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
}

func addMetaRouters() {
	metaRoutes := mrouter.Group("/meta")
	{
		metaRoutes.GET("/read", nil)

	}
}

func addUserRouters() {
	//userRoutes := mrouter.Group("/user"){}
}

func addSchoolRouters() {
	//	schoolRoutes := mrouter.Group("/school"){}
}

func addGroupRouters() {
	//groupRoutes := mrouter.Group("/group"){}

}

func addQuestionRouters() {
	//quesRoutes := mrouter.Group("/question"){}

}
