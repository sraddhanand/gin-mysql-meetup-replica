package models

import (
	"fmt"
	"time"
)

func CreateEvent(t_event *GroupEvent) (*GroupEvent, error) {
	tpl, validErr := validateEvent(t_event)
	if validErr != nil {
		return nil, validErr
	}
	grp := &GroupProfile{}
	if err := db.Model(&Group{}).Where("id = ?", tpl.GroupID).First(grp).Error; err != nil {
		return nil, err
	}
	tpl.GroupAvatar = grp.GroupAvatar
	tpl.GroupName = grp.Name
	if err := db.Model(&GroupEvent{}).Create(&tpl).Error; err != nil {
		return nil, err
	}
	return tpl, nil
}

func GetGroupEvents(group_id string) ([]GroupEvent, error) {
	var strs = []GroupEvent{}
	if err := db.Model(&GroupEvent{}).Where("group_id = ?", group_id).Find(&strs).Error; err != nil {
		return nil, err
	}
	return strs, nil
}

func GetAllEvents() ([]GroupEvent, error) {
	currnt_time := time.Now().Unix()
	var strs = []GroupEvent{}
	if err := db.Model(&GroupEvent{}).Where("starts_at >= ?", currnt_time).Order("starts_at asc").Find(&strs).Error; err != nil {
		return nil, err
	}
	return strs, nil
}

func UpdateEvent(group_id string) ([]Story, error) {
	var strs = []Story{}
	if err := db.Model(&Story{}).Where("group_id = ?", group_id).Find(&strs).Error; err != nil {
		return nil, err
	}
	return strs, nil
}

func nDaysToNextEvent(rep_days []uint8) int {
	today_now := time.Now()
	nth_day := int(today_now.Weekday())
	if nth_day == 7 {
		nth_day = 0
	}
	rt_next_dat := 0
	// sort.Ints(rep_days)
	for i := 0; i < len(rep_days); i++ {
		if int(rep_days[i]) > nth_day {
			rt_next_dat = int(rep_days[i]) - nth_day
			break
		}
	}
	if rt_next_dat == 0 {
		rt_next_dat = int(rep_days[0])
	}
	return rt_next_dat
}

func daysToXDay(x int) int {
	today_now := time.Now()
	diff := x - int(today_now.Weekday())
	if diff > 0 {
		return int(diff)
	} else {
		return int(7 + diff)
	}
}

func validateEvent(t_event *GroupEvent) (*GroupEvent, error) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	const date_time_format = "02-01-2006 3:04 PM"
	if len(t_event.Venue) == 0 {
		t_event.Venue = "@Group"
	}
	if len(t_event.Status) == 0 {
		t_event.Status = "Scheduled"
	}
	if t_event.ReccuringEvent {
		n_days_to_next := nDaysToNextEvent(t_event.ReccuringDays)
		// n_days_to_next := 2
		tomorrow := time.Now().AddDate(0, 0, n_days_to_next)
		start_time_string := fmt.Sprintf("%s %s", tomorrow.Format("02-01-2006"), t_event.StartTime)
		t, _ := time.ParseInLocation(date_time_format, start_time_string, loc)
		t_event.StartsAt = t.Unix()
		t_event.StartTime = start_time_string
	} else {
		t, _ := time.ParseInLocation(date_time_format, t_event.StartTime, loc)
		t_event.StartsAt = t.Unix()
	}
	return t_event, nil
}

// rsvp event
func UserRSVPEvent(event_id string, usr_mobile string) error {
	grp_event := &GroupEvent{}
	usr_query := &User{}
	if err1 := db.Model(&User{}).Where("mobile_number = ?", usr_mobile).First(usr_query).Error; err1 != nil {
		return err1
	}
	if err2 := db.Model(&GroupEvent{}).Where("id = ?", event_id).First(grp_event).Error; err2 != nil {
		return err2
	}
	err3 := db.Model(&usr_query).Association("RSVPEvents").Append(grp_event)
	fmt.Printf("Error while associating RSVPevents to user, Reason: %v\n", err3)
	return nil
}

// list events
func GetUserRSVPs(mobile string) ([]GroupEvent, error) {
	rsvp_event := []GroupEvent{}
	usr := &User{}
	if err1 := db.Model(&User{}).Where("mobile_number = ?", mobile).First(usr).Error; err1 != nil {
		return rsvp_event, err1
	}
	db.Model(&usr).Order("starts_at desc").Association("RSVPEvents").Find(&rsvp_event)
	return rsvp_event, nil
}

// get user's upcoming events
func GetUserUpcomingEvents(mobile string) ([]GroupEvent, error) {
	rsvp_event := []GroupEvent{}
	usr := &User{}
	if err1 := db.Model(&User{}).Where("mobile_number = ?", mobile).First(usr).Error; err1 != nil {
		return rsvp_event, err1
	}
	cur_time_unix := time.Now().Unix()
	db.Model(&usr).Where("starts_at > ?", cur_time_unix).Order("starts_at asc").Association("RSVPEvents").Find(&rsvp_event)
	return rsvp_event, nil
}

func GetEventRSVPUsers(evnt_id string) ([]UserAvatar, error) {
	user_avatars := []UserAvatar{}
	evnt := &GroupEvent{}
	if err1 := db.Model(&GroupEvent{}).Where("id = ?", evnt_id).First(evnt).Error; err1 != nil {
		return nil, err1
	}
	fmt.Println(evnt)
	err3 := db.Model(&evnt).Association("Users").Find(&user_avatars)
	fmt.Printf("Error while associating RSVPevents to user, Reason: %v\n", err3)
	return user_avatars, nil
}

// If daily? run a job daily to update the time
