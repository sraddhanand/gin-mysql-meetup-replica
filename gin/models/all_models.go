package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Location struct {
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
	Street      string `json:"street"`
	ZipCode     string `json:"zip_code"`
}

type User struct {
	UserAuth       `gorm:"embedded"`
	UserAvatar     `gorm:"embedded"`
	UserProfile    `gorm:"embedded"`
	FavoriteGroups []Group      `gorm:"many2many:user_fav_groups;"`
	RSVPEvents     []GroupEvent `gorm:"many2many:user_rsvp_events;"`
}

type UserAuth struct {
	MobileNumber string `json:"mobile_number"`
	Password     string `json:"password"`
}

type UserGenInfo struct {
	UserAvatar
	UserProfile
}
type UserAvatar struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type UserProfile struct {
	ID        int64
	Role      string `json:"role"`
	Bio       string `json:"bio"`
	Age       uint8  `json:"age"`
	BirthDate string `json:"birth_date"`
	Language  string `json:"language"`
	Location  `gorm:"embedded"`
	Gender    string `json:"gender"`
}

type Claims struct {
	MobileNumber string `json:"mobile_number"`
	Name         string `json:"name"`
	jwt.StandardClaims
}

type TokenUser struct {
	MobileNumber string `json:"mobile_number"`
	Name         string `json:"name"`
}

type Group struct {
	GroupProfile `gorm:"embedded"`
	GroupEvents  []GroupEvent
}

type GroupTimings struct {
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

type GroupProfile struct {
	ID           int64
	GID          string `json:"g_id" gorm:"unique"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Bio          string `json:"bio"`
	GroupContact string `json:"group_contact"`
	Pincode      string `json:"pincode"`
	GroupAvatar  string `json:"group_avatar"`
}

type StoryRead struct {
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	Content    string `json:"content"`
	UserAvatar string `json:"user_avatar"`
}

type Story struct {
	StoryRead `gorm:"embedded"`
}

type GroupEvent struct {
	ID             int64
	GroupID        int64   `json:"group_id"`
	GroupName      string  `json:"group_name"`
	GroupAvatar    string  `json:"group_avatar"`
	Status         string  `json:"status"`     // [scheduled, started, postponed, completed, cancelled]
	StartsAt       int64   `json:"starts_at"`  // Store unix format
	StartTime      string  `json:"start_time"` // Store DD-MM-YYYY format
	ReccuringEvent bool    `json:"reccuring_event"`
	Name           string  `json:"name"`           // Store DD-MM-YYYY format
	ReccuringDays  []uint8 `json:"reccuring_days"` // store day as number ex: Sunday=1, Saturday=7
	Duration       int64   `json:"duration"`       // in minutes
	Venue          string  `json:"venue"`
	EventType      string  `json:"event_type"` // in of ['aarthi', 'bhajan', 'annadanam', 'pooja', 'yagnam', 'special', 'celebration']
	Users          []User  `gorm:"many2many:user_rsvp_events;"`
}

type GroupSearch struct {
	Name     string `form:"name"`
	Location string `form:"location"`
	Deity    string `form:"deity"`
}

type FavoriteGroup struct {
	GroupID int64
	Group   Group
	UserID  int64
	User    User
}
