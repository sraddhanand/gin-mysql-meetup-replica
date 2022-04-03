package routes

import (
	v1 "meetup/api/v1"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userRoutes := router.Group("/v1/user")
	userRoutes.POST("/register", v1.CreateUser)
	userRoutes.POST("/login", v1.Login)
	userRoutes.GET("/profile", v1.GetUserProfile)
	userRoutes.PUT("/update_profile", v1.UpdateUserProfile)

	// userRoutes.GET("/hello", v1.GetUser)
	groupRoutes := router.Group("/v1/group")
	groupRoutes.POST("/register", v1.CreateGroup)
	groupRoutes.GET("/profile/:groupId", v1.GetGroupProfile)
	groupRoutes.GET("/search", v1.FindGroups)
	groupRoutes.GET("/all_groups", v1.GetAllGroups)
	groupRoutes.PUT("/update_profile/:groupId", v1.UpdateGroupProfile)
	groupRoutes.POST("/add_to_favorite/:groupId", v1.AddToFavorite)
	groupRoutes.GET("/favorite_groups", v1.GetUserFavoriteGroups)
	router.GET("/health", v1.Health)

	storyRoutes := router.Group("/v1/story")
	storyRoutes.POST("/new", v1.CreateGroupStory)
	storyRoutes.GET("/of_group/:groupId", v1.GetGroupStories)
	storyRoutes.GET("/of_user/:userID", v1.GetUserStories)

	eventRoutes := router.Group("/v1/event")
	eventRoutes.POST("/new", v1.CreateGroupEvent)
	eventRoutes.GET("/of_group/:groupId", v1.GetGroupEvents)
	eventRoutes.POST("/rsvp/:eventId", v1.RSVPEvent)
	eventRoutes.GET("/user_rsvps", v1.GetUserRSVPEvents)
	eventRoutes.GET("/all", v1.GetAllFutureEvents)
	eventRoutes.GET("/event_rsvps/:eventId", v1.GetEventsRSVPUsers)
	eventRoutes.GET("/user_upcoming_events", v1.GetUserUpcomingEvents)
}
