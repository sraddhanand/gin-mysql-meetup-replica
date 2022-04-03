package v1

import (
	"fmt"
	mdw "meetup/middleware"
	"meetup/models"
	"meetup/utils/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroupEvent(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var evnt models.GroupEvent
	c.BindJSON(&evnt)
	fmt.Println(&evnt)
	rt_evnt, insertError := models.CreateEvent(&evnt)
	if insertError != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusCreated, rt_evnt)
}

func GetGroupEvents(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	group_id := c.Param("groupId")
	evnt_profile, err := models.GetGroupEvents(group_id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, evnt_profile)
}

func GetAllFutureEvents(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	groups, err := models.GetAllEvents()
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, groups)
}

func UpdateGroupEvent(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var evnt models.GroupProfile
	c.BindJSON(&evnt)
	group_id := c.Param("groupId")
	err := models.UpdateGroupProfile(&evnt, group_id)
	if err != nil {
		appG.RespMsg(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, nil)
}

func RSVPEvent(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	evnt_id := c.Param("eventId")
	err := models.UserRSVPEvent(evnt_id, token_user.Mobile)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, nil)
}

func GetUserRSVPEvents(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	rt_events, err := models.GetUserRSVPs(token_user.Mobile)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, rt_events)
}

func GetUserUpcomingEvents(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	rt_events, err := models.GetUserUpcomingEvents(token_user.Mobile)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, rt_events)
}

func GetEventsRSVPUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	_, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	evnt_id := c.Param("eventId")
	rt_events, err := models.GetEventRSVPUsers(evnt_id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, rt_events)
}
