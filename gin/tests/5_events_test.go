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

func TestCreateEvent(t *testing.T) {
	testRouter := SetupRouter()

	var rf models.GroupEvent
	rf.GroupID = int64(grp_id)
	rf.StartTime = evnt_start_time
	rf.Name = evnt_name
	data, _ := json.Marshal(rf)
	// fmt.Println(bytes.NewBufferString(string(data)))
	req, err := http.NewRequest("POST", "/v1/event/new", bytes.NewBufferString(string(data)))
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
	var envt_resp struct {
		StatusCode string            `json:"code"`
		Message    string            `json:"msg"`
		Data       models.GroupEvent `json:"data"`
	}
	json.Unmarshal(body, &envt_resp)
	grp := envt_resp.Data
	assert.Equal(t, grp.Name, evnt_name)
	normal_evnt_id = grp.ID
}

func TestRsvpEvent(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/event/rsvp/" + fmt.Sprint(normal_evnt_id)
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

func TestListUserRSVPs(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/event/user_rsvps"
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
		Data       []models.GroupEvent `json:"data"`
	}

	json.Unmarshal(body, &grp_resp)

	grp := grp_resp.Data
	fmt.Println(grp)
	assert.Equal(t, grp[0].Name, evnt_name)
	assert.Equal(t, grp[0].ReccuringEvent, false)
}

func TestUserUpcomingEvents(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/event/user_upcoming_events"
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
		Data       []models.GroupEvent `json:"data"`
	}

	json.Unmarshal(body, &grp_resp)

	grp := grp_resp.Data
	fmt.Println(grp)
	assert.Equal(t, len(grp), 1)
	assert.Equal(t, grp[0].Name, evnt_name)
	assert.Equal(t, grp[0].ReccuringEvent, false)
}

func TestListEventRSVPs(t *testing.T) {
	testRouter := SetupRouter()
	url := "/v1/event/event_rsvps/" + fmt.Sprint(normal_evnt_id)
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
		Data       []models.UserAvatar `json:"data"`
	}

	json.Unmarshal(body, &grp_resp)
	users := grp_resp.Data
	fmt.Println(users)
	assert.Equal(t, users[0].FirstName, updateFirstName)
	assert.Equal(t, users[0].LastName, updateLastName)
}
