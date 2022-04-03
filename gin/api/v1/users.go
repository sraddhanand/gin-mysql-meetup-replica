package v1

import (
	"fmt"
	"meetup/models"
	"net/http"

	mdw "meetup/middleware"
	"meetup/utils/app"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var usr models.User
	err := c.BindJSON(&usr)
	if err != nil {
		fmt.Println("error while binding", err)
	}
	insertError := models.CreateUser(&usr)
	if insertError != nil {
		fmt.Printf("Error while inserting new post into db, Reason: %v\n", insertError)
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusCreated, nil)
}

func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	var usr models.User
	c.BindJSON(&usr)
	jwt_token, err := models.LoginToken(usr.MobileNumber, usr.Password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, nil)
		return
	}
	appG.Response(http.StatusOK, jwt_token)
}

func GetUserProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	profile, err := models.GetUserProfile(token_user.Mobile)
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil)
		return
	}
	appG.Response(http.StatusOK, profile)
}

func UpdateUserProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	token_user, err_1 := mdw.ExtractTokenMetadata(c.Request)
	if err_1 != nil {
		appG.Response(http.StatusForbidden, nil)
		return
	}
	var usr models.User
	c.BindJSON(&usr)
	err := models.UpdateUserProfile(token_user.Mobile, &usr)
	if err != nil {
		appG.RespMsg(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, nil)
}
