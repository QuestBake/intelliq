package controller

import (
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AddNewSchool adds new school
func AddNewSchool(ctx *gin.Context) {
	var school model.School
	err := ctx.BindJSON(&school)
	if err != nil {
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
	} else {
		res = service.AddNewSchool(&school)
	}
	ctx.JSON(http.StatusOK, res)
}

//ListAllSchools fetches all schools
func ListAllSchools(ctx *gin.Context) {
	res = service.FetchAllSchools()
	ctx.JSON(http.StatusOK, res)
}
