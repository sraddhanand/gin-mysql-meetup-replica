package v1

import (
	"fmt"
	mdw "meetup/middleware"
	"meetup/models"
	"meetup/utils/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroupStory(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var stry models.Story
	c.BindJSON(&stry)
	fmt.Println(&stry)
	insertError := models.CreateStory(&stry, token_user.Mobile)
	if insertError != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusCreated, nil)
}

func GetGroupStories(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	group_id := c.Param("groupId")
	stry_profile, err := models.GetGroupStories(group_id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, stry_profile)
}

func GetUserStories(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	usr_id := c.Param("userID")
	stry_profile, err := models.GetUserStories(usr_id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, stry_profile)
}

func UpdateGroupStory(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var stry models.GroupProfile
	c.BindJSON(&stry)
	group_id := c.Param("groupId")
	err := models.UpdateGroupProfile(&stry, group_id)
	if err != nil {
		appG.RespMsg(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, nil)
}
