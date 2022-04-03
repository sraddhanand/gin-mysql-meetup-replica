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

func TestCreateStory(t *testing.T) {
	testRouter := SetupRouter()

	var rf models.Story
	rf.StoryRead.GroupID = int64(grp_id)
	rf.Content = str_content
	data, _ := json.Marshal(rf)
	fmt.Println(bytes.NewBufferString(string(data)))
	req, err := http.NewRequest("POST", "/v1/story/new", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusCreated)

}

func TestGroupStories(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/story/of_group/" + fmt.Sprint(grp_id)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var stories_resp struct {
		StatusCode string             `json:"code"`
		Message    string             `json:"msg"`
		Data       []models.StoryRead `json:"data"`
	}

	json.Unmarshal(body, &stories_resp)

	stories := stories_resp.Data
	fmt.Println(stories)
	first_story := stories[0]

	// assert.Equal(t, usr["mobile_number"], testNumber, "user mobile number same as registered")
	assert.Equal(t, first_story.GroupID, grp_id)
	assert.Equal(t, first_story.UserID, usr_id)
	assert.Equal(t, first_story.Content, str_content)
	// grp_id = grp.ID
	// fmt.Println(grp_id)
}

func TestUserStories(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/story/of_user/" + fmt.Sprint(usr_id)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(resp.Body)

	var stories_resp struct {
		StatusCode string         `json:"code"`
		Message    string         `json:"msg"`
		Data       []models.Story `json:"data"`
	}

	json.Unmarshal(body, &stories_resp)

	stories := stories_resp.Data
	fmt.Println(stories)
	first_story := stories[0]

	assert.Equal(t, first_story.GroupID, grp_id)
	assert.Equal(t, first_story.UserID, usr_id)
	assert.Equal(t, first_story.Content, str_content)
}
