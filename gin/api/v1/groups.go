package v1

import (
	"fmt"
	mdw "meetup/middleware"
	"meetup/models"
	"meetup/utils/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var tmpl models.GroupProfile
	err := c.BindJSON(&tmpl)
	// fmt.Println(&tmpl)
	if err != nil {
		fmt.Println("error while binding", err)
	}
	temp, insertError := models.CreateGroup(&tmpl)
	if insertError != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusCreated, temp)
}

func GetGroupProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	postId := c.Param("groupId")
	tmpl_profile, err := models.GetGroup(postId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, tmpl_profile)
}

func FindGroups(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var search models.GroupSearch
	c.ShouldBindQuery(&search)
	tmpl_profile, err := models.FindGroups(&search)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, tmpl_profile)
}

func UpdateGroupProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var tmpl models.GroupProfile
	c.BindJSON(&tmpl)
	group_id := c.Param("groupId")
	err := models.UpdateGroupProfile(&tmpl, group_id)
	if err != nil {
		appG.RespMsg(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, nil)
}

func GetUserFavoriteGroups(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	profile, err := models.GetFavoriteGroup(token_user.Mobile)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, profile)
}

func GetAllGroups(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	groups, err := models.ListGroups()
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, groups)
}

func AddToFavorite(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}

	group_id := c.Param("groupId")
	err := models.AddGroupToFavorite(group_id, token_user.Mobile)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, nil)
}
