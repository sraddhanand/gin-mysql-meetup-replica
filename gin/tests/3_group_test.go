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

func TestGroupRegister(t *testing.T) {
	testRouter := SetupRouter()

	var rf models.GroupProfile
	rf.Name = grp_name
	rf.Location = grp_location
	rf.GroupAvatar = grp_avatar
	rf.GroupContact = grp_group_contact
	rf.Pincode = grp_pincode
	rf.GID = grp_gid
	data, _ := json.Marshal(rf)
	fmt.Println(bytes.NewBufferString(string(data)))
	req, err := http.NewRequest("POST", "/v1/group/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusCreated)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var grp_resp struct {
		StatusCode string              `json:"code"`
		Message    string              `json:"msg"`
		Data       models.GroupProfile `json:"data"`
	}

	json.Unmarshal(body, &grp_resp)

	grp := grp_resp.Data
	fmt.Println(grp)
	// assert.Equal(t, usr["mobile_number"], testNumber, "user mobile number same as registered")
	assert.Equal(t, grp.Name, grp_name)
	assert.Equal(t, grp.GroupContact, grp_group_contact)
	grp_id = grp.ID
	fmt.Println(grp_id)
}

func TestUpdateGroupProfile(t *testing.T) {
	testRouter := SetupRouter()
	var up models.GroupProfile
	up.Bio = grp_bio

	data, _ := json.Marshal(up)
	url := "/v1/group/update_profile/" + fmt.Sprint(grp_gid)
	fmt.Println(bytes.NewBufferString(string(data)))
	req, err := http.NewRequest("PUT", url, bytes.NewBufferString(string(data)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	fmt.Println(resp.Body)
	assert.Equal(t, resp.Code, http.StatusOK)
}

func TestGroupProfile(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/group/profile/" + fmt.Sprint(grp_id)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var grp_resp struct {
		StatusCode string              `json:"code"`
		Message    string              `json:"msg"`
		Data       models.GroupProfile `json:"data"`
	}

	json.Unmarshal(body, &grp_resp)

	grp := grp_resp.Data
	fmt.Println(grp)
	// assert.Equal(t, usr["mobile_number"], testNumber, "user mobile number same as registered")
	assert.Equal(t, grp.Name, grp_name)
	assert.Equal(t, grp.GroupContact, grp_group_contact)
	grp_id = grp.ID
	fmt.Println(grp_id)
}

func TestGroupSearchByName(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/group/search?name=" + fmt.Sprint(grp_query_name)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var grp_resp struct {
		StatusCode string                `json:"code"`
		Message    string                `json:"msg"`
		Data       []models.GroupProfile `json:"data"`
	}

	json.Unmarshal(body, &grp_resp)

	grps := grp_resp.Data
	grp := grps[len(grps)-1]
	fmt.Println(grp)
	// assert.Equal(t, usr["mobile_number"], testNumber, "user mobile number same as registered")
	assert.Equal(t, grp.Name, grp_name)
	assert.Equal(t, grp.GroupContact, grp_group_contact)
	grp_id = grp.ID
	fmt.Println(grp_id)
}

func TestAddGroupToFavorite(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/group/add_to_favorite/" + fmt.Sprint(grp_id)
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}

// Get user favorite groups

func TestUserFavoriteGroups(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/group/favorite_groups"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var grp_resp struct {
		StatusCode string                `json:"code"`
		Message    string                `json:"msg"`
		Data       []models.GroupProfile `json:"data"`
	}

	json.Unmarshal(body, &grp_resp)

	grp := grp_resp.Data
	fmt.Println(grp)
	// assert.Equal(t, usr["mobile_number"], testNumber, "user mobile number same as registered")
	assert.Equal(t, grp[0].Name, grp_name)
	assert.Equal(t, grp[0].GroupContact, grp_group_contact)
}
