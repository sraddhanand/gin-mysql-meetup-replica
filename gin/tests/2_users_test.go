package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	models "meetup/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	testRouter := SetupRouter()

	var rf models.UserAuth
	rf.MobileNumber = testNumber
	rf.Password = testPassword
	data, _ := json.Marshal(rf)
	fmt.Println(bytes.NewBufferString(string(data)))
	req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusCreated)

}

func TestLogin(t *testing.T) {
	testRouter := SetupRouter()

	var rf models.UserAuth
	rf.MobileNumber = testNumber
	rf.Password = testPassword
	data, _ := json.Marshal(rf)
	fmt.Println(bytes.NewBufferString(string(data)))
	req, err := http.NewRequest("POST", "/v1/user/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &ret_resp)
	accessToken = fmt.Sprintf("%v", ret_resp.Data)

	assert.Equal(t, resp.Code, http.StatusOK)
	fmt.Println(ret_resp.Data)
}

func TestUpdateUserProfile(t *testing.T) {
	testRouter := SetupRouter()
	var up models.UserGenInfo
	up.FirstName = updateFirstName
	up.LastName = updateLastName
	// up.MobileNumber = updateMobile
	up.Bio = updateBio
	up.BirthDate = updateBirthDate
	up.Gender = updateGender
	up.Language = updateLanguage
	// up.UID = usr_id

	data, _ := json.Marshal(up)
	fmt.Println(bytes.NewBufferString(string(data)))
	req, err := http.NewRequest("PUT", "/v1/user/update_profile", bytes.NewBufferString(string(data)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, http.StatusOK)
}

func TestUserProfile(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/v1/user/profile", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var usr_resp struct {
		StatusCode string             `json:"code"`
		Message    string             `json:"msg"`
		Data       models.UserGenInfo `json:"data"`
	}

	json.Unmarshal(body, &usr_resp)

	usr := usr_resp.Data
	// assert.Equal(t, usr["mobile_number"], testNumber, "user mobile number same as registered")
	assert.Equal(t, usr.FirstName, updateFirstName)
	assert.Equal(t, usr.LastName, updateLastName)
	assert.Equal(t, usr.Language, updateLanguage)
	// assert.Equal(t, usr["u_id"], usr_id)
	usr_id = usr.ID
}
