package models

import "fmt"

func CreateStory(str *Story, usr_mobile string) (err error) {
	vl_str := map[string]interface{}{}
	vl_str["group_id"] = str.StoryRead.GroupID
	vl_str["content"] = str.Content
	usr_query := &User{}
	if err := db.Model(&User{}).Where("mobile_number = ?", usr_mobile).First(usr_query).Error; err != nil {
		return err
	}
	fmt.Println(usr_query)
	vl_str["user_id"] = usr_query.ID
	vl_str["user_avatar"] = usr_query.UserAvatar.Avatar
	if err = db.Model(&Story{}).Create(vl_str).Error; err != nil {
		return err
	}
	return nil
}

func GetUserStories(user_id string) ([]StoryRead, error) {
	// var str = Story
	var strs = []StoryRead{}

	if err := db.Model(&Story{}).Where("user_id = ?", user_id).Find(&strs).Error; err != nil {
		return nil, err
	}
	return strs, nil
}

func GetGroupStories(group_id string) ([]StoryRead, error) {
	// var str = Story
	var strs = []StoryRead{}

	if err := db.Model(&Story{}).Where("group_id = ?", group_id).Find(&strs).Error; err != nil {
		return nil, err
	}
	return strs, nil
}
