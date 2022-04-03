package tests

import (
	v1 "meetup/api/v1"
	models "meetup/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	rt := gin.Default()
	gin.SetMode(gin.TestMode)

	userRoutes := rt.Group("/v1/user")
	userRoutes.POST("/register", v1.CreateUser)
	userRoutes.POST("/login", v1.Login)
	userRoutes.GET("/profile", v1.GetUserProfile)
	userRoutes.PUT("/update_profile", v1.UpdateUserProfile)

	// userRoutes.GET("/hello", v1.GetUser)
	groupRoutes := rt.Group("/v1/group")
	// groupRoutes.GET("/welcome", v1.CreateGroup)
	groupRoutes.POST("/register", v1.CreateGroup)
	groupRoutes.GET("/profile/:groupId", v1.GetGroupProfile)
	groupRoutes.GET("/search", v1.FindGroups)
	groupRoutes.PUT("/update_profile/:groupId", v1.UpdateGroupProfile)
	groupRoutes.POST("/add_to_favorite/:groupId", v1.AddToFavorite)
	groupRoutes.GET("/favorite_groups", v1.GetUserFavoriteGroups)

	storyRoutes := rt.Group("/v1/story")
	storyRoutes.POST("/new", v1.CreateGroupStory)
	storyRoutes.GET("/of_group/:groupId", v1.GetGroupStories)
	storyRoutes.GET("/of_user/:userID", v1.GetUserStories)
	eventRoutes := rt.Group("/v1/event")
	eventRoutes.POST("/new", v1.CreateGroupEvent)
	eventRoutes.POST("/rsvp/:eventId", v1.RSVPEvent)
	eventRoutes.GET("/user_rsvps", v1.GetUserRSVPEvents)
	eventRoutes.GET("/event_rsvps/:eventId", v1.GetEventsRSVPUsers)
	eventRoutes.GET("/user_upcoming_events", v1.GetUserUpcomingEvents)
	rt.GET("/health", v1.Health)
	return rt
}

func main() {
	models.Setup()
	r := SetupRouter()
	r.Run()
}

var testNumber = "8765431234"
var testPassword = "123456"
var accessToken string

// var usr_id = "AM1234"
var usr_id int64
var updateFirstName = "Soumya"
var updateLastName = "Shree"

var grp_name = "BLR Docker Meetup"
var grp_query_name = "BLR"
var grp_location = "Bengaluru, Kaarnata"

var grp_gid = "TNRS001"
var grp_id int64
var grp_avatar = "https://images.unsplash.com/photo-1572372815316-fa5474e31ddc?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1932&q=80"
var grp_group_contact = "9265431234"
var grp_bio = "Learn, Collaborate & Dockerize! Meet other developers and ops engineers in your community that are using and learning about Docker. Docker is an open platform that helps you build, ship and run applications anytime and anywhere. Developers use Docker to modify code and to streamline application development, while operations gain support to quickly and flexibly respond to their changing needs. Docker ensures agility, portability and control for all your distributed apps."
var grp_pincode = "560103"
var str_content = "Awesome events with great minds in Bangalure DevOps"
var evnt_start_time = "07-03-2026 5:30 AM"
var evnt_name = "DevSecOps Conference"
var normal_evnt_id int64

// var event_unix_time = 1646611200

// var timings'] = [{'open_time' : '5 AM','close_time' : '8 PM'}]

// var updateMobile = "8765431244"
var updateBio = "Senior Software Engineer"
var updateBirthDate = "05/04/1995"
var updateGender = "F"
var updateLanguage = "EN"

var ret_resp struct {
	StatusCode string      `json:"code"`
	Message    string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func TestInitDB(t *testing.T) {
	models.Setup()
}
